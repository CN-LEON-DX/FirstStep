package biz

import (
	"awesomeProject/module/item/model"
	"context"
)

// handler => biz [=>repository => Storage]
// handler => parse request check => convert json => format in business
// biz     => using input by handler => combine the requirement => send to repository
// repo    => get info in some where in db, mysql, mongodb, ....
// storage => communicate with db then get the data return by request !

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}
type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
