package biz

// handler => biz [=>repository => Storage]
// handler => parse request check => convert json => format in business
// biz     => using input by handler => combine the requirement => send to repository
// repo    => get info in some where in db, mysql, mongodb, ....
// storage => communicate with db then get the data return by request !
type createItemBiz struct {
}
