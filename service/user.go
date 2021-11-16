package service

import (
	"github.com/yeyudekuangxiang/imagedesign/internal/util"
	"github.com/yeyudekuangxiang/imagedesign/model/auth"
	"github.com/yeyudekuangxiang/imagedesign/model/entity"
	"github.com/yeyudekuangxiang/imagedesign/repository"
)

var DefaultUserService = NewUserService(repository.DefaultUserRepository)

type Iu interface {
	//根据token获取用户
	GetUserByToken(string) (*entity.User, error)
	//根据用户id获取用户信息
	GetUserById(int) (*entity.User, error)
}

func NewUserService(r repository.IUserRepository) UserService {
	return UserService{
		r: r,
	}
}

type UserService struct {
	r repository.IUserRepository
}

func (u UserService) GetUserById(id int) (*entity.User, error) {
	return u.r.GetUserById(id)
}

func (u UserService) GetUserByToken(token string) (*entity.User, error) {
	var authUser auth.User
	err := util.ParseToken(token, &authUser)
	if err != nil {
		return nil, err
	}
	return u.r.GetUserByGuid(authUser.Guid)
}
