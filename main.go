package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"go.i3wm.org/i3/v4"
)

func ContainerByTabIndex(tc *i3.Node, index uint64) (*i3.Node, error) {
	if index < 1 {
		return nil, fmt.Errorf("tab index is %d, but must be 1 or greater", index)
	}

	highestTabIndex := uint64(len(tc.Nodes))
	if index > highestTabIndex {
		return nil, fmt.Errorf("tab index is %d, but highest index is %d", index, highestTabIndex)
	}

	return tc.Nodes[index-1], nil
}

func FindFocusedTabbedContainer() (*i3.Node, error) {
	tree, err := i3.GetTree()
	if err != nil {
		return nil, err
	}

	return tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Layout == i3.Tabbed
	}), nil
}

func Focus(n *i3.Node) error {
	if _, err := i3.RunCommand(fmt.Sprintf(`[con_id="%d"] focus`, n.ID)); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s TABINDEX", os.Args[0])
	}

	index, err := strconv.ParseUint(os.Args[1], 10, 8)
	if err != nil {
		log.Fatalf("tab index must be a number: %q: %s", os.Args[1], err)
	}

	tc, err := FindFocusedTabbedContainer()
	if err != nil {
		log.Fatalf("failed to find focused tabbed container: %s", err)
	}
	if tc == nil {
		log.Fatalf("focused container is not tabbed")
	}

	c, err := ContainerByTabIndex(tc, index)
	if err != nil {
		log.Fatalf("failed to find container by index: %s", err)
	}

	if err := Focus(c); err != nil {
		log.Fatalf("could not focus container at tab index %d", index)
	}
}
