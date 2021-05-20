package clients

import (
	"bytes"
	"encoding/json"
	"gin-api/app/utility/app"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Option http.Request 配置方法
type Option func(*http.Request)

// HTTPClientOption HTTPClient 配置
type HTTPClientOption func(*HTTPClient)

// Timeout 超时时间
func Timeout(timeout time.Duration) HTTPClientOption {
	return func(client *HTTPClient) {
		client.client.Timeout = timeout
	}
}

// Header 设置 http.Request 头部信息
func Header(key, value string) Option {
	return func(request *http.Request) {
		request.Header.Set(key, value)
	}
}

// Body 设置 http.Request body 参数
func Body(body io.Reader) Option {
	return func(req *http.Request) {
		if body == nil {
			return
		}
		rc, ok := body.(io.ReadCloser)
		if !ok && body != nil {
			rc = ioutil.NopCloser(body)
		}
		req.Body = rc
		if body != nil {
			switch v := body.(type) {
			case *bytes.Buffer:
				req.ContentLength = int64(v.Len())
				buf := v.Bytes()
				req.GetBody = func() (io.ReadCloser, error) {
					r := bytes.NewReader(buf)
					return ioutil.NopCloser(r), nil
				}
			case *bytes.Reader:
				req.ContentLength = int64(v.Len())
				snapshot := *v
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return ioutil.NopCloser(&r), nil
				}
			case *strings.Reader:
				req.ContentLength = int64(v.Len())
				snapshot := *v
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return ioutil.NopCloser(&r), nil
				}
			default:
				// This is where we'd set it to -1 (at least
				// if body != NoBody) to mean unknown, but
				// that broke people during the Go 1.8 testing
				// period. People depend on it being 0 I
				// guess. Maybe retry later. See Issue 18117.
			}
		}
	}
}

// Form Post 请求
func Form(params map[string]string, files ...*os.File) Option {
	bodyBuffer := &bytes.Buffer{}
	contentType := ContentTypeForm

	// 文件 form 提交
	if len(files) > 0 {
		bodyWriter := multipart.NewWriter(bodyBuffer)

		for _, file := range files {
			part, err := bodyWriter.CreateFormFile("files", file.Name())
			app.Panic(err)

			_, err = io.Copy(part, file)
			app.Panic(err)

			err = file.Close()
			app.Panic(err)
		}

		if params != nil {
			// 其他参数列表写入 body
			for k, v := range params {
				if err := bodyWriter.WriteField(k, v); err != nil {
					panic(err)
				}
			}
		}

		contentType = ContentType(bodyWriter.FormDataContentType())
		err := bodyWriter.Close()
		app.Panic(err)
	} else {
		// 普通 form 提交
		urlValues := url.Values{}
		for key, val := range params {
			urlValues.Set(key, val)
		}
		reqBody := urlValues.Encode()
		bodyBuffer.WriteString(reqBody)
	}

	return func(request *http.Request) {
		contentType(request)
		Body(bodyBuffer)(request)
	}
}

// JSONBody Post 请求
func JSONBody(body interface{}) Option {
	jsonBytes, _ := json.Marshal(body)
	bodyReader := strings.NewReader(string(jsonBytes))
	return func(request *http.Request) {
		ContentTypeJSON(request)
		Body(bodyReader)(request)
	}
}

// ContentType 设置 http.Request 请求类型
func ContentType(t string) Option {
	return func(request *http.Request) {
		request.Header.Set("Content-Type", t)
	}
}

// ContentTypeJSON 设置 http.Request Content-Type 为 `application/json`
var ContentTypeJSON = ContentType("application/json")

// ContentTypeForm 设置 http.Request Content-Type 为 `application/x-www-form-urlencoded`
var ContentTypeForm = ContentType("application/x-www-form-urlencoded")


// Cookies 设置 http.Request cookie
func Cookies(cookies ...*http.Cookie) Option {
	return func(request *http.Request) {
		for _, cookie := range cookies {
			request.AddCookie(cookie)
		}
	}
}
