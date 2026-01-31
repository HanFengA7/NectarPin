package repository

import "NectarPin/api/models"

type UserRepository interface {
}

// CreateUser
// @version 1.0.260131
// @desc 创建用户
// @param user [model.User]
// @return error
func CreateUser(user *models.User) error {
	return nil
}
