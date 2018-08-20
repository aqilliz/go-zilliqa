// Copyright 2018 The go-zilliqa Authors
// This file is part of go-zilliqa.
//
// go-zilliqa is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-zilliqa is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-zilliqa. If not, see <http://www.gnu.org/licenses/>.

// gzil is the official command-line client for Zilliqa.
package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	godebug "runtime/debug"
	"sort"
	"strconv"

	"github.com/aqilliz/go-zilliqa/cmd/utils"
	"github.com/aqilliz/go-zilliqa/node"
	"github.com/elastic/gosigar"
	"gopkg.in/urfave/cli.v1"
)

const (
	clientIdentifier = "gzil" // Client identifier to advertise over the network
)

var (
	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// The app that holds all commands and flags.
	app = utils.NewApp(gitCommit, "the go-zilliqa command line interface")
	// flags that configure the node
	nodeFlags = []cli.Flag{
		utils.IdentityFlag,
	}
)

func init() {
	// Initialize the CLI app and start Geth
	app.Action = gzil
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2018 The go-zilliqa Authors"
	app.Commands = []cli.Command{
		// See config.go
		dumpConfigCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = append(app.Flags, nodeFlags...)

	app.Before = func(ctx *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())

		// Cap the cache allowance and tune the garbage collector
		var mem gosigar.Mem
		if err := mem.Get(); err == nil {
			allowance := int(mem.Total / 1024 / 1024 / 3)
			if cache := ctx.GlobalInt(utils.CacheFlag.Name); cache > allowance {
				fmt.Println("Sanitizing cache to Go's GC limits", "provided", cache, "updated", allowance)
				ctx.GlobalSet(utils.CacheFlag.Name, strconv.Itoa(allowance))
			}
		}
		// Ensure Go's GC ignores the database cache for trigger percentage
		cache := ctx.GlobalInt(utils.CacheFlag.Name)
		gogc := math.Max(20, math.Min(100, 100/(float64(cache)/1024)))

		fmt.Println("Sanitizing Go's GC trigger", "percent", int(gogc))
		godebug.SetGCPercent(int(gogc))

		return nil
	}

	app.After = func(ctx *cli.Context) error {
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// gzil is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
func gzil(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}
	node := makeFullNode(ctx)
	startNode(ctx, node)
	node.Wait()
	return nil
}

// startNode boots up the system node and all registered protocols, after which
// it unlocks any requested accounts, and starts the RPC/IPC interfaces and the
// miner.
func startNode(ctx *cli.Context, stack *node.Node) {
	// Start up the node itself
	utils.StartNode(stack)
}
