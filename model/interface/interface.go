package _interface

type CRUD interface {
	CreateUser()       //创建用户
	UpdateUserInfo()   //修改用户信息
	RetrieveLogin()    //用户登录查询数据库密码
	DeleteUser()       //注销用户
	RetrieveUserInfo() //查看用户信息
	ReadUserId()       //查询用户id
}
