package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/supernet/game/hashrocket/pkg/log"
	"github.com/supernet/game/hashrocket/pkg/viper"

	"time"
)

var (
	MasterDB *xorm.Engine
	err      error
	Slave1DB *xorm.Engine
	err1     error
)

func MasterInit() *xorm.Engine {
	open := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", viper.Vp.GetString("mysql-master.username"),
		viper.Vp.GetString("mysql-master.password"),
		viper.Vp.GetString("mysql-master.addr"),
		viper.Vp.GetInt64("mysql-master.port"),
		viper.Vp.GetString("mysql-master.database"))

	MasterDB, err = xorm.NewEngine("mysql", open)
	if err != nil {
		log.ZapLog.Error(fmt.Sprintf("Open mysql-master failed,err:%s\n", err.Error()))
		panic(err)
	}

	MasterDB.SetConnMaxLifetime(100 * time.Second)
	MasterDB.SetMaxOpenConns(100)
	MasterDB.SetMaxIdleConns(16)
	err = MasterDB.Ping()
	if err != nil {
		log.ZapLog.Error(fmt.Sprintf("Failed to connect to mysql-master, err:%s" + err.Error()))
		panic(err.Error())
	}

	// 显示打印语句
	if viper.Vp.GetString("active") == "dev" || viper.Vp.GetString("active") == "test" {
		MasterDB.ShowSQL(true)
	}

	log.ZapLog.Info("mysql-master connect success\r\n")
	return MasterDB
}

func Slave1Init() *xorm.Engine {
	open := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", viper.Vp.GetString("mysql-slave1.username"),
		viper.Vp.GetString("mysql-slave1.password"),
		viper.Vp.GetString("mysql-slave1.addr"),
		viper.Vp.GetInt64("mysql-slave1.port"),
		viper.Vp.GetString("mysql-slave1.database"))

	Slave1DB, err1 = xorm.NewEngine("mysql", open)
	if err != nil {
		log.ZapLog.Error(fmt.Sprintf("Open mysql-slave1 failed,err:%s\n", err.Error()))
		panic(err)
	}

	Slave1DB.SetConnMaxLifetime(100 * time.Second)
	Slave1DB.SetMaxOpenConns(100)
	Slave1DB.SetMaxIdleConns(16)

	// 显示打印语句
	if viper.Vp.GetString("active") == "dev" || viper.Vp.GetString("active") == "test" {
		Slave1DB.ShowSQL(true)
	}

	err1 = Slave1DB.Ping()
	if err != nil {
		log.ZapLog.Error(fmt.Sprintf("Failed to connect to mysql-slave1, err:%s" + err1.Error()))
		panic(err.Error())
	}

	log.ZapLog.Info("mysql-slave1 connect success\r\n")
	return Slave1DB
}
