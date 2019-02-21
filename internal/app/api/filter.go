package api

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"go-sensitive/internal/pkg/server"
	"net/http"
	"strconv"
	"strings"
)

func FilterHandel(w http.ResponseWriter, req *http.Request)  {
	req.ParseForm()
	content := req.PostForm["content"][0]
	treeNum := req.PostForm["tree_num"][0]
	treeNumList := strings.Split(treeNum,",")

	Ac := server.GetAcDictIns()
	result := make(map[int][]string)
	for _,v:= range treeNumList {
		index,_ := strconv.Atoi(v)
		if Ac.Tree[index] != nil{
			fmt.Println("----------" + strconv.Itoa(index))
			result[index] = Ac.Tree[index].Search(content)
		}

	}
	res,_ := json.Marshal(result)
	w.Write([]byte(res))
}
