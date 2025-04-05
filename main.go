package main

import (
	"fmt"
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

	response, err := api.SearchPackage(packageName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Only list the search results, for now
	api.DisplayResults(response)
}
