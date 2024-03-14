package mheader

import "net/http"

type AddHeader struct {
}

// 假设每个package中必须使用New函数返回
func New() http.Handler {
	return &AddHeader{}
}

func (this *AddHeader) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("name", "mheader--v3")
}
