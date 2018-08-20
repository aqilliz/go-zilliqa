// Copyright 2014 The go-zilliqa Authors
// This file is part of the go-zilliqa library.
//
// The go-zilliqa library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-zilliqa library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-zilliqa library. If not, see <http://www.gnu.org/licenses/>.

package node

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Config represents a small collection of configuration values to fine tune the
// P2P network layer of a protocol stack. These values can be further extended by
// all registered services.
type Config struct {
	// Name sets the instance name of the node. It must not contain the / character and is
	// used in the devp2p node identifier. The instance name of gzil is "gzil". If no
	// value is specified, the basename of the current executable is used.
	Name string `toml:"-"`

	// UserIdent, if set, is used as an additional component in the devp2p node identifier.
	UserIdent string `toml:",omitempty"`

	// Version should be set to the version number of the program. It is used
	// in the devp2p node identifier.
	Version string `toml:"-"`

	// DataDir is the file system folder the node should use for any data storage
	// requirements. The configured data directory will not be directly shared with
	// registered services, instead those can use utility mzilods to create/access
	// databases or flat files. This enables ephemeral nodes which can fully reside
	// in memory.
	DataDir string
}

// NodeName returns the devp2p node identifier.
func (c *Config) NodeName() string {
	name := c.name()
	// Backwards compatibility: previous versions used title-cased "Gzil", keep that.
	if name == "gzil" || name == "gzil-testnet" {
		name = "Gzil"
	}
	if c.UserIdent != "" {
		name += "/" + c.UserIdent
	}
	if c.Version != "" {
		name += "/v" + c.Version
	}
	name += "/" + runtime.GOOS + "-" + runtime.GOARCH
	name += "/" + runtime.Version()
	return name
}

func (c *Config) name() string {
	if c.Name == "" {
		progname := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
		if progname == "" {
			panic("empty executable name, set Config.Name")
		}
		return progname
	}
	return c.Name
}

// DefaultConfig contains reasonable default settings.
var DefaultConfig = Config{}
