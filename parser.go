package filtr

import (
	"strings"
)

type Item struct {
	Pos int
	End int
}

// Checks if a slice could possibly be an exception - if there is an exception that begins on a slice.
func IsPossibleException(slice string) bool {
	for _, v := range Exceptions {
		if strings.HasPrefix(v, slice) {
			return true
		}
	}
	return false
}

// Checks if a slice is perfectly equal to one of exceptions.
func IsException(slice string) bool {
	for _, v := range Exceptions {
		if v == slice {
			return true
		}
	}
	return false
}

// Applies all possible replacements to bypass profanities (leet, cyrillic, etc.)
func ApplyReplacements(slice string) string {
	res := slice
	for k, v := range Replacements {
		res = strings.ReplaceAll(res, string(k), string(v))
	}
	return res
}

// Checks if a slice of a string is a profanity.
func IsProfane(slice string) bool {
	str := slice
	for k, v := range Replacements {
		str = strings.ReplaceAll(str, string(k), string(v))
	}

	if IsException(slice) {
		return false
	}

	for _, v := range Dictionary {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}

// Parses a string character by character, analyzing for existence of any profanities.
func Parse(text string) []Item {
	lower := strings.ToLower(text)
	res := []Item{}

	for i := 0; i < len(text); i++ {
		for j := i; j < len(text); j++ {
			if IsProfane(lower[i:j]) {
				if IsPossibleException(lower[i:j]) {
					continue
				}
				res = append(res, Item{i, j})
				break
			}
		}
	}

	return res
}
