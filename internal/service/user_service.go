package service

import (
	"nectarpin/internal/models"
	"nectarpin/internal/repository"
	"nectarpin/internal/utils"
)

// UserService 用户业务逻辑层
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService 创建新的用户 Service 实例
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

/*
CreateUser 创建用户 [0.2][2025.12.07]

	@param user *models.User
	@return msg string 提示信息
	@return code int 状态码
*/
func (s *UserService) CreateUser(user *models.User) (msg string, code int) {
	//[1]数据校验 - 检查用户名和邮箱是否已存在
	msg, code, exist := s.ExistUserOrEmail(user.UserName, user.Email)
	if exist {
		return msg, code
	}

	//[2]密码加密
	msg, code, user.Password = utils.EncryptUserPassword(user.Password)
	if code != 200 {
		return msg, 500
	}

	//[3]写入数据库
	err := s.userRepo.Create(user)
	if err != nil {
		return "创建用户失败", 500
	}

	return "创建用户成功", 200
}

/*
ExistUserOrEmail 检测(用户/邮箱)是否存在 [0.2][2025.12.07]

	@param userName string 用户名(可选)
	@param email string 邮箱(可选)
	@return msg string 提示信息
	@return code int 状态码
	@return bool 检测结果
*/
func (s *UserService) ExistUserOrEmail(userName string, email string) (msg string, code int, exist bool) {
	// 检查用户名是否存在
	if userName != "" {
		user, err := s.userRepo.ExistsByUsername(userName)
		if err != nil {
			return "数据库查询失败", 500, false
		}
		if user != nil {
			// 如果同时传了邮箱，且邮箱属于该用户，则跳过邮箱检测
			if email != "" && user.Email == email {
				return "检测完成,无重复!", 200, false
			}
			return "用户名已存在!", 2001, true
		}
	}

	// 检查邮箱是否存在
	if email != "" {
		emailExists, err := s.userRepo.ExistsByEmail(email)
		if err != nil {
			return "数据库查询失败", 500, false
		}
		if emailExists {
			return "邮箱已存在!", 2002, true
		}
	}

	return "检测完成,无重复!", 200, false
}

/*
GetUserByID 根据ID获取用户信息

	@param id uint
	@return *models.User
	@return error
*/
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

/*
GetUserByUsername 根据用户名获取用户信息

	@param userName string
	@return *models.User
	@return error
*/
func (s *UserService) GetUserByUsername(userName string) (*models.User, error) {
	return s.userRepo.FindByUsername(userName)
}

/*
UpdateUser 更新用户信息

	@param user *models.User
	@return msg string 提示信息
	@return code int 状态码
*/
func (s *UserService) UpdateUser(user *models.User) (msg string, code int) {
	// 检查邮箱是否被其他用户占用
	if user.Email != "" {
		existUser, err := s.userRepo.ExistsByUsername(user.UserName)
		if err != nil {
			return "数据库查询失败", 500
		}
		// 如果邮箱被占用且不是当前用户
		if existUser != nil && existUser.ID != user.ID {
			emailExists, err := s.userRepo.ExistsByEmail(user.Email)
			if err != nil {
				return "数据库查询失败", 500
			}
			if emailExists {
				return "邮箱已被使用!", 2002
			}
		}
	}

	err := s.userRepo.Update(user)
	if err != nil {
		return "更新用户失败", 500
	}
	return "更新用户成功", 200
}

/*
DeleteUser 删除用户

	@param id uint
	@return msg string 提示信息
	@return code int 状态码
*/
func (s *UserService) DeleteUser(id uint) (msg string, code int) {
	err := s.userRepo.Delete(id)
	if err != nil {
		return "删除用户失败", 500
	}
	return "删除用户成功", 200
}
