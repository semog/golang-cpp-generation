package main

// Run the pre-process script when go generate is run
//go:generate gopp

// The driver function
func main() {
	reflected()
	reflected_sum()
	generated()
	generated_sum()
}
