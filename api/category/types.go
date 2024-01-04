package category

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
type CreateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
