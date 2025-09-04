# Go Practice

A collection of Go programming exercises, algorithm implementations, and coding challenges.

## Project Structure

This repository contains various Go programming exercises and implementations organized into the following categories:

### Algorithm Implementations

- **grokking-algorithms**: Implementations of algorithms from the book "Grokking Algorithms" by Aditya Bhargava
  - Binary Search (Chapter 1)
  - Selection Sort (Chapter 2)
  - Recursion (Chapter 3)
  - Quicksort (Chapter 4)
  - Hash Tables (Chapter 5)
  - Breadth-First Search (Chapter 6)
  - Dijkstra's Algorithm (Chapter 7)
  - Set Covering Problem (Chapter 8)

### LeetCode Solutions

- **leet-code/top150interview**: Solutions to LeetCode's Top 150 Interview Questions
  - #26: Remove Duplicates from Sorted Array
  - #27: Remove Element
  - #88: Merge Sorted Array
  - #169: Majority Element

### Standalone Practice Exercises

- **fibonacci**: Fibonacci sequence implementation
- **factorial**: Factorial function implementation
- **prime**: Prime number checking and generation
- **min-max**: Min and max function implementations
- **count-characters**: Character counting in strings
- **structs**: Practice with Go structures
- **concurrency**: Go concurrency patterns and examples

### Gophercises

- **gophercises**: Solutions to Jon Calhoun's Gophercises course

## Getting Started

Each directory contains standalone Go modules. To run the code in any directory:

1. Navigate to the specific directory

   ```
   cd directory-name
   ```

2. Run the Go code

   ```
   go run *.go
   ```

3. Run tests (if available)
   ```
   go test -v
   ```

## Testing

Most implementations include test files (with `_test.go` suffix) to verify correctness. Run the tests using the `go test` command.

## License

This project is licensed under the terms of the LICENSE file included in this repository.

