package responses

type ListResponse[T any] struct {
	Count int `json:"count"`
	Items []T `json:"items"`
}

func NewListResponse[T any](items []T) ListResponse[T] {
	return ListResponse[T]{
		Count: len(items),
		Items: items,
	}
}
