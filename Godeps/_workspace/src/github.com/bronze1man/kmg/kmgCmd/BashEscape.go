package kmgCmd

import (
	"strings"
)

// escape string to put it into bash
func BashEscape(inS string) (outS string) {
	return "'" + strings.Replace(inS, "'", "'''", -1) + "'"
}
