package cart

import "errors"

var (
	ErrItemAlreadyExistInCart = errors.New("商品已经存在")
	ErrCountInvalid           = errors.New("商品数量不能是负数")
)
