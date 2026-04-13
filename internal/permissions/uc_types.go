package permissions

type CreateInput struct {
	Name string `json:"name"`
}

type UpdateInput struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type ListInput struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
