package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	fmt.Println("List all orders:", r.URL.Query())
	cursorstr := r.URL.Query().Get("cursor")
	cursor, err := strconv.ParseUint(cursorstr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to parse cursor: %w", err).Error(), http.StatusBadRequest)
	}
	fmt.Println("Cursor:", cursor)
	pageresult, err := o.Repo.FindAll(r.Context(), order.FindAllPage{
		Size:   50,
		Offset: cursor,
	})

	if err != nil {
		http.Error(w, fmt.Errorf("failed to find all orders: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(pageresult)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to marshal order: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
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
