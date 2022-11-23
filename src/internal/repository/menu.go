package repository

import (
	"menu/internal/models"

	"github.com/jmoiron/sqlx"
)

type MenuItem struct {
}

type MenuRespository struct {
	db *sqlx.DB
}

func NewMenuRespository(db *sqlx.DB) *MenuRespository {
	return &MenuRespository{db}
}

func (r *MenuRespository) SearchItems(query string) ([]models.MenuItem, error) {
	sqlQuery := `SELECT * FROM menu_items WHERE (title ILIKE :query OR description ILIKE :query)`
	queryArgs := map[string]interface{}{
		"query": "%" + query + "%",
	}
	rows, err := r.db.NamedQuery(sqlQuery, queryArgs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]models.MenuItem, 0)
	for rows.Next() {
		var item models.MenuItem
		err := rows.StructScan(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *MenuRespository) GetByIDs(ids []int) ([]models.MenuItem, error) {
	sqlQuery := `SELECT * FROM menu_items WHERE id IN (:product_ids)`
	queryArgs := map[string]interface{}{
		"product_ids": ids,
	}
	query, args, err := sqlx.Named(sqlQuery, queryArgs)
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(r.db.Rebind(query), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]models.MenuItem, 0)
	for rows.Next() {
		var item models.MenuItem
		err := rows.StructScan(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
