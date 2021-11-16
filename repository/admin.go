package repository

import (
	"github.com/yeyudekuangxiang/imagedesign/core/app"
	"github.com/yeyudekuangxiang/imagedesign/model/entity"
	"gorm.io/gorm"
)

var DefaultAdminRepository IAdminRepository = NewAdminRepository()

type IAdminRepository interface {
	//根据管理员id获取管理员信息
	GetAdminById(int) (*entity.Admin, error)
}

func NewAdminRepository() AdminRepository {
	return AdminRepository{}
}

type AdminRepository struct {
}

func (a AdminRepository) GetAdminById(id int) (*entity.Admin, error) {
	var admin entity.Admin
	if err := app.DB.First(&admin, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}
