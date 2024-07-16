package valueobject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

func TestParseOrderStatus(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected valueobject.OrderStatus
		err      bool
	}{
		{
			name:     "Valid order status",
			input:    "PREPARING",
			expected: valueobject.PREPARING,
			err:      false,
		},
		{
			name:     "Valid order status with lower case",
			input:    "preparing",
			expected: valueobject.PREPARING,
			err:      false,
		},
		{
			name:     "Invalid order status",
			input:    "INVALID",
			expected: "",
			err:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			o, err := valueobject.ParseOrderStatus(tc.input)

			if tc.err && err == nil {
				t.Errorf("expected error but got none")
			}

			if !tc.err && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if o != tc.expected {
				t.Errorf("expected %s but got %s", tc.expected, o)
			}
		})
	}
}

func TestGetPreviousStatus(t *testing.T) {
	t.Run("should return previous status for PREPARING", func(t *testing.T) {
		previousStatus := valueobject.PREPARING.GetPreviousStatus()
		expectedStatus := []valueobject.OrderStatus{valueobject.AWAITING_PREPARATION}

		assert.Equal(t, expectedStatus, previousStatus)
	})

	t.Run("should return previous status for READY_TO_TAKEOUT", func(t *testing.T) {
		previousStatus := valueobject.READY_TO_TAKEOUT.GetPreviousStatus()
		expectedStatus := []valueobject.OrderStatus{valueobject.AWAITING_PREPARATION, valueobject.PREPARING}

		assert.Equal(t, expectedStatus, previousStatus)
	})

	t.Run("should return previous status for COMPLETED", func(t *testing.T) {
		previousStatus := valueobject.COMPLETED.GetPreviousStatus()
		expectedStatus := []valueobject.OrderStatus{valueobject.AWAITING_PREPARATION, valueobject.PREPARING, valueobject.READY_TO_TAKEOUT}

		assert.Equal(t, expectedStatus, previousStatus)
	})

	t.Run("should return empty previous status for unknown status", func(t *testing.T) {
		unknownStatus := valueobject.OrderStatus("UNKNOWN")
		previousStatus := unknownStatus.GetPreviousStatus()

		assert.Empty(t, previousStatus)
	})
}

func TestIsValidNextStatus(t *testing.T) {
	t.Run("should return true when next status is valid", func(t *testing.T) {
		currentStatus := valueobject.OrderStatus("AWAITING_PREPARATION")
		nextStatus := "PREPARING"

		result := currentStatus.IsValidNextStatus(nextStatus)

		assert.True(t, result)
	})

	t.Run("should return false when next status is not valid", func(t *testing.T) {
		currentStatus := valueobject.OrderStatus("AWAITING_PREPARATION")
		nextStatus := "COMPLETED"

		result := currentStatus.IsValidNextStatus(nextStatus)

		assert.False(t, result)
	})

	t.Run("should return false when next status is the same as current status", func(t *testing.T) {
		currentStatus := valueobject.OrderStatus("PREPARING")
		nextStatus := "PREPARING"

		result := currentStatus.IsValidNextStatus(nextStatus)

		assert.False(t, result)
	})

	t.Run("should return false when next status cannot be parsed", func(t *testing.T) {
		currentStatus := valueobject.OrderStatus("AWAITING_PREPARATION")
		nextStatus := "INVALID_STATUS"

		result := currentStatus.IsValidNextStatus(nextStatus)

		assert.False(t, result)
	})
}

func TestOrderStatusAsString(t *testing.T) {
	result := valueobject.AWAITING_PREPARATION.String()

	assert.Equal(t, "AWAITING_PREPARATION", result)
}
