package repository

type CreateOptions struct {
	Name string
}

type UpdateOptions struct {
	ID   string
	Name string
}

type ListOptions struct {
	Limit  int32
	Offset int32
}
