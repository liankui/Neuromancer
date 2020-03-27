package model

import (
	"github.com/yueekee/Neuromancer/initDB"
	"log"
)

type UserModel struct {
	Id       int64  `from:"id"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
}

func (user *UserModel) Save() int64 {
	result, err := initDB.Db.Exec("insert into ginhello.user (email, password) values (?, ?)", user.Email, user.Password)
	if err != nil {
		log.Panicln("insert user err:", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id err:", err.Error())
	}
	return id
}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * form user where email = ?", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password)
	if e != nil {
		log.Panicln(e)
	}
	return u
}
