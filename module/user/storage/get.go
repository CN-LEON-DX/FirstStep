package storage

import (
	"awesomeProject/common"
	"awesomeProject/module/user/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(cxt context.Context, conditions map[string]interface{}, moreInfor ...string) (*model.User, error) {
	db := s.db.Table(model.UserLogin{}.TableName())

	//for i := range moreInfor {
	//	db = db.Preload(moreInfor[i])
	//}

	var user model.User
	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFoundError
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil

}
