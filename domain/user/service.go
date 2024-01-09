package user

import "shopping_go/utils/hash"

// Service 用户 service构造体
type Service struct {
	r Repository
}

// NewUserService 实例化user service
func NewUserService(repository Repository) *Service {
	repository.Migration()        //创建表
	repository.InsertSampleData() //创建测试数据
	return &Service{
		r: repository,
	}
}

// Create 新增用户
func (c *Service) Create(user *User) error {
	if user.Password != user.Password2 {
		return ErrMismatchedPasswords
	}
	//判断用户名是否存在
	_, err := c.r.GetByName(user.Username)
	if err != nil {
		return ErrUserExistWithName
	}
	//验证用户名
	if ValidateUserName(user.Username) {
		return ErrInvalidUsername
	}
	//验证密码
	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}
	err = c.r.Create(user)
	return err
}

// GetUser 查询用户
func (c *Service) GetUser(username string, password string) (User, error) {
	user, err := c.r.GetByName(username)
	if err != nil {
		return User{}, err
	}
	passwordHash := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !passwordHash {
		return User{}, err
	}
	return user, err
}

// UpdateUser 更新用户数据
func (c *Service) UpdateUser(user *User) error {
	return c.r.Update(user)
}
