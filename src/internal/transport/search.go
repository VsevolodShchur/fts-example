package transport

import (
	"errors"
	"menu/internal/models"
	"menu/pkg/httputils"
	"menu/pkg/responses"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type menuSearch interface {
	SearchItems(str string) (ids []int, err error)
}

type menuRepository interface {
	SearchItems(str string) ([]models.MenuItem, error)
	GetByIDs(ids []int) ([]models.MenuItem, error)
}

type MenuSearchHandler struct {
	repository menuRepository
	search     menuSearch
}

func NewMenuSearchHandler(repo menuRepository, search menuSearch) *MenuSearchHandler {
	return &MenuSearchHandler{repo, search}
}

func (s *MenuSearchHandler) searchDB(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		err := errors.New("empty query")
		httputils.RespondJSON(w, http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	items, err := s.repository.SearchItems(query)

	if err != nil {
		httputils.RespondJSON(w, http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}
	httputils.RespondJSON(w, http.StatusOK, responses.NewListResponse(items))
}

func (s *MenuSearchHandler) searchFTS(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		err := errors.New("empty query")
		httputils.RespondJSON(w, http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	itemIDs, err := s.search.SearchItems(query)
	if err != nil {
		httputils.RespondJSON(w, http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}
	items, err := s.repository.GetByIDs(itemIDs)
	if err != nil {
		httputils.RespondJSON(w, http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}
	httputils.RespondJSON(w, http.StatusOK, responses.NewListResponse(items))
}

func (s *MenuSearchHandler) GetRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/db", s.searchDB)
	r.Get("/fts", s.searchFTS)

	return r
}
