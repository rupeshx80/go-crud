package models

type User struct {
	//gorm.Model  adds ID(autoincrement),createdAt,updatedAt,deletedAt . best practise
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
