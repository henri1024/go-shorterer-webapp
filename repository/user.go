package repository

import "go-shorterer/model"

func (db *dbm) SaveUser(user *model.User) error {
	err := db.PsqlDB.Model(&model.User{}).Save(user).Error
	return err
}

func (db *dbm) CheckAPIKey(apiKey string) bool {
	user := &model.User{}
	err := db.PsqlDB.Model(&model.User{}).Where(&model.User{APIKEY: apiKey}).Find(user).Error
	if err == nil && user.APIKEY != "" && user.Email != "" {
		return true
	}
	return false
}
