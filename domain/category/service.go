package category

import (
	"mime/multipart"
	"shopping_go/utils/csv_helper"
	"shopping_go/utils/pagination"
)

type Service struct {
	r Repository
}

func NewCategoryService(r Repository) *Service {
	r.Migration()        //创建商品分类表
	r.InsertSampleData() //插入商品分类测试数据
	return &Service{
		r: r,
	}
}

// Create 新增商品分类数据
func (s *Service) Create(category *Category) error {
	name := s.r.GetByName(category.Name)
	if len(name) > 0 {
		return ErrCategoryExistWithName
	}
	err := s.r.Create(category)
	if err != nil {
		return err
	}
	return nil
}

// BulkCreate 批量传教商品分类数据（csv导入）
func (s *Service) BulkCreate(header *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	csv, err := csv_helper.ReadCsv(header)
	if err != nil {
		return 0, err
	}
	//遍历读取的csv数据
	for _, strings := range csv {
		categories = append(categories, NewCategory(strings[0], strings[1]))
	}
	count, err := s.r.BulkCreate(categories)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetAll 分页获取商品分类数据
func (s *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	all, count := s.r.GetAll(page.Page, page.PageSize)
	page.Items = all
	page.TotalCount = count
	return page
}
