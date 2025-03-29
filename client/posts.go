package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type ListPostsParams struct {
	Page          int    `url:"page,omitempty"`
	PerPage       int    `url:"per_page,omitempty"`
	Sort          string `url:"sort,omitempty"`
	Direction     string `url:"direction,omitempty"`
	Visibility    string `url:"visibility,omitempty"`
	CreatedAfter  string `url:"createdAfter,omitempty"`
	CreatedBefore string `url:"createdBefore,omitempty"`
	UpdatedAfter  string `url:"updatedAfter,omitempty"`
	UpdatedBefore string `url:"updatedBefore,omitempty"`
}

type Post struct {
	Slug               string    `json:"slug"`
	Title              string    `json:"title"`
	BodyCharacterCount int       `json:"bodyCharacterCount"`
	Visibility         string    `json:"visibility"`
	Tags               []string  `json:"tags"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	BodyMarkdown       string    `json:"bodyMarkdown,omitempty"`
	BodyHTML           string    `json:"bodyHTML,omitempty"`
}

type Pagination struct {
	CurrentPage int    `json:"currentPage"`
	NextPage    int    `json:"nextPage"`
	PrevPage    int    `json:"prevPage"`
	PerPage     int    `json:"perPage"`
	Sort        string `json:"sort"`
	Direction   string `json:"direction"`
}

type ListPostsResponse struct {
	Posts      []Post     `json:"posts"`
	Pagination Pagination `json:"pagination"`
}

func (c *Client) ListPosts(params *ListPostsParams) (*ListPostsResponse, error) {
	q, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	url := c.base + "/posts?" + q.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ret := &ListPostsResponse{}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) GetPost(slug string) (*Post, error) {
	url := c.base + "/posts/" + slug
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ret := &Post{}
	if err := json.NewDecoder(res.Body).Decode(ret); err != nil {
		return nil, err
	}

	return ret, nil
}
