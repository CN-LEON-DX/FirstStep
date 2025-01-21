package storage

import (
	"awesomeProject/module/item/model"
	"context"
)

func (s *sqlStore) DeleteItem(tcx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
