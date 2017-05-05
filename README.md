
# kvstore - A Simple Transactional in-Memory Key-Value Store

## Overview
The goal of this project is to design a simple in-memory key-value store, which supports nested transactions. Each transaction can then be committed or aborted.

### Commands
1. **READ** <key> Reads and prints, to stdout, the val associated with key. If the value is not present an error is printed to stderr.
2. **WRITE** <key> <val> Stores val in key.
3. **DELETE** <key> Removes all key from store. Future READ commands on that key will return an error.
4. **START** Start a transaction.
5. **COMMIT** Commit a transaction. All actions in the current transaction are committed to the parent transaction or the root store. If there is no current transaction an error is output to stderr.
6. **ABORT** Abort a transaction. All actions in the current transaction are discarded.
7. **QUIT** Exit the REPL cleanly. A message to stderr may be output.

## Build from the source code
1. Install *glide* for dependency management: https://glide.sh/

2. Run `make init`

3. Run `make build`

and then run the executable (`cmd/main`) to enter the REPL interface.

## To Do
- Add support for highly concurrent access to the kv store
- Scaling the kv store (sharding, repliation, etc.)
- Performance profiling and improvement
- Add logging and creating a docker image
- Perform more rigorous testing
