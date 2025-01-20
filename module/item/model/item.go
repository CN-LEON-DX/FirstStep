package model

import (
	"awesomeProject/common"
	"errors"
	"strings"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
)

type TodoItem struct {
	common.SQLModel        // embed struct
	Title           string `json:"string" gorm:"column:title"`
	Description     string `json:"description" gorm:"column:description"`
	Status          string `json:"status" gorm:"column:status"`
}
type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id"` // tag name json, json encode
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}
	return nil
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItem) TableName() string         { return "todo_items" }
func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }
func (TodoItemUpdate) TableName() string   { return TodoItem{}.TableName() }
