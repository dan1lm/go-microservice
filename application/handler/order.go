package handler

import (
	"fmt"
	"net/http"

	"github.com/dan1lm/go-microservice.git/repository/order"
)

type Order struct{
	Repo *order.RedisRepo
}

func(o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get an order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update an order by ID")
}
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete an order by ID")
}