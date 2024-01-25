package models

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	ArticleID uint
	Article   Article
	Content   string `gorm:"size:100"`
}
