package handler

import (
	"category/domain/model"
	"category/domain/service"
	category "category/proto/category"
	"context"
)

type Category struct{
     CategoryDataService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest,
	response *category.CreateCategoryResponse) error {
	category := &model.Category{}

}
