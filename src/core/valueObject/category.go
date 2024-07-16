package valueobject

import (
	"slices"
	"strings"
)

type Category string

func (category Category) IsValid() bool {
	categoryList := []string{
		"acompanhamento",
		"bebida",
		"sobremesa",
		"lanche",
	}

	return slices.Contains(categoryList, strings.ToLower(string(category)))
}
