package modules

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/illidan33/wow_api/global"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var DbConn *gorm.DB

func init() {
	var err error
	// 确保使用你从环境变量中获取到的 DbHost 变量
	// 注意这里的 %d
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Local",
		global.Config.DbUser,
		global.Config.DbPwd,
		global.Config.DbHost,
		global.Config.DbPort,
	)

	db, err := gorm.Open("mysql", dsnWithoutDB)
	if err != nil {
		panic(fmt.Sprintf("无法连接到 MySQL 实例: %v", err))
	}
	defer db.Close()
	// 3. 第二步：执行 SQL 创建数据库（如果不存在）
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", global.Config.DbName)
	err = db.Exec(createSql).Error
	if err != nil {
		panic(fmt.Sprintf("自动创建数据库失败: %v", err))
	}
	fmt.Printf("数据库检查完毕: %s\n", global.Config.DbName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		global.Config.DbUser,
		global.Config.DbPwd,
		global.Config.DbHost,
		global.Config.DbPort,
		global.Config.DbName,
	)

	DbConn, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	DbConn.SingularTable(true)
	DbConn.SetLogger(global.Log)
	if global.Config.LogLevel == logrus.DebugLevel {
		DbConn.LogMode(true)
	}
}
