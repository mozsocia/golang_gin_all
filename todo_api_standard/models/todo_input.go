package models

type CreateTodoInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
