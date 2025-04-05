package api

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func DisplayResults(searchRes *SearchResponse) {
	// Can't print anything if there are no results
	if len(searchRes.Results) == 0 {
		fmt.Println("no packages match search")
		return
	}

	// pretty print results
	for i, res := range searchRes.Results {
		color.New(color.FgHiMagenta).Fprintf(os.Stdout, "%d ", i+1)
		fmt.Fprintf(os.Stdout, "aur/")
		color.New(color.Bold).Fprintf(os.Stdout, "%s ", res.Name)
		color.New(color.FgGreen).Fprintf(os.Stdout, "%s\n", res.Version)
	}
}
