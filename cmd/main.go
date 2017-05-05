package main

import (
	"fmt"
	"os"

	"github.com/golang-collections/collections/stack"
	"github.com/mazarmi/kvstore"
	"github.com/peterh/liner"
)

func main() {
	store := &kvstore.Store{
		KVStack: stack.New(),
		Depth:   0,
	}
	// kvlayer is the default/initial storage layer
	kvlayer := map[string]string{}
	store.KVStack.Push(kvlayer)

	line := liner.NewLiner()
	defer line.Close()

	for {
		cmd, err := line.Prompt("> ")
		v, err := kvstore.Handle(cmd, store)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			if v != "" {
				fmt.Println(v)
			}
		}
	}
}
