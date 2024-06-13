package sortbyclick

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// PrintTracks !+PrintTracks
func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-PrintTracks

// !+titlecode
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-titlecode

// !+albumcode
type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-albumcode

// !+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

// !+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

// !+lencode
type byLen []*Track

func (x byLen) Len() int           { return len(x) }
func (x byLen) Less(i, j int) bool { return x[i].Length < x[j].Length }
func (x byLen) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// !-lencode

func SortByType(sortType string) []*Track {
	switch sortType {
	case "Title":
		sort.Sort(byTitle(tracks))
	case "Album":
		sort.Sort(byAlbum(tracks))
	case "Artist":
		sort.Sort(byArtist(tracks))
	case "Year":
		sort.Sort(byYear(tracks))
	default:
		sort.Sort(byLen(tracks))
	}
	return tracks
}

func SortByTypeReverse(sortType string) []*Track {
	switch sortType {
	case "Title":
		sort.Sort(sort.Reverse(byTitle(tracks)))
	case "Album":
		sort.Sort(sort.Reverse(byAlbum(tracks)))
	case "Artist":
		sort.Sort(sort.Reverse(byArtist(tracks)))
	case "Year":
		sort.Sort(sort.Reverse(byYear(tracks)))
	default:
		sort.Sort(sort.Reverse(byLen(tracks)))
	}
	return tracks
}
