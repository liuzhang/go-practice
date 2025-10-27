package storage

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"todo/model"
)

type YamlStorage struct {
	filePath string
}

func NewYamlStorage(filePath string) *YamlStorage {
	return &YamlStorage{filePath: filePath}
}

func (s *YamlStorage) Load() ([]model.Todo, error) {
	content, err := os.ReadFile(s.filePath)

	if err != nil {
		return nil, err
	}
	var todo []model.Todo

	err = yaml.Unmarshal(content, &todo)
	if err != nil {
		return nil, err
	}
	return todo, nil

}

func (s *YamlStorage) List() ([]model.Todo, error) {
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

func (s *YamlStorage) Add(todo *model.Todo) error {
	todos, _ := s.Load()
	todo.Id = len(todos) + 1
	todos = append(todos, *todo)
	content, err := yaml.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, content, 0644)
}

func (s *YamlStorage) Done(id int) error {
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

	content, err := yaml.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, content, 0644)
}
