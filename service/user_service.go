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

func GetAllUsers() ([]model.UserJson, error) {
	zap.L().Info("开始获取所有用户列表")
	// 这里可以做额外业务逻辑处理，比如过滤、排序、缓存等
	users, err := model.ListUsers()
	if err != nil {
		zap.L().Error("获取用户列表失败", zap.Error(err))
		return nil, err
	}

	zap.L().Info("获取用户列表成功", zap.Int("count", len(users)))
	return users, nil
}

func RegisterUser(user *model.UserJson) (*model.UserJson, error) {
	zap.L().Info("尝试注册用户", zap.String("username", user.Username), zap.String("email", user.Email))

	existingUser, err := model.GetUserByUsername(user.Username)
	if err != nil {
		zap.L().Error("查询用户名失败", zap.String("username", user.Username), zap.String("error", err.Error()))
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
		zap.L().Error("创建用户失败", zap.String("username", user.Username), zap.String("error", err.Error()))
		return nil, fmt.Errorf("发生错误: %w", err)
	}

	zap.L().Info("用户注册成功", zap.Uint64("userId", user.Id), zap.String("username", user.Username))
	return user, nil
}

func LoginUser(username, password string) (string, *model.UserJson, error) {

	zap.L().Info("尝试登录用户", zap.String("username", username))

	user, err := model.GetUserByUsername(username)
	if err != nil {
		zap.L().Error("查询用户失败", zap.String("username", username), zap.String("error", err.Error()))
		return "", nil, err
	}
	if user == nil || user.Password != password {
		zap.L().Warn("用户名或密码错误", zap.String("username", username))
		return "", nil, errors.New("用户名或密码错误")
	}

	token, err := utils.GenerateToken(user.Id, user.IsAdmin)
	if err != nil {
		zap.L().Error("生成 Token 失败", zap.Uint64("userId", user.Id), zap.String("error", err.Error()))
		return "", nil, err
	}

	zap.L().Info("用户登录成功", zap.Uint64("userId", user.Id), zap.String("username", user.Username))
	return token, user, nil
}

// 绑定用户和学生
func BindUserStudent(userId, studentId uint64) error {
	if userId == 0 || studentId == 0 {
		return errors.New("用户ID和学生ID不能为空")
	}

	exists, err := model.ExistsUserStudent(userId, studentId)
	if err != nil {
		zap.L().Error("查询绑定关系失败", zap.Uint64("userId", userId), zap.Uint64("studentId", studentId), zap.Error(err))
		return err
	}
	if exists {
		return errors.New("绑定关系已存在")
	}

	if err := model.CreateUserStudent(userId, studentId); err != nil {
		zap.L().Error("创建绑定关系失败", zap.Uint64("userId", userId), zap.Uint64("studentId", studentId), zap.Error(err))
		return err
	}

	zap.L().Info("绑定用户和学生成功", zap.Uint64("userId", userId), zap.Uint64("studentId", studentId))
	return nil
}

// 解绑用户和学生
// 解绑用户和学生
func UnbindUserStudent(userId uint64) error {
	if userId == 0 {
		return errors.New("用户ID不能为空")
	}

	rows, err := model.DeleteUserStudent(userId)
	if err != nil {
		zap.L().Error("解绑失败", zap.Uint64("userId", userId), zap.Error(err))
		return err
	}

	if rows == 0 {
		return errors.New("绑定关系不存在")
	}

	zap.L().Info("解绑用户和学生成功", zap.Uint64("userId", userId))
	return nil
}
