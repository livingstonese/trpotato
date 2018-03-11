package crawler

import "testing"
import (
	"github.com/livingstonese/trpotato"
)

var testCases = []struct{
	title string
	movie *trpotato.Movie
}{
	{
		"Savarakathi (2018)[720p HDRip - x265 - HEVC - AC3 5.1 - 900MB - ESubs - Tamil]",
		&trpotato.Movie{Title:"Savarakathi", Year:2018, Resolution:"720p", Size:"900MB"},
	},
	{
		"Akilandakodi Brahmandanayagan (2018)[1080p v3 HD - AVC - MP4 - DD 5.1 - 6.4GB - ESubs - Tamil]",
		&trpotato.Movie{Title:"Akilandakodi Brahmandanayagan", Year:2018, Resolution:"1080p", Size:"6.4GB"},
	},
	{
		"Nee Enna Maayam Seidhai (2017) Tamil 720p HDRip x264 5.1 1.4GB",
		&trpotato.Movie{Title:"Nee Enna Maayam Seidhai", Year:2017, Resolution:"720p", Size:"1.4GB"},
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
				if err != nil  {
					st.Errorf("Failed to parse movie")
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
