package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"osquey/model"
	"time"
)

func dsn() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)
}

func DB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(model.OsVersion))
	orm.RegisterModel(new(model.OsQueryVersion))
	orm.RegisterModel(new(model.ProgramList))
	dsn := dsn()
	k := 0
	for k == 0 {
		err := orm.RegisterDataBase("default", "mysql", dsn, 30)
		fmt.Println("error is", err)
		if err == nil {
			k = 1
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("DB connection success...")

	orm.RunSyncdb("default", false, true)

}
