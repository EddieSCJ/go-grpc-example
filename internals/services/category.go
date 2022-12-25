package services

import (
	"context"
	"github.com/EddieSCJ/go-grpc-example/internals/database"
	"github.com/EddieSCJ/go-grpc-example/internals/pb"
	"io"
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

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categoryList := make([]*pb.Category, 0)
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CategoryResponse{Category: categoryList})
		}

		result, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categoryItem := &pb.Category{
			Id:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		}
		categoryList = append(categoryList, categoryItem)
	}
}

func (c *CategoryService) CreateCategoryStreamBoth(stream pb.CategoryService_CreateCategoryStreamBothServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		result, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categoryItem := &pb.Category{
			Id:          result.ID,
			Name:        result.Name,
			Description: result.Description,
		}

		err = stream.Send(categoryItem)
		if err != nil {
			return err
		}
	}
}
