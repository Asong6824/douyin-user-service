package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id              int64  
	Name            string 
	Password        string
	FollowCount     int64  
	FollowerCount   int64     
	Avatar          string 
	BackgroundImage string 
	Signature       string 
	TotalFavorited  int64  
	WorkCount       int64  
	FavoriteCount  int64  
}

func (u User) IsUserExists(db *gorm.DB) (bool, error) {
    var user User

    err := db.Select("id").Where("name = ?", u.Name).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return false, nil
        }
        // 其他错误
        return false, err
    }

    // 找到了用户，返回true
    return true, nil
}

func (u User) Register(db *gorm.DB) error {
	if err := db.Table("user").Model(&User{}).Create(&u).Error; err != nil {
		return err
	}
	return nil
}

// func (u User) Login(db *gorm.DB) (int64, error) {
// 	var user User

// 	// 使用 GORM 的 Where 方法查询数据库中的记录
// 	result := db.Table("users").
// 		Where("user_name = ? AND password = ?", u.UserName, u.Password).
// 		First(&user)

// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return 0, result.Error
// 		}
// 		return 0, result.Error
// 	}

// 	return user.UserID, nil
// }

// func (u User) GetUser(db *gorm.DB) (*User, error) {
// 	var user User

// 	result := db.Table("users").Where("user_id = ?", u.UserID).First(&user)

// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, result.Error
// 		}
// 		return nil, result.Error
// 	}
// 	return &user, nil
// }

// func (u User) PlusFollowingCount(db *gorm.DB) error {
// 	err := db.Table("users").
// 		Where("user_id = ?", u.UserID).
// 		Model(&u).
// 		UpdateColumn("following_count", gorm.Expr("following_count + ?", 1)).Error 
// 	if err != nil { 
// 		return err 
// 	} 
// 	return nil
// }

// func (u User) MinusFollowingCount(db *gorm.DB) error {
// 	err := db.Table("users").
// 		Where("user_id = ?", u.UserID).
// 		Model(&u).
// 		UpdateColumn("following_count", gorm.Expr("following_count - ?", 1)).Error 
// 	if err != nil { 
// 		return err 
// 	} 
// 	return nil
// }

// func (u User) PlusFollowersCount(db *gorm.DB) error {
// 	err := db.Table("users").
// 		Where("user_id = ?", u.UserID).
// 		Model(&u).
// 		UpdateColumn("followers_count", gorm.Expr("followers_count + ?", 1)).Error 
// 	if err != nil { 
// 		return err 
// 	} 
// 	return nil
// }

// func (u User) MinusFollowersCount(db *gorm.DB) error {
// 	err := db.Table("users").
// 		Where("user_id = ?", u.UserID).
// 		Model(&u).
// 		UpdateColumn("followers_count", gorm.Expr("followers_count - ?", 1)).Error 
// 	if err != nil { 
// 		return err 
// 	} 
// 	return nil
// }