package core

import (
	"log"
	"os"
	"time"
	"x_admin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db = initMysql()

func GetDB() *gorm.DB {
	return db
}

// initMysql 初始化mysql会话
func initMysql() *gorm.DB {
	// fmt.Printf("%#v\n", config.DBConfig)
	// 日志配置
	slowThreshold := time.Second
	ignoreRecordNotFoundError := true
	logLevel := logger.Warn
	if config.DBConfig.LogLevel == "info" {
		logLevel = logger.Info
		ignoreRecordNotFoundError = false
	}
	if config.DBConfig.SlowThreshold > 0 {
		slowThreshold = time.Duration(config.DBConfig.SlowThreshold) * time.Second
	}

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             slowThreshold,             // 慢 SQL 阈值
			LogLevel:                  logLevel,                  // 日志级别
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError, // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                      // 彩色打印
		},
	)
	// 初始化会话
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DBConfig.Dsn,               // DSN data source name
		DefaultStringSize:         config.DBConfig.DefaultStringSize, // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                             // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.DBConfig.TablePrefix, // 表名前缀
			SingularTable: true,                        // 使用单一表名, eg. `User` => `user`
		},
		DisableForeignKeyConstraintWhenMigrating: true,   // 禁用自动创建外键约束
		Logger:                                   logger, // 自定义Logger
	})
	if err != nil {
		log.Fatal("initMysql gorm.Open err:", err)
	}
	db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
	// 🚀 注册插件
	// if err := db.Use(&ExistPlugin{}); err != nil {
	// 	panic(err)
	// }

	sqlDB, err := db.DB() //通用的数据库接口 *sql.DB
	if err != nil {
		log.Fatal("initMysql db.DB err:", err)
	}
	// 数据库空闲连接池最大值
	sqlDB.SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	// 数据库连接池最大值
	sqlDB.SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.DBConfig.ConnMaxLifetimeSeconds) * time.Second)
	// 定时打印DBStats
	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 1)
	// 		stats := sqlDB.Stats()
	// 		log.Printf("DBStats: OpenConnections =%d, 正在使用InUse=%d, 空闲连接数Idle=%d, 等待的连接总数WaitCount=%d, 阻塞等待新连接的总时间WaitDuration=%s,MaxIdleTimeClosed=%d,MaxLifetimeClosed=%d \n",
	// 			stats.OpenConnections, stats.InUse, stats.Idle, stats.WaitCount, stats.WaitDuration, stats.MaxIdleTimeClosed, stats.MaxLifetimeClosed)
	// 	}
	// }()
	return db
}

func DBTableName(model any) string {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	return stmt.Schema.Table
}
