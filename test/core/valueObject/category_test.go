package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	vo "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

func TestCategory(t *testing.T) {
	t.Run("should return true when category is Lanche", func(t *testing.T) {
		isValid := vo.Category("Lanche").IsValid()

		assert.True(t, isValid)
	})
	t.Run("should return true when category is Bebida", func(t *testing.T) {
		isValid := vo.Category("Bebida").IsValid()

		assert.True(t, isValid)
	})
	t.Run("should return true when category is Acompanhamento", func(t *testing.T) {
		isValid := vo.Category("Acompanhamento").IsValid()

		assert.True(t, isValid)
	})
	t.Run("should return true when category is Sobremesa", func(t *testing.T) {
		isValid := vo.Category("Sobremesa").IsValid()

		assert.True(t, isValid)
	})
	t.Run("should return false when category is unkown", func(t *testing.T) {
		isValid := vo.Category("Não mapeada").IsValid()

		assert.False(t, isValid)
	})
}
