package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// DAO 数据访问层
var (
	_db     *gorm.DB
	_config *Config
	_tables = make([]interface{}, 0)
)

// Init 初始化
func Init(config *Config, tables ...interface{}) {
	// 数据校验
	if config == nil {
		panic("dao Config Can't be nil")
	}
	if config.DB == "" ||
		config.Host == "" ||
		config.User == "" ||
		config.DBName == "" {
		panic("dao Config filed can not be empty")
	}
	if config.IDZones == nil ||
		len(config.IDZones) == 0 {
		panic("IDZones must set ")
	}
	checkIDZones(config.IDZones)

	_config = config
	// 添加id区段表
	_tables = append(_tables, &IDZone{})
	if tables != nil {
		_tables = append(_tables, tables...)
	}
}

// Init2 初始化
func Init2(config map[string]interface{}, tables ...interface{}) {
	if config == nil {
		panic("dao Config Can't be nil")
	}

	daoConfig := &Config{
		DB:       config["DB"].(string),
		Host:     config["Host"].(string),
		User:     config["User"].(string),
		DBName:   config["DBName"].(string),
		PassWord: config["PassWord"].(string),
	}
	Init(daoConfig, tables...)
}

// Connect 连接数据库
func Connect() (db *gorm.DB, err error) {
	_db, err = gorm.Open(_config.DB, fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		_config.Host, _config.User, _config.DBName, _config.PassWord))
	if err == nil {
		initTables()
		initIDZone(_db)
		//initFactory(_db)
		db = _db
	}
	return
}

// DisConnect 断开链接
func DisConnect() {
	if _db != nil {
		_db.Close()
		_db = nil
	}
}

// initTables 合并数据库表
func initTables() {
	// 初始化用户自定义表
	if _tables != nil && len(_tables) > 0 {
		if err := _db.AutoMigrate(_tables...).Error; err != nil {
			panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
		}
	}
}
