package service

import (
	"time"
	"web-demo/cgo/dao"
	"web-demo/cgo/entity"
)

type UserService struct {
}

var userDao = new(dao.UserDao)

func (p *UserService) Insert(username, password string) int64 {
	return userDao.Insert(&entity.User{Username: username, Password: password, CreateTime: time.Now()})
}

func (p *UserService) SelectUserByName(username string) []entity.User {
	return userDao.SelectUserByName(username)
}

func (p *UserService) SelectAllUser() []entity.User {
	return userDao.SelectAllUser()
}
