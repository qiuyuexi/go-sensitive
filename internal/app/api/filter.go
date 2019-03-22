package api

import (
	"encoding/json"
	"fmt"
	"go-sensitive/internal/pkg/ahocorasick"
	"go-sensitive/internal/pkg/server"
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
		result := server.ServerErr(400, "content 不能为空")
		fmt.Println(result)
		res, _ := json.Marshal(result)
		_, err := w.Write(res)
		if err != nil {
			fmt.Println("error")
			w.WriteHeader(500)
		}
		return
	}

	if _, ok := req.PostForm["tree_num"]; !ok {
		result := server.ServerErr(400, "tree_num 不能为空")
		res, _ := json.Marshal(result)
		fmt.Println(res)
		_, err := w.Write([]byte(res))
		if err != nil {
			fmt.Println("error")
			w.WriteHeader(500)
		}
		return
	}

	content := req.PostForm["content"][0]
	treeNum := req.PostForm["tree_num"][0]
	treeNumList := strings.Split(treeNum, ",")

	//获取ac自动机
	Ac := ahocorasick.GetAcDictIns()
	result := make(map[int][]string)
	for _, v := range treeNumList {
		index, _ := strconv.Atoi(v)
		if Ac.Tree[index] != nil {
			result[index] = Ac.Tree[index].Search(content)
		}
	}
	resSuccess := server.ServerSuccess(result)
	res, _ := json.Marshal(resSuccess)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println("error")
		w.WriteHeader(500)
	}
	return
}
