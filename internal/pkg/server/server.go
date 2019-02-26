package server

import (
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
	buildAhocorasickDict()
}

/**
生成ac自动机字典
 */
func buildAhocorasickDict() {
	sensitiveWordConfig := config.LoadSensitiveWords()
	AcDictIns = new(acDict)
	AcDictIns.isUpdate = 0
	AcDictIns.ac[0] = ahocorasick.GetAhocorasick()
	for k, v := range sensitiveWordConfig {
		AcDictIns.ac[0].Tree[k] = ahocorasick.GetTire(v)
	}
	AcDictIns.ac[1] = AcDictIns.ac[0]
}

/**
获取ac自动机
 */
func GetAcDictIns() *ahocorasick.Ahocorasick {
	var acIns *ahocorasick.Ahocorasick
	if AcDictIns.isUpdate == 0 {
		acIns = AcDictIns.ac[0]
	} else {
		acIns = AcDictIns.ac[1]
	}
	return acIns
}
