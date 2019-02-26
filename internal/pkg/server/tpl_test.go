package server

import "testing"

func TestServerErr(t *testing.T) {
	code := 500
	msg := "error"
	res := ServerErr(code, msg)
	if res.Head.Code != 500 || res.Head.Msg != msg || res.Body.Data != struct{}{} {
		t.Errorf("code:%v,msg:%v,body:%v", res.Head.Code, res.Head.Msg, res.Body.Data)
	}
}

func TestServerSuccess(t *testing.T) {
	data := "success";
	res := ServerSuccess(data)
	if res.Head.Code != 200 || res.Head.Msg != "" || res.Body.Data != data {
		t.Errorf("code:%v,msg:%v,body:%v", res.Head.Code, res.Head.Msg, res.Body.Data)
	}
}
