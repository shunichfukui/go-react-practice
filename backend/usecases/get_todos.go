package usecases

import "github.com/shunichfukui/go-react-practice/backend/entities"

func GetTodos(repo TodoRepository) ([]entities.Todo, error) {
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}
	return todos, nil
}
