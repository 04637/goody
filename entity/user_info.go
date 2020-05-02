package entity

type UserInfo struct {
	Id int `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Age string `gorm:"size:11;DEFAULT NULL" json:"age"`
	Name string `gorm:"size:255;DEFAULT NULL" json:"name"`
	Bio string `gorm:"size:255;DEFAULT NULL" json:"bio"`
}
