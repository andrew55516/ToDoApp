package main

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type Todo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"isDone"`
}

type TodoManager struct {
	todos []Todo
	m     sync.Mutex
}

type CreateTodoRequest struct {
	Title string `json:"title"`
}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		todos: make([]Todo, 0),
		m:     sync.Mutex{},
	}
}

func (tm *TodoManager) GetAll() []Todo {
	return tm.todos
}

func (tm *TodoManager) Create(request CreateTodoRequest) Todo {
	tm.m.Lock()
	defer tm.m.Unlock()

	newTodo := Todo{
		ID:         strconv.FormatInt(time.Now().UnixMilli(), 10),
		Title:      request.Title,
		IsComplete: false,
	}

	tm.todos = append(tm.todos, newTodo)

	return newTodo
}

func (tm *TodoManager) Complete(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	var todo *Todo
	index := -1

	for i, t := range tm.todos {
		if t.ID == ID {
			todo = &t
			index = i
			break
		}
	}

	if todo == nil {
		return errors.New("todo not found")
	}

	if todo.IsComplete {
		return errors.New("todo is already complete")
	}

	tm.todos[index].IsComplete = true

	return nil
}

func (tm *TodoManager) Remove(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	index := -1

	for i, t := range tm.todos {
		if t.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("todo not found")
	}

	tm.todos = append(tm.todos[:index], tm.todos[index+1:]...)

	return nil
}