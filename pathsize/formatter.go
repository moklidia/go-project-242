package pathsize

import (
	"fmt"
	"strconv"

	"github.com/inhies/go-bytesize"
)

func Format(size int, path string, human bool) string {
	if human {
		humanSize := bytesize.New(float64(size))
		return fmt.Sprintf("%s\t%s", humanSize, path)
	}

	return fmt.Sprintf("%s\t%s", strconv.Itoa(size), path)
}
