package cli

import (
	"fmt"
	"icetube/core"
	"log"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search YouTube",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		html, err := core.Search(query)
		if err != nil {
			log.Fatal(err)
		}
		jsonStr := core.Extractor(html)
		results, err := core.Parse(jsonStr)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range results.Videos {
			fmt.Println(v.Title)
		}
	},
}
