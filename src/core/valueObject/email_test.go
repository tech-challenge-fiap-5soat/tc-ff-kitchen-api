package valueobject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

func TestEmail(t *testing.T) {

	t.Run("should return true when email is valid", func(t *testing.T) {

		email := Email("test@gmail.com")

		assert.True(t, email.IsValid())
	})

	t.Run("should return false when email is invalid", func(t *testing.T) {
		email := Email("testgmail.com")

		assert.False(t, email.IsValid())
	})
}
