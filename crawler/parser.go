package crawler

import (
	"regexp"
	"github.com/livingstonese/trpotato"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

var expressions = []struct{
	title, quality *regexp.Regexp
}{
	{
		title: regexp.MustCompile(`^([a-zA-Z0-9 ]+) \(([\d]+)\)(.*)$`),
		quality: regexp.MustCompile(`(1080p|720p|4K|4k).*\s+([\d.]+(MB|GB))`),
	},
}

func parse(title string) (*trpotato.Movie, error) {
	for _, expression := range expressions {
		movie := &trpotato.Movie{}
		var err error
		match := expression.title.FindStringSubmatch(title)
		if len(match) == 0 {
			continue
		}
		movie.Title = strings.TrimSpace(match[1])
		movie.Year, err = strconv.Atoi(match[2])
		if err != nil {
			return nil, errors.Errorf("Failed to parse year to integer %s", match[2])
		}

		if err = parseQuality(match[3], expression.quality, movie); err != nil {
			return nil, err
		}
		return movie, nil
	}
	return nil, errors.New("Failed to parse movie")
}

func parseQuality(quality string, expression *regexp.Regexp, movie *trpotato.Movie) error {
	match := expression.FindStringSubmatch(quality)
	if len(match) == 0 {
		return errors.Errorf("Failed to parse quality %s", quality)
	}
	movie.Resolution = strings.TrimSpace(match[1])
	movie.Size = strings.TrimSpace(match[2])
	return nil
}