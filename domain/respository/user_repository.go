package respository

import (
	"git.imooc.com/dfs/dfsmirco/domain/model"
	"github.com/jinzhu/gorm"
)

// IUserRepository 用户相关
type IUserRepository interface {
	// 初始化数据表
	InitTable() error
	// 根据用户名查找用户信息
	FindUserByName(string) (*model.User, error)
	// 根据用户id查找用户信息
	FindUserById(int64) (*model.User, error)
	// 创建用户
	CreateUser(*model.User) (int64, error)
	// 根据id删除用户
	DeleteUserById(int64) error
	// 用户数据更新
	UpdateUser(*model.User) error
	// 查找所有用户
	FindAll() ([]model.User, error)
}

// 新建用户
func NewUserRepository(db *gorm.DB) {

}
