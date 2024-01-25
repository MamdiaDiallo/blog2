package models

type Article struct {
    ID      uint   `gorm:"primaryKey"`
    UserID  uint
    Content string
    Likes   int
}