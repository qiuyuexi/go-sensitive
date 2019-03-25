package config

import (
	"encoding/json"
	"fmt"
	"go-sensitive/internal/pkg/model"
	"io/ioutil"
)

//加载敏感词列表
func LoadSensitiveWords() map[int][]string {
	words := make(map[int][]string)
	inputReader, err := ioutil.ReadFile(appPath + "/env/words/words.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(inputReader, &words)
	wordsFromDb := model.GetSensitiveWordList()

	for k, v := range wordsFromDb {
		words[k] = append(words[k], v...)
	}
	return words
}
