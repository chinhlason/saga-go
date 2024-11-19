package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	r IRepository
}

func NewHandler(r IRepository) *Handler {
	return &Handler{r: r}
}

func (h *Handler) Insert(res http.ResponseWriter, req *http.Request) {
	err := h.r.Insert(req.Context())
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ProduceMessage("localhost:29092", "inventory", "ORDER CREATED")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{"message": "success"}`))
}

func (h *Handler) Get(res http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	order, err := h.r.Get(req.Context(), id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(order); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Update(res http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	status := req.URL.Query().Get("status")
	err := h.r.Update(req.Context(), id, status)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"message": "success"}`))
}

func (h *Handler) Write(res http.ResponseWriter, req *http.Request) {
	order := &Order{}
	if err := json.NewDecoder(req.Body).Decode(order); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	err := ProduceMessage("localhost:29092", "test", order.Status)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{"message": "success"}`))
}

func (h *Handler) OnMessage() {
	err := ConsumeMessages("localhost:29092", "order", "group1")
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
	}
}
