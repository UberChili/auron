package main

import (
	"fmt"
	"log"
	"os"

	"github.com/UberChili/auron/api"
)

func main() {
	args := os.Args
	var packageName string

	if len(args) != 2 {
		fmt.Println("No package provided, do a system upgrade")
		os.Exit(1)
	} else {
		packageName = args[1]
	}

	// Make API call to search for package
	data, err := api.FetchData(packageName)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	fmt.Println("API Response: ", data)
}
