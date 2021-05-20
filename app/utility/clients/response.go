package clients

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// -----------------------------------------------------------------------------

// ResponseProxy 读取代理类
type ResponseProxy struct {
	c        context.Context
	request  *http.Request
	response *http.Response
}

// ResponseWrapper http.Response 包装
func ResponseWrapper(c context.Context, req *http.Request, resp *http.Response) *ResponseProxy {
	return &ResponseProxy{c: c, request: req, response: resp}
}

// Transform 数据转换
type Transform func([]byte) (interface{}, error)

// Observer 数据流操作
type Observer func(c context.Context, reader io.Reader) error

// Read 读取数据
func (p *ResponseProxy) Read(transform Transform) (interface{}, error) {
	defer func() {
		if err := p.response.Body.Close(); err != nil {
			log.Panicf("httpclient 流关闭失败 错误信息：[%s]", err)
		}
	}()
	b, err := ioutil.ReadAll(p.response.Body)
	if err != nil {
		return nil, err
	}
	return transform(b)
}

// Stream 流式读取数据
func (p *ResponseProxy) Stream(observer Observer) error {
	defer func() {
		if err := p.response.Body.Close(); err != nil {
			log.Panicf("httpclient 流关闭失败 错误信息：[%s]", err)
		}
	}()
	return observer(p.c, p.response.Body)
}

// Request 返回 http 请求配置对象
func (p *ResponseProxy) Request() *http.Request {
	return p.request
}

// Response 返回 http 应答对象
func (p *ResponseProxy) Response() *http.Response {
	return p.response
}

// Int64 转换 64 位整数
func Int64(b []byte) (interface{}, error) {
	return strconv.ParseInt(string(b), 10, 64)
}

// Text 读取字符串
func Text(b []byte) (interface{}, error) {
	return string(b), nil
}

// JSON 读取 JSON 数据
func JSON(v interface{}) Transform {
	return func(b []byte) (i interface{}, err error) {
		err = json.Unmarshal(b, v)
		return v, err
	}
}

// Writer 数据写入到 writer
func Writer(writer io.Writer) Observer {
	return func(c context.Context, reader io.Reader) error {
		w := bufio.NewWriter(writer)
		n, err := bufio.NewReader(reader).WriteTo(w)
		if err != nil {
			return err
		}
		log.Println(fmt.Sprintf("cope length [%d]", n))
		return w.Flush()
	}
}

// File 数据写入到文件
func File(path string, flag int, perm os.FileMode) Observer {
	return func(c context.Context, reader io.Reader) error {
		file, err := os.OpenFile(path, flag, perm)
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Println("文件关闭错误:", err)
			}
		}()
		return Writer(file)(c, reader)
	}
}