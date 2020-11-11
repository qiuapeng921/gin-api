package curl

import (
	"bytes"
	"gin-api/app/utility/system"
	"io"
)

// response http response wrapper
type response struct {
	body []byte
}

// String return response as string
func (wrap *response) String() string {
	return string(wrap.body)
}

// Byte return response as byte
func (wrap *response) Byte() []byte {
	return wrap.body
}

// BindJson bind json object with pointer
func (wrap *response) BindJson(object interface{}) error {
	return system.JsonToStruct(string(wrap.body), object)
}

// Reader
func (wrap *response) Reader() io.Reader {
	return bytes.NewReader(wrap.body)
}
