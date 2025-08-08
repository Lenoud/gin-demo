package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Lenoud/gin-demo/model"
	"github.com/Lenoud/gin-demo/utils"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func RegisterUser(user *model.UserJson) (*model.UserJson, error) {
	zap.L().Info("尝试注册用户", zap.String("username", user.Username), zap.String("email", user.Email))

	existingUser, err := model.GetUserByUsername(user.Username)
	if err != nil {
		zap.L().Error("查询用户名失败", zap.String("username", user.Username), zap.Error(err))
		return nil, err
	}
	if existingUser != nil {
		zap.L().Warn("用户名已存在", zap.String("username", user.Username))
		return nil, errors.New("用户名已存在")
	}

	// user.IsAdmin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := model.CreateUser(user); err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			zap.L().Warn("邮箱已存在", zap.String("email", user.Email))
			return nil, fmt.Errorf("邮箱已经存在！")
		}
		zap.L().Error("创建用户失败", zap.String("username", user.Username), zap.Error(err))
		return nil, fmt.Errorf("发生错误: %w", err)
	}

	zap.L().Info("用户注册成功", zap.Uint64("userId", user.Id), zap.String("username", user.Username))
	return user, nil
}

func LoginUser(username, password string) (string, *model.UserJson, error) {

	zap.L().Info("尝试登录用户", zap.String("username", username))

	user, err := model.GetUserByUsername(username)
	if err != nil {
		zap.L().Error("查询用户失败", zap.String("username", username), zap.Error(err))
		return "", nil, err
	}
	if user == nil || user.Password != password {
		zap.L().Warn("用户名或密码错误", zap.String("username", username))
		return "", nil, errors.New("用户名或密码错误")
	}

	token, err := utils.GenerateToken(user.Id, user.IsAdmin)
	if err != nil {
		zap.L().Error("生成 Token 失败", zap.Uint64("userId", user.Id), zap.Error(err))
		return "", nil, err
	}

	zap.L().Info("用户登录成功", zap.Uint64("userId", user.Id), zap.String("username", user.Username))
	return token, user, nil
}
