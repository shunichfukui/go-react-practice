package usecases

import "github.com/shunichfukui/go-react-practice/backend/entities"

type TodoRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
