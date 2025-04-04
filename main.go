package main

import (
	"fmt"
	"os"

	"github.com/UberChili/auron/api"
	"github.com/fatih/color"
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

	for i, result := range response.Results {
		color.New(color.FgRed).Fprintf(os.Stdout, "%d ", i+1)
		fmt.Fprintf(os.Stdout, "aur/")
		color.New(color.Bold).Fprintf(os.Stdout, "%s\n", result.Name)
		// color.New(color.FgCyan).Fprintf(os.Stdout, "%s\n", result.Name)

		// fmt.Printf("%s  %s\n   %s\n", result.Name, result.Version, result.Description)
	}

}
