package storage

import "todo/model"

type Operator interface {
	Load() ([]model.Todo, error)
	List() ([]model.Todo, error)
	Add(todo *model.Todo) error
	Done(id int) error
}
