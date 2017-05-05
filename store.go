package kvstore

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

// Store is a struct that defines an in-memory kv storage with support for transactions.
// To enable transaction support, we define a stack which keeps a layer of KV maps for each
// open transactions (except for the initial layer). The Depth variable keeps track of
// nesting level, which will be used by commit or abort operations.
type Store struct {
	Depth   int
	KVStack *stack.Stack
}

// read reads a key out of the most recent tranaction layer or the default layer.
func (s *Store) Read(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect input. Usage: READ <key>")
	}
	topLayer := s.KVStack.Peek().(map[string]string)
	i, ok := topLayer[args[0]]
	if !ok {
		return "", fmt.Errorf("Key not found: %v", args[0])
	}
	return i, nil
}

// write writes a key-value pair to the most recent transaction or the default layer.
func (s *Store) Write(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect input. Usage: WRITE <key> <value>")
	}
	topLayer := s.KVStack.Pop().(map[string]string)
	topLayer[args[0]] = args[1]
	s.KVStack.Push(topLayer)
	return nil
}

// del deletes a key-value pair from the most recent transaction or the default layer.
func (s *Store) Delete(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Incorrect input. Usage: DELETE <key>")
	}
	topLayer := s.KVStack.Pop().(map[string]string)
	_, ok := topLayer[args[0]]
	if !ok {
		return fmt.Errorf("Key not found.")
	}
	delete(topLayer, args[0])
	s.KVStack.Push(topLayer)
	return nil
}
