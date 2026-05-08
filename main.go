package main

import (
	"fmt"
	"icetube/core"
)

func main() {
	var search_term string
	fmt.Print("search term: ")
	fmt.Scan(&search_term)
	core.Search(search_term)
}
