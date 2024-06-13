package main

import "sortbyclick"

func main() {
	res := sortbyclick.SortByType("Artist")
	res = sortbyclick.SortByTypeReverse("Artist")
	sortbyclick.PrintTracks(res)
}
