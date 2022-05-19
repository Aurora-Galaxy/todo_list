package serializer

import "todo_list/model"

type User struct {
	ID       uint   `json:"id"form:"id"`
	UserName string `json:"user_name"form:"username"`
	Status   string `json:"status"form:"status"`
	CreateAt int64  `json:"create_at"form:"createat"`
}

// BuildUser 序列化用户
func BuildUser(user model.User)User{
	return User{
		ID: user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}