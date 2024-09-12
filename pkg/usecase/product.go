package usecase

import (
  entity "goProductExam/pkg/entity"
)


type ProductInterface interface{
  ProductGet() ([]entity.Product, error)
  ProductPost(entity.Product) (int, error)
  ProductPut(entity.Product) (entity.Product, error)
  ProductDelete(int) (int, error)
}

type ProductRepo interface{
  GetProducts() ([]entity.Product, error)
  CreateProduct(entity.Product) (int, error)
  UpdateCostProduct(entity.Product) (entity.Product, error)
  DeleteProduct(int) (int, error)
}

type ProductUseCase struct {
  repo ProductRepo
}

func New (repo ProductRepo) *ProductUseCase{
  return &ProductUseCase{repo: repo}
}

func (us *ProductUseCase) ProductGet () ([]entity.Product, error) {
  entitiesSlice, err := us.repo.GetProducts()
  return entitiesSlice, err
}

func (us *ProductUseCase) ProductPost (entity entity.Product) (id int, err error) {
  id, err = us.repo.CreateProduct(entity)
  return id, err
}

func (us *ProductUseCase) ProductPut (entity entity.Product) (entity.Product, error) {
  entity, err := us.repo.UpdateCostProduct(entity)
  return entity, err
}

func (us *ProductUseCase) ProductDelete (id int) (int, error) {
  id, err := us.repo.DeleteProduct(id)
  return id, err
}
