package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lolitsgab/dist-kv-store/application/model"
	"github.com/lolitsgab/dist-kv-store/repository/order"
)

type Order struct {
	Repo *order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
	var body struct {
		CustomerID uuid.UUID        `json:"customer_id"`
		LineItems  []model.LineItem `json:"line_items"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, fmt.Errorf("failed to read body: %w", err).Error(), http.StatusBadRequest)
		return
	}
	now := time.Now()
	order := model.Order{
		OrderID:    rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}
	err := o.Repo.Insert(r.Context(), order)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to insert order: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(order)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to marshal order: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
	w.WriteHeader(http.StatusOK)
	return
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
