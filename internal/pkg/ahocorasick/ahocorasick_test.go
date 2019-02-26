package ahocorasick

import "testing"

func TestGetAhocorasick(t *testing.T) {

}

func TestGetTire(t *testing.T) {
	words := []string{"测试", "玩笑", "啊哈哈哈哈", "哦哦哦哦哦", "12345", "aaaabb", "aaaAAAA"}
	trie := GetTire(words)
	result := trie.Search("测试玩笑")
	var exist int
	exist = 0
	for _, v := range result {
		if v == "测试" {
			exist = 1
		}
	}
	if exist != 1 {
		t.Errorf("serarch result:%v", result)
	}
	exist = 0
	result2 := trie.Search("玩测笑")
	for _, v := range result2 {
		if v == "测试" {
			exist = 1
		}
	}
	if exist != 0 {
		t.Errorf("serarch result:%v", result)
	}
}
