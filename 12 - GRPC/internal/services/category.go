package services

import (
	"context"
	"renebizelli/grpc/internal/database"
	"renebizelli/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer // necessário
	CategoryDB                            *database.Category
}

func NewCategoryService(db *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

// essa assinaqtura foi obtida do arquivo /internal/pb/course_category_grpc.pb.go
//
//	type CategoryServiceServer interface {
//		CreateCategory(context.Context, *CreateCategoryRequest) (*CategoryResponse, error)  <-----
//		mustEmbedUnimplementedCategoryServiceServer()
//	}
func (s *CategoryService) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {

	category, err := s.CategoryDB.Create(request.Name, request.Description)

	if err != nil {
		panic(err)
	}

	categoryResponse := pb.CategoryResponse{
		Category: &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	return &categoryResponse, nil

}
