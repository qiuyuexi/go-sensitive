package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var appPath string

func SetWorkPath() {
	workPath, err := os.Getwd()
	if err != nil {
		fmt.Println()
	}
	appPath = filepath.Dir(workPath)
}

//文件中读取配置
func LoadSensitiveWords() map[int][]string {
	words := make(map[int][]string)
	inputReader,err := ioutil.ReadFile(appPath + "/env/words/words.json")
	if err!= nil{

	}
	json.Unmarshal(inputReader,&words)
	return words
}
