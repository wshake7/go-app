package model

type User struct {
	Id       int64
	NickName string `xorm:"varchar(56) notnull comment('昵称')"`
	Account  string `xorm:"varchar(32) notnull unique comment('账号')"`
	Password string `xorm:"varchar(255) notnull comment('密码')"`
	Status   int    `xorm:"default(0) comment('状态')"`
}
