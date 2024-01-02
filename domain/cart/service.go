package cart

import (
	"errors"
	"shopping_go/domain/product"
)

type Service struct {
	cartRepository     Repository
	cartItemRepository ItemRepository
	productRepository  product.Repository
}

// NewService 实例化service
func NewService(cartRepository Repository, itemRepository ItemRepository, productRepository product.Repository) *Service {
	return &Service{
		cartRepository:     cartRepository,
		cartItemRepository: itemRepository,
		productRepository:  productRepository,
	}
}

// AddItem 购物车中添加商品 Item
func (s *Service) AddItem(userID uint, sku string, count int) error {
	currentProduct, err := s.productRepository.FindBySKU(sku) //根据商品唯一编码查找商品
	if err != nil {
		return err
	}
	currentCart, err := s.cartRepository.FindOrCreateByUserID(userID) //根据userid查找或者创建购物车
	if err != nil {
		return err
	}
	_, err = s.cartItemRepository.FindByID(currentProduct.ID, currentCart.ID) //根据商品id和购物车id查找购物车中所有商品
	if err == nil {
		//err 是空 说明找到了 所以返回商品已存在
		return ErrItemAlreadyExistInCart
	}
	if currentProduct.StockCount < count { //判断加入购物车的商品数量与库存数量对比 如果加入购物车数量大于库存数量 则返回库存不足
		return product.ErrProductStockIsNotEnough
	}
	if count < 0 {
		//这里判断加入购物车的数量 不能小于0 如果小于等于0 则返回商品数量不能是负数
		return ErrCountInvalid
	}
	err = s.cartItemRepository.Create(NewCartItem(currentProduct.ID, currentCart.ID, count))
	return err
}

// UpdateItem 更新购物车中的商品 item
func (s *Service) UpdateItem(userID uint, sku string, count int) error {
	currentProduct, err := s.productRepository.FindBySKU(sku) //根据商品唯一编号查找商品
	if err != nil {
		return err
	}
	cart, err := s.cartRepository.FindOrCreateByUserID(userID) //根据用户id查找或者创建购物车
	if err != nil {
		return err
	}
	item, err := s.cartItemRepository.FindByID(currentProduct.ID, cart.ID)
	if err != nil {
		return errors.New("item 购物车中商品不存在")
	}
	if currentProduct.StockCount+item.Count < count {
		return product.ErrProductStockIsNotEnough
	}
	item.Count = count
	err = s.cartItemRepository.Update(*item)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItems 查询 购物车中所有商品 items
func (s *Service) GetCartItems(userId uint) ([]Item, error) {
	cart, err := s.cartRepository.FindOrCreateByUserID(userId)
	if err != nil {
		return nil, err
	}
	items, err := s.cartItemRepository.GetItems(cart.ID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
