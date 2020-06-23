package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"go.i3wm.org/i3/v4"
)

func GetTab(i int) (*i3.Node, error) {
	tree, err := i3.GetTree()
	if err != nil {
		return nil, err
	}

	tc := tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Layout == i3.Tabbed
	})
	if tc == nil {
		return nil, fmt.Errorf("could not find tabbed container with focus")
	}

	if i < 0 || i > len(tc.Nodes) {
		return nil, fmt.Errorf("no tab at index %d", i)
	}
	return tc.Nodes[i], nil
}

func Focus(n *i3.Node) error {
	if _, err := i3.RunCommand(fmt.Sprintf(`[con_id="%d"] focus`, n.ID)); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) == 2 {
		tab := os.Args[1]
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("invalid tab %s", tab)
		}
		n, err := GetTab(i)
		if err != nil {
			log.Fatal(err)
		}
		if err := Focus(n); err != nil {
			log.Fatalf("could not focus on tab %d", i)
		}
	}
}
