package kvstore

import (
	"fmt"
	"os"
	"strings"
)

// Handle executes the commands entered through REPL interface
func Handle(s string, store *Store) (string, error) {
	args := strings.Split(s, " ")

	switch strings.ToLower(args[0]) {
	case "read":
		return store.Read(args[1:])
	case "write":
		return "", store.Write(args[1:])
	case "delete":
		return "", store.Delete(args[1:])
	case "start":
		if len(args) != 1 {
			return "", fmt.Errorf("Incorrect input. Usage: START")
		}

		// creating a new KV layer for the new transaction and pushing it
		// into the stack
		layer := make(map[string]string)
		for k, v := range store.KVStack.Peek().(map[string]string) {
			layer[k] = v
		}
		store.KVStack.Push(layer)
		store.Depth++
	case "commit":
		if len(args) != 1 {
			return "", fmt.Errorf("Incorrect input. Usage: COMMIT")
		}
		if store.Depth == 0 {
			return "", fmt.Errorf("Not in a transaction")
		}

		// apply the most recent updates from the deepest transaction
		// into its previous transaction. Common keys will be overwritten.
		topLayer := store.KVStack.Pop().(map[string]string)
		secondTopLayer := store.KVStack.Pop().(map[string]string)
		for k, v := range topLayer {
			secondTopLayer[k] = v
		}
		//deleting the existing keys which have been deleted in this transaction
		for k, _ := range secondTopLayer {
			_, ok := topLayer[k]
			if !ok {
				delete(secondTopLayer, k)
			}
		}
		store.KVStack.Push(secondTopLayer)
		store.Depth--
	case "abort":
		if len(args) != 1 {
			return "", fmt.Errorf("Incorrect input. Usage: ABORT")
		}
		store.KVStack.Pop()
		store.Depth--
	case "quit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
	}
	return "", nil
}
