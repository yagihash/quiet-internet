package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"

	"github.com/yagihash/quiet-internet/client"
	"github.com/yagihash/quiet-internet/cmd/qi/config"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type listCmd struct {
	client *client.Client

	page          int
	perPage       int
	sort          string
	direction     string
	visibility    string
	createdAfter  string
	createdBefore string
	updatedAfter  string
	updatedBefore string
}

func (*listCmd) Name() string     { return "list" }
func (*listCmd) Synopsis() string { return "Fetch the list of posts in JSON format" }
func (*listCmd) Usage() string {
	return `list:
	List posts.
`
}

func (l *listCmd) SetFlags(f *flag.FlagSet) {
	// TODO:validate values
	f.IntVar(&l.page, "page", 0, "page number")
	f.IntVar(&l.perPage, "per-page", 0, "per-page")
	f.StringVar(&l.sort, "sort", "", "sort field [created, updated]")
	f.StringVar(&l.direction, "direction", "", "direction field [asc, desc]")
	f.StringVar(&l.visibility, "visibility", "", "visibility field [MYSELF, ANYONE, URL_ONLY]")
	f.StringVar(&l.createdAfter, "created-after", "", "filter with created field in ISO8601 format")
	f.StringVar(&l.createdBefore, "created-before", "", "filter with created field in ISO8601 format")
	f.StringVar(&l.updatedAfter, "updated-after", "", "filter with updated field in ISO8601 format")
	f.StringVar(&l.updatedBefore, "updated-before", "", "filter with updated field in ISO8601 format")
}

func (l *listCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	params := &client.ListPostsParams{
		Page:          l.page,
		PerPage:       l.perPage,
		Sort:          l.sort,
		Direction:     l.direction,
		Visibility:    l.visibility,
		CreatedAfter:  l.createdAfter,
		CreatedBefore: l.createdBefore,
		UpdatedAfter:  l.updatedAfter,
		UpdatedBefore: l.updatedBefore,
	}
	res, err := l.client.ListPosts(params)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	fmt.Printf("%s\n", data)
	return subcommands.ExitSuccess
}

type getCmd struct {
	client *client.Client

	slug string
}

func (*getCmd) Name() string     { return "get" }
func (*getCmd) Synopsis() string { return "Fetch a post in JSON format" }
func (*getCmd) Usage() string {
	return `get:
	Get a post.
`
}

func (g *getCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&g.slug, "slug", "", "slug")
}

func (g *getCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	res, err := g.client.GetPost(g.slug)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	fmt.Printf("%s\n", data)
	return subcommands.ExitSuccess
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	ua := "qi-v0.0.1" // TODO:use git tag info
	qi := client.New(cfg.Token, client.WithUserAgent(ua))

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&listCmd{client: qi}, "")
	subcommands.Register(&getCmd{client: qi}, "")

	flag.Parse()
	ctx := context.Background()

	return int(subcommands.Execute(ctx))
}
