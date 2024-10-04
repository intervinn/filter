package filtr

import (
	"strings"
)

type Item struct {
	Pos int
	End int
}

func IsPossibleException(slice string) bool {
	for _, v := range Exceptions {
		if strings.HasPrefix(v, slice) {
			return true
		}
	}
	return false
}

func ApplyReplacements(slice string) string {
	res := slice
	for k, v := range Replacements {
		res = strings.ReplaceAll(res, string(k), string(v))
	}
	return res
}

func IsProfane(slice string) bool {
	str := slice
	for k, v := range Replacements {
		str = strings.ReplaceAll(str, string(k), string(v))
	}

	for _, v := range Exceptions {
		if str == v {
			return false
		}
	}

	for _, v := range Dictionary {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}

func IsException(slice string) bool {
	for _, v := range Exceptions {
		if v == slice {
			return true
		}
	}
	return false
}

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
