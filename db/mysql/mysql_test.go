package tymysql

import (
	"log"
	"reflect"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// 表结构
type TUser struct {
	Name    string
	Age     int
	Sex     int
	Address string
}

var (
	jsonCfg *JsonConfig
	// yamlCfg *YamlConfig
	wg sync.WaitGroup
)

func SqlQuery(s *Mysql) {
	rows, err := s.DB().Query("SELECT * FROM user")
	defer func() {
		_ = rows.Close()
	}()

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

func SqlQueryRow(s *Mysql) {
	row := s.DB().QueryRow("SELECT * FROM user WHERE name='levi';")
	var u TUser
	err := row.Scan(&u.Name, &u.Age, &u.Sex, &u.Address)
	if err == nil {
		log.Println(u)
	}
}

func SqlQueryWhere(s *Mysql) {
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

func SqlDelete(s *Mysql) {
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

func SqlUpdate(s *Mysql) {
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

func SqlInsert(s *Mysql) {
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
	jsonCfg, err = LoadJsonWithData([]byte(`{
		"username": "root",
		"password": "Aabbcc123!@#",
		"network": "tcp",
		"host": "47.97.103.250:3307",
		"database": "kingshard",
		"MaxOpenConns": 2,
		"MaxOdleConns": 1
}`))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(jsonCfg)

	db := &Mysql{}
	err = db.OpenWithJsonConfig(jsonCfg)
	if err != nil {
		t.Error(err)
		return
	}

	// SqlQuery(db)
	// SqlQueryRow(db)
	// SqlQueryWhere(db)
	// SqlUpdate(db)
	// SqlInsert(db)
	// sqlDelete(db)
}

func TestLoadWithYamlData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *YamlConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: struct{ data []byte }{data: []byte(`
												username: root
												password: lwstar
												host: 127.0.0.1:3306
												port: 3306
												database: lw
												maxOpenConns: 2
												maxOdleConns: 1
											`)}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadWithYamlData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadWithYamlData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadWithYamlData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadWithYaml(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *YamlConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: struct{ file string }{file: "conf.yaml"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadWithYaml(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadWithYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadWithYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadJsonWithData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *JsonConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: struct{ data []byte }{data: []byte(`{
										"username": "root",
										"password": "Aabbcc123!@#",
										"host": "47.96.230.81:3306",
										"database": "mservice",
										"MaxOpenConns": 20,
										"MaxOdleConns": 5
									}`)}, want: &JsonConfig{
			Username:     "root",
			Password:     "Aabbcc123!@#",
			Database:     "47.96.230.81:3306",
			Host:         "mservice",
			MaxOpenConns: 20,
			MaxOdleConns: 5,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadJsonWithData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadJsonWithData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadJsonWithData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadJsonConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *JsonConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: struct{ file string }{file: "conf.json"}, want: &JsonConfig{
			Username:     "root",
			Password:     "Aabbcc123!@#",
			Database:     "47.96.230.81:3306",
			Host:         "mservice",
			MaxOpenConns: 20,
			MaxOdleConns: 5,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadJsonConfig(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadJsonConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadJsonConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
