package storage

import (
	"awesomeProject/module/item/model"
	"context"
)

func (s *sqlStore) CreateItem(tcx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
