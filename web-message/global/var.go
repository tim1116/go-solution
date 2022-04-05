package global

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

var (
	user      string
	pass      string
	address   string
	dbname    string
	charset   string
	parseTime bool
)

func init() {
	viper.SetConfigName("database") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	viper.AddConfigPath("config")   //设置配置文件的搜索目录
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("配置文件读取异常,err: %s", err))
	}
	user = viper.GetString(`database.username`)
	pass = viper.GetString(`database.password`)
	address = viper.GetString(`database.address`)
	dbname = viper.GetString(`database.dbname`)
	charset = viper.GetString(`database.charset`)
	parseTime = viper.GetBool(`database.parseTime`)
}

// InitDb 初识化数据库链接
func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", user, pass, address, dbname, charset, parseTime)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gorm_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	Db = db
}
