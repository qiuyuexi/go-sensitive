package template

/**
	需要大写 否者属性无法解析
 */
type headTpl struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type bodyTpl struct {
	Data interface{} `json:"data"`
}

type response struct {
	Head headTpl `json:"head"`
	Body bodyTpl `json:"body"`
}

func ServerErr(code int, msg string) *response {
	res := new(response)
	res.Body.Data = struct{}{}
	res.Head.Code = code
	res.Head.Msg = msg
	return res
}

func ServerSuccess(data interface{}) *response {
	res := new(response)
	res.Body.Data = data
	res.Head.Code = 200
	res.Head.Msg = ""
	return res
}
