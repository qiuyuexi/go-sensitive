package server

import (
	"fmt"
	"go-sensitive/internal/pkg/ahocorasick"
	"go-sensitive/internal/pkg/config"
)

var AcDictIns *acDict

type acDict struct {
	isUpdate int
	ac       [2]*ahocorasick.Ahocorasick
}

func Start() {
	config.SetWorkPath()
	ahocorasick.BuildAhocorasickDict()
	Watch()
	printOut()
}

func printOut() {
	fmt.Println("启动...")
}
