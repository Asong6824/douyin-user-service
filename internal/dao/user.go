package dao

import (
	"github.com/Asong6824/douyin-user-service/internal/model"
	//"github.com/Asong6824/douyin-user-service/internal/cache"
	"github.com/Asong6824/douyin-user-service/package/utils"
)

func (d *Dao) Register(username string, password string) (int64, error) {
	user := model.User{
		Name: username,
	}
	isUserExists, err := user.IsUserExists(d.engine)
	if err != nil && isUserExists {
		return 0, err
	}
	id, err := utils.GenerateSnowflakeID()
	if err != nil {
		return 0, err
	}
	user.Id = id
	user.Password = password
	err = user.Register(d.engine)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Dao) Login(username string, password string) (int64, error) {
	user := model.User{
		Name: username,
		Password: password,
	}
	return user.Login(d.engine)
}

// func (d *Dao) GetUser(userid uint32) (*model.User, error) {
// 	user := model.User{
// 		UserID: userid,
// 	}
// 	return user.GetUser(d.engine)
// }

// //ActionType 0 minus, 1 plus
// func (d *Dao) ModifyFollowingCount(userid uint32, ActionType int) error {
// 	user := model.User{
// 		UserID: userid,
// 	}
// 	if ActionType == 1 {
// 		return user.PlusFollowingCount(d.engine)
// 	} else {
// 		return user.MinusFollowingCount(d.engine)
// 	}
// }

// func (d *Dao) ModifyFollowersCount(userid uint32, ActionType int) error {
// 	user := model.User{
// 		UserID: userid,
// 	}
// 	if ActionType == 1 {
// 		return user.PlusFollowersCount(d.engine)
// 	} else {
// 		return user.MinusFollowersCount(d.engine)
// 	}
// }