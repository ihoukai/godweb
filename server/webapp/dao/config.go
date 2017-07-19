package dao

// Config dao层配置信息
type Config struct {
	DB       string    // 数据库的名字(postgres、mysql、mssql)
	Host     string    // 主机
	User     string    // 用户
	DBName   string    // 数据库
	PassWord string    // 密码
	IDZones  []*IDZone // id区段表
}
