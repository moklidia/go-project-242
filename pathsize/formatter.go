package pathsize

import (
	"fmt"
	"strconv"
)

func Format(size int, path string, _ bool) string {
	return fmt.Sprintf("%s\t%s", strconv.Itoa(size), path)
}
