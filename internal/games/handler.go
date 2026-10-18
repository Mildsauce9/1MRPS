package games

import (
	"log"
	"net/http"

	"github.com/Mildsauce9/1MRPS/internal/json"
)

type handler struct {
	processor GameProcessor
}

func NewHandler(processor GameProcessor) *handler {
	return &handler{
		processor: processor,
	}
}

func (h *handler) SelectGames(w http.ResponseWriter, r *http.Request) {

	games, err := h.processor.ListGames(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, games)
}
