package storage

import (
	"encoding/json"
	"errors"
	"os"
	"todo/model"
)

type JsonStorage struct {
	filePath string
}

func NewJsonStorage(filePath string) *JsonStorage {
	return &JsonStorage{filePath: filePath}
}

func (s *JsonStorage) Load() ([]model.Todo, error) {
	content, err := os.ReadFile(s.filePath)

	if err != nil {
		return nil, err
	}
	var todo []model.Todo

	err = json.Unmarshal(content, &todo)
	if err != nil {
		return nil, err
	}
	return todo, nil

}

func (s *JsonStorage) List() ([]model.Todo, error) {
	todos, _ := s.Load()
	var newTodos []model.Todo

	for _, todo := range todos {
		if !todo.Done {
			newTodos = append(newTodos, todo)
		}
	}
	if len(newTodos) == 0 {
		return nil, errors.New("no todos, please add")
	}
	return newTodos, nil

}

func (s *JsonStorage) Add(todo *model.Todo) error {
	todos, _ := s.Load()
	todo.Id = len(todos) + 1
	todos = append(todos, *todo)
	content, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, content, 0644)
}

func (s *JsonStorage) Done(id int) error {
	todos, _ := s.List()
	flag := false
	for i, todo := range todos {
		if todo.Id == id {
			flag = true
			todos[i].Done = true
		}
	}
	if !flag {
		return errors.New("todo not found")
	}

	content, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, content, 0644)
}
