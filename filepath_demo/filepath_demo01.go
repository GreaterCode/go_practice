package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// base
	fmt.Println("base path:", filepath.Base(os.Args[0]))

	// abs
	fmt.Println("abs path:", filepath.Abs(".go")

}
