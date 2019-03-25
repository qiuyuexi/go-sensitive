package ahocorasick

import (
	"go-sensitive/internal/pkg/config"
)

var AcDictIns *acDict

type acDict struct {
	isUpdate int
	ac       [2]*Ahocorasick
	isBuild  int
}

/**
生成ac自动机字典
 */
func BuildAhocorasickDict() {
	sensitiveWordConfig := config.LoadSensitiveWords()
	AcDictIns = new(acDict)
	AcDictIns.ac[0] = GetAhocorasick()
	for k, v := range sensitiveWordConfig {
		AcDictIns.ac[0].Tree[k] = GetTire(v)
	}
	AcDictIns.ac[1] = AcDictIns.ac[0]
	AcDictIns.isBuild = 1
	AcDictIns.isUpdate = 0
}

/**
敏感词列表更新，重建ac自动机
 */
func RebuildAhocorasickDict() {
	AcDictIns.isUpdate = 1
	sensitiveWordConfig := config.LoadSensitiveWords()
	newAcDictIns := new(acDict)
	newAcDictIns.isUpdate = 0
	newAcDictIns.ac[0] = GetAhocorasick()
	for k, v := range sensitiveWordConfig {
		newAcDictIns.ac[0].Tree[k] = GetTire(v)
	}
	newAcDictIns.ac[1] = newAcDictIns.ac[0]
	AcDictIns.ac[0] = newAcDictIns.ac[0]
	AcDictIns.ac[1] = newAcDictIns.ac[1]
	AcDictIns.isUpdate = 0
}

/**
获取ac自动机
 */
func GetAcDictIns() *Ahocorasick {
	var acIns *Ahocorasick

	//需要判断是否构建
	if AcDictIns == nil {
		BuildAhocorasickDict()
	}

	if AcDictIns.isUpdate == 0 {
		acIns = AcDictIns.ac[0]
	} else {
		acIns = AcDictIns.ac[1]
	}
	return acIns
}
