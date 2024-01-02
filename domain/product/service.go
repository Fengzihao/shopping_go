package product

import "shopping_go/utils/pagination"

type Service struct {
	ProductRepository Repository
}

// NewService 实例化service
func NewService(repository Repository) *Service {
	repository.Migration() //生成表
	return &Service{
		ProductRepository: repository,
	}
}

// GetAll 获得所有商品 分页
func (s *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	all, i := s.ProductRepository.GetAll(page.Page, page.PageSize)
	page.Items = all
	page.TotalCount = i
	return page
}

// CreateProduct 创建商品
func (s *Service) CreateProduct(name string, desc string, count int, price float32, cid uint) error {
	product := NewProduct(name, desc, count, price, cid)
	err := s.ProductRepository.Create(product)
	return err
}

// DeleteProduct 删除商品
func (s *Service) DeleteProduct(sku string) error {
	err := s.ProductRepository.Delete(sku)
	return err
}

// UpdateProduct 更新商品
func (s *Service) UpdateProduct(product *Product) error {
	err := s.ProductRepository.Update(*product)
	return err
}

// SearchProduct 查找商品
func (s *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	byString, i := s.ProductRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = byString
	page.TotalCount = i
	return page
}
