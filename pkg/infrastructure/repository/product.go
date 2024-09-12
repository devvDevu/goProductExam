package repository

import (
	"context"
	entity "goProductExam/pkg/entity"
  "log"
)

func (repo *PGRepo) CreateProduct(product entity.Product) (id int, err error) {
  ctx := context.Background()
  err = repo.pool.QueryRow(ctx, "insert into product (name, cost) values($1, $2) returning id", product.Name, product.Cost).Scan(&id)
  if err != nil {
    log.Fatal(err.Error())
  }

  return id, nil
}

func (repo *PGRepo) GetProducts() ([]entity.Product, error) {
  ctx := context.Background()
  var productSlice []entity.Product
  var product entity.Product  

  rows, err := repo.pool.Query(ctx, "select id, name, cost from product")
  if err != nil {
    log.Fatal(err.Error)
  }
  for rows.Next() {
    err = rows.Scan(
      &product.Id,
      &product.Name,
      &product.Cost,
    )
    if err != nil {
      log.Fatal(err.Error())
    }
    productSlice = append(productSlice, product)
  }
  return productSlice, nil
}

func (repo *PGRepo) UpdateCostProduct (product entity.Product) (updatedProduct entity.Product, err error) {
  ctx := context.Background()
  err = repo.pool.QueryRow(ctx, "update product set cost=$1 where id=$2 returning id, name, cost", product.Cost, product.Id).Scan(
    &updatedProduct.Id,
    &updatedProduct.Name,
    &updatedProduct.Cost,
  )
  if err != nil {
    log.Fatal(err.Error())
  }
  return updatedProduct, nil
}

func (repo *PGRepo) DeleteProduct (id int) (int, error) {
  ctx := context.Background()
  _, err := repo.pool.Exec(ctx, "delete from product where id=$1", id)
  if err != nil {
    log.Fatal(err.Error())
  }
  return id, nil
}
