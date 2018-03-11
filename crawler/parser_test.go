package crawler

import "testing"
import (
	"github.com/livingstonese/trpotato"
	"github.com/pkg/errors"
)

var testCases = []struct{
	title string
	movie *trpotato.Movie
	err error
}{
	{
		"Savarakathi (2018)[720p HDRip - x265 - HEVC - AC3 5.1 - 900MB - ESubs - Tamil]",
		&trpotato.Movie{Title:"Savarakathi", Year:2018, Resolution:"720p", Size:"900MB"},
		nil,
	},
	{
		"Akilandakodi Brahmandanayagan (2018)[1080p v3 HD - AVC - MP4 - DD 5.1 - 6.4GB - ESubs - Tamil]",
		&trpotato.Movie{Title:"Akilandakodi Brahmandanayagan", Year:2018, Resolution:"1080p", Size:"6.4GB"},
		nil,
	},
	{
		title: "TamilRockers Rare Movie Picks in True HD Special Thread Links",
		movie: nil,
		err: errors.New("Failed to parse movie"),
	},
}

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.movie.Title, func(st *testing.T) {
			tc := tc
			st.Parallel()
			st.Logf("Testing %s", tc.title)
			{
				movie, err := parse(tc.title)
				if tc.err != nil  {
					return
				}
				if tc.err != err {
					st.Errorf("Expected error value %v. Got %v", tc.err, err)
					return
				}
				st.Log("Successfully parsed movie.")
				if movie.Title != tc.movie.Title {
					st.Errorf("Should parse title to %s. Received %s", tc.movie.Title, movie.Title)
				}
				st.Log("Successfully parsed movie title")
				if movie.Year != tc.movie.Year {
					st.Errorf("Should parse year to %d. Received %d", tc.movie.Year, movie.Year)
				}
				st.Log("Successfully parsed movie year")
				if movie.Resolution != tc.movie.Resolution {
					st.Errorf("Should parse resolution to %s. Received %s", tc.movie.Resolution, movie.Resolution)
				}
				st.Log("Successfully parsed movie resolution")
				if movie.Year != tc.movie.Year {
					st.Errorf("Should parse year to %d. Received %d", tc.movie.Year, movie.Year)
				}
				st.Log("Successfully parsed movie year")
			}
		})
	}
}
