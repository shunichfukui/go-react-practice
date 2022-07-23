package ui

import "github.com/shunichfukui/go-react-practice/backend/entities"

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
