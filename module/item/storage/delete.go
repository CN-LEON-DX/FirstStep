package storage

import (
	"context"
)

func (s *sqlStore) DeleteItem(tcx context.Context, cond map[string]interface{}) error {
	deleteStatus := "Deleted"

	if err := s.db.Where(cond).Updates(map[string]interface{}{
		"status": deleteStatus,
	}).Error; err != nil {
		return err
	}
	return nil
}
