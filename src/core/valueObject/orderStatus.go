package valueobject

import (
	"fmt"
	"slices"
	"strings"
)

type OrderStatus string

const (
	AWAITING_PREPARATION OrderStatus = "AWAITING_PREPARATION"
	PREPARING            OrderStatus = "PREPARING"
	READY_TO_TAKEOUT     OrderStatus = "READY_TO_TAKEOUT"
	COMPLETED            OrderStatus = "COMPLETED"
)

func (o OrderStatus) String() string {
	return string(o)
}

func ParseOrderStatus(s string) (o OrderStatus, err error) {
	statuses := map[OrderStatus]struct{}{
		AWAITING_PREPARATION: {},
		PREPARING:            {},
		READY_TO_TAKEOUT:     {},
		COMPLETED:            {},
	}

	orderStatus := OrderStatus(strings.ToUpper(s))
	_, ok := statuses[orderStatus]

	if !ok {
		return o, fmt.Errorf(`cannot parse:[%s] as order status`, s)
	}
	return orderStatus, nil
}

func (o OrderStatus) AvailableNextStatus(status OrderStatus) []OrderStatus {
	switch status {
	case AWAITING_PREPARATION:
		return []OrderStatus{PREPARING}
	case PREPARING:
		return []OrderStatus{READY_TO_TAKEOUT}
	case READY_TO_TAKEOUT:
		return []OrderStatus{COMPLETED}
	default:
		return []OrderStatus{}
	}
}

func (o OrderStatus) IsValidNextStatus(nextStatus string) bool {
	currentStatus, err := ParseOrderStatus(o.String())
	allowSameStatus := false

	if err != nil {
		return false
	}

	nextStatusParsed, err := ParseOrderStatus(nextStatus)

	if err != nil {
		return false
	}

	if nextStatusParsed == currentStatus {
		return allowSameStatus
	}

	return slices.Contains(currentStatus.AvailableNextStatus(currentStatus), nextStatusParsed)
}

func (o OrderStatus) GetPreviousStatus() []OrderStatus {
	switch o {
	case PREPARING:
		return []OrderStatus{AWAITING_PREPARATION}
	case READY_TO_TAKEOUT:
		return []OrderStatus{AWAITING_PREPARATION, PREPARING}
	case COMPLETED:
		return []OrderStatus{AWAITING_PREPARATION, PREPARING, READY_TO_TAKEOUT}
	default:
		return []OrderStatus{}
	}
}

func (o OrderStatus) OrderCanBeUpdated() bool {
	return o == AWAITING_PREPARATION
}

func (o OrderStatus) OrderIsCompleted() bool {
	return o == COMPLETED
}
