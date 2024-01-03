package order

import "errors"

var (
	ErrEmptyCartFound       = errors.New("购物车是空的")
	ErrInvalidOrderID       = errors.New("无效订单")
	ErrCancelDurationPassed = errors.New("已通过取消持续时间")
	ErrNotEnoughStock       = errors.New("没有足够库存")
)
