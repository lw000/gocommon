package test

import (
	"github.com/lw000/gocommon/db/mysql"
	"log"
	"testing"
)

// 表结构
type TUser struct {
	Name    string `json:"Name" form:"name"`
	Age     int    `json:"Age" form:"age"`
	Sex     int    `json:"Sex" form:"sex"`
	Address string `json:"Address" form:"address"`
}

var dbcfg *tymysql.JsonConfig
var dbYamlcfg *tymysql.YamlConfig

func SqlQuery(s *tymysql.Mysql) {
	rows, err := s.DB().Query("SELECT * FROM user")
	defer rows.Close()

	if err != nil {
		log.Panic(err)
	}

	var us []TUser
	for rows.Next() {
		var u TUser
		if err = rows.Scan(&u.Name, &u.Age, &u.Sex, &u.Address); err == nil {
			us = append(us, u)
		}
	}
	if err = rows.Err(); err != nil {
		log.Panic(err)
	}

	for _, v := range us {
		log.Println(v)
	}
}

func SqlQueryRow(s *tymysql.Mysql) {
	row := s.DB().QueryRow("SELECT * FROM user WHERE name='levi';")
	var u TUser
	err := row.Scan(&u.Name, &u.Age, &u.Sex, &u.Address)
	if err == nil {
		log.Println(u)
	}
}

func SqlQueryWhere(s *tymysql.Mysql) {
	stms, err := s.DB().Prepare("SELECT * FROM user WHERE name=?;")
	if err != nil {
		log.Panic(err)
	}

	defer stms.Close()

	rows, err := stms.Query("levi")
	if err != nil {
		log.Panic(err)
	}

	var us []TUser
	for rows.Next() {
		var u TUser
		if err = rows.Scan(&u.Name, &u.Age, &u.Sex, &u.Address); err == nil {
			us = append(us, u)
		}
	}
	if err = rows.Err(); err != nil {
		log.Panic(err)
	}

	for _, v := range us {
		log.Println(v)
	}
}

func SqlDelete(s *tymysql.Mysql) {
	rs, err := s.DB().Exec("DELETE FROM user WHERE name=?;", "levi")
	if err != nil {
		log.Panic(err)
	}
	id, err := rs.RowsAffected()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("sqlDelete:%d\n", id)
}

func SqlUpdate(s *tymysql.Mysql) {
	rs, err := s.DB().Exec("UPDATE user SET name=? WHERE name=?;", "levi1", "levi")
	if err != nil {
		log.Panic(err)
	}

	id, err := rs.RowsAffected()
	if err != nil {
		log.Panic(err)
	}

	log.Printf("sqlUpdate:%d\n", id)
}

func SqlInsert(s *tymysql.Mysql) {
	rs, err := s.DB().Exec("INSERT INTO user(name, age, sex ,address) VALUES(?,?,?,?);", "hjt", 30, 1, "beijingshi")
	if err != nil {
		log.Panic(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Panic(err)
	}

	log.Println("sqlInsert: ", id)
}

func TestRunTest(t *testing.T) {
	var err error
	dbcfg, err = tymysql.LoadJsonConfig("../conf/mysql.json")
	if err != nil {
		log.Panic(err)
	}
	log.Println(dbcfg)

	dbYamlcfg, err = tymysql.LoadWithYaml("../conf/mysql.yaml")
	if err != nil {
		log.Panic(err)
	}
	log.Println(dbYamlcfg)

	db := &tymysql.Mysql{}
	err = db.OpenWithJsonConfig(dbcfg)
	err = db.OpenWithYamlConfig(dbYamlcfg)
	if err != nil {
		log.Panic(err)
	}
	SqlQuery(db)
	SqlQueryRow(db)
	SqlQueryWhere(db)
	SqlUpdate(db)
	//SqlInsert(db)
	// sqlDelete(db)
}
