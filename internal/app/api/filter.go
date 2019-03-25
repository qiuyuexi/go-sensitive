package api

import (
	"encoding/json"
	"fmt"
	"go-sensitive/internal/pkg/ahocorasick"
	"go-sensitive/internal/pkg/template"
	"net/http"
	"strconv"
	"strings"
)

func FilterHandel(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	}()
	parseErr := req.ParseForm()

	if parseErr != nil {
		w.WriteHeader(500)
		return
	}

	if _, ok := req.PostForm["content"]; !ok {
		result := template.ServerErr(400, "content 不能为空")
		res, _ := json.Marshal(result)
		_, err := w.Write(res)
		if err != nil {
			fmt.Println("error")
			w.WriteHeader(500)
		}
		return
	}

	if _, ok := req.PostForm["group_id"]; !ok {
		result := template.ServerErr(400, "group_id 不能为空")
		res, _ := json.Marshal(result)
		_, err := w.Write([]byte(res))
		if err != nil {
			fmt.Println("error")
			w.WriteHeader(500)
		}
		return
	}

	content := req.PostForm["content"][0]
	groupId := req.PostForm["group_id"][0]
	groupIdList := strings.Split(groupId, ",")

	//获取ac自动机
	Ac := ahocorasick.GetAcDictIns()
	result := make(map[int][]string)
	for _, v := range groupIdList {
		index, _ := strconv.Atoi(v)
		if Ac.Tree[index] != nil {
			result[index] = Ac.Tree[index].Search(content)
		}
	}
	resSuccess := template.ServerSuccess(result)
	res, _ := json.Marshal(resSuccess)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	return
}
