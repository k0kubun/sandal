package data

import (
	"fmt"
)

type (
	Pos struct {
		Line   int
		Column int
	}
)

func (x Pos) String() string {
	return fmt.Sprintf("Line: %d, Column %d", x.Line, x.Column)
}
