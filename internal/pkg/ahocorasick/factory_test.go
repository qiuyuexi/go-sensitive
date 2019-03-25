package ahocorasick

import (
	"go-sensitive/internal/pkg/config"
	"testing"
)

func TestGetAcDictIns(t *testing.T) {
	config.SetWorkPath("/www/go/src/go-sensitive")
	acDit := GetAcDictIns()
	searchStr := "test";
	result := acDit.Tree[1].Search(searchStr)
	if len(result) != 1 {
		t.Errorf("serarch error:%d", len(result))
	}
}
