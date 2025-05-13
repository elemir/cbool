package cbool

import "golang.org/x/exp/constraints"

func FromInt[I constraints.Integer](i I) bool {
	if i == 0 {
		return false
	}

	return true
}

func ToInt[I constraints.Integer](b bool) I {
	if b {
		return 1
	}

	return 0
}
