package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"

	"github.com/yagihash/quiet-internet/client"
	"github.com/yagihash/quiet-internet/config"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type listCmd struct {
	client *client.Client
}

func (*listCmd) Name() string     { return "list" }
func (*listCmd) Synopsis() string { return "List posts to stdout in JSON format" }
func (*listCmd) Usage() string {
	return `list:
	List posts.
`
}

func (l *listCmd) SetFlags(f *flag.FlagSet) {}

func (l *listCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	res, _ := l.client.ListPosts(nil)
	data, _ := json.Marshal(res)
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

	qi := client.New(cfg.Token, client.WithUserAgent(cfg.UserAgent))

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&listCmd{client: qi}, "")

	flag.Parse()
	ctx := context.Background()

	return int(subcommands.Execute(ctx))
}
