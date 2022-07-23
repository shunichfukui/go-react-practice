package usecases_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/shunichfukui/go-react-practice/backend/entities"
	"github.com/shunichfukui/go-react-practice/backend/usecases"
)

var dummyTodos = []entities.Todo{
	{
		Title:       "todo 1",
		Description: "description of todo1",
		IsCompleted: false,
	},
	{
		Title:       "todo 2",
		Description: "description of todo2",
		IsCompleted: true,
	},
	{
		Title:       "todo 3",
		Description: "description of todo3",
		IsCompleted: true,
	},
}

type MockTodosRepo struct {
}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil
}

type BadTodoRepo struct{}

func (BadTodoRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("something went wrong")
}

func TestGetTodos(t *testing.T) {
	// テスト
	t.Run("Returens ErrInternal when TodosRepository returns err", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(BadTodoRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(usecases.ErrInternal)

		if todos != nil {
			t.Fatalf("Expected todos to be nil: Got: %v", todos)
		}
	})

	t.Run("Returns todos from TodoRepository", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(MockTodosRepo)

		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(nil)
		expect(todos).ToEqual(dummyTodos)
	})
}
