package db

import (
	"fmt"
	"go-app/config"
	"go-app/lib/database"
	"go-app/lib/logger"

	"gorm.io/gorm"
)

var Conn *gorm.DB

func init() {

	var err error

	// 如果有多个数据库链接需求，或者有主从，负载均衡等
	// InitMultiDatabase()
	// Conn.Logger = logger.NewGormLogger()
	// // 创建数据表
	// createTable()

	// 如果只是单一数据库
	Conn, err = database.ConnDB(config.DB_APP.DSN, config.DB_APP.DB_TYPE)
	if err != nil {
		panic(err)
	}
	Conn.Logger = logger.NewGormLogger()
}

func createTable() {
	// if !Conn.Migrator().HasTable(&model.Customer{}) {
	// 	Conn.Migrator().CreateTable(&model.Customer{})
	// }
	// if !Conn.Migrator().HasTable(&model.CustomerBrand{}) {
	// 	Conn.Migrator().CreateTable(&model.CustomerBrand{})
	// }
	// if !Conn.Migrator().HasTable(&model.User{}) {
	// 	Conn.Migrator().CreateTable(&model.User{})
	// }
	// if !Conn.Migrator().HasTable(&model.UserSns{}) {
	// 	Conn.Migrator().CreateTable(&model.UserSns{})
	// }
	// if !Conn.Migrator().HasTable(&model.PhoneCheck{}) {
	// 	Conn.Migrator().CreateTable(&model.PhoneCheck{})
	// }
	// if !Conn.Migrator().HasTable(&model.PreSimilar{}) {
	// 	Conn.Migrator().CreateTable(&model.PreSimilar{})
	// }
	// if !Conn.Migrator().HasTable(&model.PreMissingBrand{}) {
	// 	Conn.Migrator().CreateTable(&model.PreMissingBrand{})
	// }

	// ConnBusiness, _ := gorm.Open(database.GenDialector(config.DB_BUSINESS.DSN, database.POSTGRES), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// 	PrepareStmt:                              true,
	// })

	// if !ConnBusiness.Migrator().HasTable(&model.BusinessCompany{}) {
	// 	ConnBusiness.Migrator().CreateTable(&model.BusinessCompany{})
	// }
	// if !ConnBusiness.Migrator().HasTable(&model.Business{}) {
	// 	ConnBusiness.Migrator().CreateTable(&model.Business{})
	// }
	// if !ConnBusiness.Migrator().HasTable(&model.BusinessNode{}) {
	// 	ConnBusiness.Migrator().CreateTable(&model.BusinessNode{})
	// }
	// if !ConnBusiness.Migrator().HasTable(&model.BusinessSendHistory{}) {
	// 	ConnBusiness.Migrator().CreateTable(&model.BusinessSendHistory{})
	// }
	// if !ConnBusiness.Migrator().HasTable(&model.BusinessTakeHistory{}) {
	// 	ConnBusiness.Migrator().CreateTable(&model.BusinessTakeHistory{})
	// }

	// addGinIndex(ConnBusiness, "company", "name_used")
	// addGinIndex(ConnBusiness, "company", "phone")
	// addGinIndex(ConnBusiness, "company", "phone_check")
	// addGinIndex(ConnBusiness, "node", "detail")
}

// postgres 添加json数据的gin索引
func addGinIndex(Conn *gorm.DB, tableName, fieldName string) {
	indexName := "idx_" + tableName + "_" + fieldName
	createIndexStatement := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s USING gin (%s)", indexName, tableName, fieldName)
	Conn.Exec(createIndexStatement)
}

func InitMultiDatabase() {
	// 具体可参考 https://github.com/go-gorm/dbresolver
	// var err error
	// Conn, err = gorm.Open(database.GenDialector(config.DB_APP.DSN, database.POSTGRES), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// 	PrepareStmt:                              true,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// // 设置主库的线程池
	// sqlDB, err := Conn.DB()
	// sqlDB.SetMaxIdleConns(config.Pool.MaxIdleConns)
	// sqlDB.SetMaxOpenConns(config.Pool.MaxOpenConns)
	// sqlDB.SetConnMaxIdleTime(config.Pool.ConnMaxIdleTime)
	// sqlDB.SetConnMaxLifetime(config.Pool.ConnMaxLifetime)
	// if err != nil {
	// 	panic(err)
	// }

	// // 这里指定特定的表去特定的数据库
	// // 还可以设置从库，负载均衡等
	// slover := dbresolver.Register(
	// 	dbresolver.Config{
	// 		Sources: []gorm.Dialector{database.GenDialector(config.DB_DATA.DSN, database.POSTGRES)},
	// 	},
	// 	&model.Brand{},
	// 	&model.Coowner{},
	// 	&model.Flow{},
	// 	&model.FlowLast{},
	// 	&model.Global{},
	// 	&model.Item{},
	// 	&model.Owner{},
	// 	// &model.OwnerHistory{},
	// ).Register(
	// 	dbresolver.Config{
	// 		Sources: []gorm.Dialector{database.GenDialector(config.DB_FLOW.DSN, database.POSTGRES)},
	// 	},
	// 	&model.TmFlowStatus{},
	// ).Register(
	// 	dbresolver.Config{
	// 		Sources: []gorm.Dialector{database.GenDialector(config.DB_BUSINESS.DSN, database.POSTGRES)},
	// 	},
	// 	&model.BusinessCompany{},
	// 	&model.Business{},
	// 	&model.BusinessNode{},
	// 	&model.BusinessSendHistory{},
	// 	&model.BusinessTakeHistory{},
	// )

	// // 设置连接池信息
	// slover.SetConnMaxIdleTime(config.Pool.ConnMaxIdleTime).
	// 	SetConnMaxLifetime(config.Pool.ConnMaxLifetime).
	// 	SetMaxIdleConns(config.Pool.MaxIdleConns).
	// 	SetMaxOpenConns(config.Pool.MaxOpenConns)

	// Conn.Use(slover)
}
