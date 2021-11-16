package entity

type User struct {
	ID       int    `gorm:"primary_key;column:id" json:"id"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Avatar   string `gorm:"avatar" json:"avatar"`
	Guid     string `gorm:"guid" json:"guid"`
}
