package code

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/inhies/go-bytesize"
)

func Format(size int, path string, human bool) string {
	if human {
		humanSize := bytesize.New(float64(size))
		return normalizeHumanByteSize(humanSize)
	}

	return fmt.Sprintf("%sB", strconv.Itoa(size))
}

func normalizeHumanByteSize(humanSize bytesize.ByteSize) string {
	str := humanSize.String()
	divider := 0

	for i, r := range str {
		if !unicode.IsDigit(r) && r != '.' {
			divider = i
			break
		}
	}

	num := str[:divider]
	unit := str[divider:]

	if unit == "B" {
		num = strings.TrimRight(num, "0")
		num = strings.TrimRight(num, ".")
	} else {
		if strings.Contains(num, ".") && strings.HasSuffix(num, "00") {
			num = strings.TrimSuffix(num, "0")
		}
	}

	return fmt.Sprintf("%s%s", num, unit)
}
