package config

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strings"
)

type DBConf struct {
	user string
	password string
	ip string
	port string
	database string
}

// admin:pwd@tcp(192.168.56.2:3306)/table?charset=utf8
func (conf DBConf) connStr() string  {
	return conf.user + ":" + conf.password + "@tcp(" + conf.ip + ":" + conf.port + ")/" + conf.database + "?charset=utf8"
}

var mysql *sql.DB

func init() {
	_map := ReadProperties("/home/langyu/WorkSpace/go_project/go_mysql_demo/resource/db.properties")
	conf := DBConf{
		user:     _map["mysql.user"],
		password: _map["mysql.password"],
		ip:       _map["mysql.ip"],
		port:     _map["mysql.port"],
		database: _map["mysql.database"],
	}
	db, err := sql.Open("mysql", conf.connStr())
	if err != nil {
		_ = fmt.Errorf("open conn error: %s\n", err)
	} else {
		mysql = db
	}
}

func ReadProperties(path string) map[string]string {
	_map := make(map[string]string)
	f, fErr := os.Open(path)
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()
	if fErr != nil {
		panic(fErr)
		return _map
	}
	r := bufio.NewReader(f)
	for i := 0; ; i++ {
		b, _, fErr := r.ReadLine()
		if fErr != nil {
			if fErr == io.EOF {
				fmt.Printf("total %d line, file end\n", i)
				break
			}
			panic(fErr)
		}
		s := strings.TrimSpace(string(b))
		if len(s) == 0 {
			fmt.Printf("this line content len is 0\n")
			continue
		}
		if s[:1] == "#" {
			fmt.Printf("this line content is annotation : %s\n", s)
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			fmt.Printf("this line content format error : %s\n", s)
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			fmt.Printf("this line content not has key : %s\n", s)
			continue
		}
		v := strings.TrimSpace(s[index+1:])
		if len(v) == 0 {
			fmt.Printf("this line content not has value : %s\n", v)
			continue
		}
		_map[key] = v
	}
	return _map
}

func GetDB() (*sql.DB, error) {
	if mysql == nil{
		return nil, errors.New("获取mysql连接失败")
	}
	return mysql, nil
}
