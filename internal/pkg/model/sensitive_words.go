package model

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"reflect"
	"time"
)

type WordEntry struct {
	Id         int
	Word       string
	Group_id   int
	Created_at int
}

func getDb() (*sql.DB, error) {
	config := &mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		DBName:               "test",
		ReadTimeout:          1 * time.Second,
		WriteTimeout:         1 * time.Second,
		AllowNativePasswords: true,
	}
	formarDsn := config.FormatDSN()
	db, err := sql.Open("mysql", formarDsn)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

//获取单词列表
func GetSensitiveWordList() map[int][]string {
	wordList := make(map[int][]string)

	db, dbConnectErr := getDb()
	if dbConnectErr != nil {
		return wordList
	}

	defer db.Close()

	//查询
	maxId := 0
	cnt := 0
	limit := 2
	for {
		rows, dbQueryErr := db.Query("SELECT * FROM words WHERE id > ?  limit ?", maxId, limit)

		if dbQueryErr != nil {
			fmt.Println(dbQueryErr)
			return wordList
		}
		cnt = 0
		for rows.Next() {
			entry := new(WordEntry)
			addrList := getAddrList(entry) //获取结构体中每个变量的地址
			rows.Scan(addrList...)
			wordList[entry.Group_id] = append(wordList[entry.Group_id], entry.Word)
			maxId = entry.Id
			cnt++
		}
		if cnt < limit {
			break
		}
	}
	return wordList
}

//获取结构体变量的地址
func getAddrList(entry *WordEntry) []interface{} {
	s := reflect.ValueOf(entry).Elem()
	addr := make([]interface{}, 0)
	for i := 0; i < s.NumField(); i++ {
		add := s.Field(i).Addr().Interface()
		addr = append(addr, add)
	}
	return addr
}
