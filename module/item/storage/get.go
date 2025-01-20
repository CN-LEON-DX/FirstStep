package storage

import (
	"awesomeProject/module/item/model"
	"context"
)

func (s *sqlStore) GetItem(tcx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	// get data
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
