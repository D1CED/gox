package gox

import "fmt"

// Symbol holds a character representing a player. Make sure they are uniqe and non-zero.
type Symbol byte

func (s Symbol) String() string {
	if s == 0 {
		return " "
	}
	return fmt.Sprintf("%c", s)
}
