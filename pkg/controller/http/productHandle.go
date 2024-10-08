package api

import (
	"encoding/json"
	entity "goProductExam/pkg/entity"
	"log"
	"net/http"
	"strconv"
)


func (api *api) product (w http.ResponseWriter, r *http.Request) {
  switch r.Method{
  case http.MethodGet:
    api.mu.Lock()
    products, err := api.p.ProductGet()
    api.mu.Unlock()
    if err != nil {
      log.Fatal("GET db error")
    }
    err = json.NewEncoder(w).Encode(products)
    if err != nil {
      log.Fatal("GET encode error")
    }
  case http.MethodPost:
    product := entity.Product{}
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
      log.Fatal("POST decode error")
    }
    api.mu.Lock()
    id, err := api.p.ProductPost(product)
    api.mu.Unlock()
    if err != nil {
      log.Fatal("GET db error")
    }

    err = json.NewEncoder(w).Encode(id)
    if err != nil {
      log.Fatal("POST encode error")
    }
  case http.MethodPut:
    product := entity.Product{}
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
      log.Fatal("PUT decode error")
    }
    api.mu.Lock()
    product, err = api.p.ProductPut(product)
    api.mu.Unlock()
    if err != nil {
      log.Fatal("PUT db error")
    }
    err = json.NewEncoder(w).Encode(product)
    if err != nil {
      log.Fatal("PUT encode error")
    }
  case http.MethodDelete:
    id := r.URL.Query().Get("id")
    convId, _ := strconv.Atoi(id)

    api.mu.Lock()
    convId, err := api.p.ProductDelete(convId)
    api.mu.Unlock()
    if err != nil {
      log.Fatal("DELETE db error")
    }
    
    err = json.NewEncoder(w).Encode(convId)
    if err != nil {
      log.Fatal("DELETE encode error")
    }
  }
} 
