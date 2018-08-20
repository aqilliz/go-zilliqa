// Copyright 2015 The go-zilliqa Authors
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
	"path/filepath"
	"sync"
)

// Node is a container on which services can be registered.
type Node struct {
	stop chan struct{} // Channel to wait for termination notifications
	lock sync.RWMutex
}

// New creates a new node, ready for protocol registration.
func New(conf *Config) (*Node, error) {
	confCopy := *conf
	conf = &confCopy
	if conf.DataDir != "" {
		absdatadir, err := filepath.Abs(conf.DataDir)
		if err != nil {
			return nil, err
		}
		conf.DataDir = absdatadir
	}

	return &Node{}, nil
}

// Start create a live P2P node and starts running it.
func (n *Node) Start() error {
	n.lock.Lock()
	defer n.lock.Unlock()

	n.stop = make(chan struct{})

	return nil
}

// Stop terminates a running node along with all it's services. In the node was
// not started, an error is returned.
func (n *Node) Stop() error {
	n.lock.Lock()
	defer n.lock.Unlock()

	// unblock n.Wait
	close(n.stop)

	return nil
}

// Wait blocks the thread until the node is stopped. If the node is not running
// at the time of invocation, the method immediately returns.
func (n *Node) Wait() {
	n.lock.RLock()

	stop := n.stop
	n.lock.RUnlock()

	<-stop
}

// Restart terminates a running node and boots up a new one in its place. If the
// node isn't running, an error is returned.
func (n *Node) Restart() error {
	if err := n.Stop(); err != nil {
		return err
	}
	if err := n.Start(); err != nil {
		return err
	}
	return nil
}
