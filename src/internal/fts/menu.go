package fts

import (
	"menu/pkg/manticore"
)

type MenuSearch struct {
	client *manticore.ManticoreClient
}

func NewMenuSearch(client *manticore.ManticoreClient) *MenuSearch {
	return &MenuSearch{client}
}

func (s *MenuSearch) SearchItems(query string) ([]int, error) {
	rows, err := s.client.Queryx(`SELECT id FROM menu_items_index WHERE MATCH(?)`, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itemIDs := make([]int, 0)
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		itemIDs = append(itemIDs, id)
	}
	return itemIDs, nil
}
