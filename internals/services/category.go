package services

import (
	"context"
	"github.com/EddieSCJ/go-grpc-example/internals/database"
	"github.com/EddieSCJ/go-grpc-example/internals/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	pbCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	categorySlice := make([]*pb.Category, 0, 1)
	categorySlice = append(categorySlice, pbCategory)

	response := &pb.CategoryResponse{
		Category: categorySlice,
	}

	return response, nil
}

func (c *CategoryService) GetCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryResponse, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	categorySlice := make([]*pb.Category, 0, len(categories))
	for _, category := range categories {
		pbCategory := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categorySlice = append(categorySlice, pbCategory)
	}

	response := &pb.CategoryResponse{
		Category: categorySlice,
	}
	return response, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.FindCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	pbCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	categorySlice := make([]*pb.Category, 0, 1)
	categorySlice = append(categorySlice, pbCategory)

	response := &pb.CategoryResponse{
		Category: categorySlice,
	}

	return response, nil
}
