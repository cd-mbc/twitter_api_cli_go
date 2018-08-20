package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"

	"github.com/cd-mbc/twitter_api_cli_go/model"
	"github.com/cd-mbc/twitter_api_cli_go/setting"
)

var db_user = setting.DB_user
var db_pass = setting.DB_pass
var db_name = setting.DB_name

var dbms = setting.DBMS

var q_create_table_trend = "create table if not exists trend (name varchar(100), url varchar(200), tweet_volume int);"
var q_select_trend_all = "select * from trend"
var q_delete_trend_all = "delete from trend"

func get_db_path() string {
	return db_user + ":" + db_pass + "@" + "tcp(localhost:3306)/" + db_name
}

func Init_db() {
	db, err := sql.Open(dbms, get_db_path())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	_, err = dbmap.Exec(q_create_table_trend)
	if err != nil {
		panic(err.Error())
	}
}

func Insert_trend(trend *model.Trend) {
	db, err := sql.Open(dbms, get_db_path())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(model.Trend{}, "trend")

	err = dbmap.Insert(trend)
	if err != nil {
		panic(err.Error())
	}
}

func Select_trend_all() []model.Trend {
	db, err := sql.Open(dbms, get_db_path())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(model.Trend{}, "trend")

	var trend_seq []model.Trend
	_, err = dbmap.Select(&trend_seq, q_select_trend_all)
	if err != nil {
		panic(err.Error())
	}

	return trend_seq
}

func Delete_trend_all() {
	db, err := sql.Open(dbms, get_db_path())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(model.Trend{}, "trend")

	_, err = dbmap.Exec(q_delete_trend_all)
	if err != nil {
		panic(err.Error())
	}
}
