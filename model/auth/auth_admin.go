package auth

import "github.com/yeyudekuangxiang/imagedesign/model"

type Admin struct {
	ID int
}

func (au Admin) Valid() error {
	return nil
}

type OldAdmin struct {
	UserId model.StrToInt `json:"userId"`
}
