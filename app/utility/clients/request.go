package clients

import (
	"context"
	"gin-api/app/utility/system"
	"log"
	"net/http"
	"strings"
	"time"
)

// HTTPClient http.HTTPClient 封装
type HTTPClient struct {
	client http.Client
}

// HTTPClientOption HTTPClient 配置
type HTTPClientOption func(*HTTPClient)

// Timeout 超时时间
func Timeout(timeout time.Duration) HTTPClientOption {
	return func(client *HTTPClient) {
		client.client.Timeout = timeout
	}
}

// NewHTTPClient 返回一个 http 请求客户端
func NewHTTPClient(options ...HTTPClientOption) *HTTPClient {
	client := &HTTPClient{
		client: http.Client{
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second * 5,
		},
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}


// HTTPRequest http.Request 构造器
func HTTPRequest(method, url string, options ...Option) (*http.Request, error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		return nil, err
	}
	for _, opt := range options {
		opt(req)
	}

	return req, nil
}


// Request 发送请求
func (hc *HTTPClient) Request(c context.Context, method, url string, options ...Option) (*ResponseProxy, error) {
	return hc.do(c, method, url, options)
}

// Post 发送 Post 请求
func (hc *HTTPClient) Post(c context.Context, url string, options ...Option) (*ResponseProxy, error) {
	return hc.do(c, http.MethodPost, url, options, ContentTypeJSON)
}

// Get 发送 Get 请求
func (hc *HTTPClient) Get(c context.Context, url string, options ...Option) (*ResponseProxy, error) {
	return hc.do(c, http.MethodGet, url, options)
}

// Delete 发送 Delete 请求
func (hc *HTTPClient) Delete(c context.Context, url string, options ...Option) (*ResponseProxy, error) {
	return hc.do(c, http.MethodDelete, url, options, ContentTypeJSON)
}

// Put 发送 Put 请求
func (hc *HTTPClient) Put(c context.Context, url string, options ...Option) (*ResponseProxy, error) {
	return hc.do(c, http.MethodPut, url, options, ContentTypeJSON)
}

// Patch 发送 Patch 请求
func (hc *HTTPClient) Patch(c context.Context, url string, options ...Option) (*ResponseProxy, error) {
	return hc.do(c, http.MethodPatch, url, options, ContentTypeJSON)
}

// do 内部使用，设置 http 请求的默认配置
func (hc *HTTPClient) do(c context.Context, method, url string, options []Option, defaultOpts ...Option) (*ResponseProxy, error) {
	req, err := HTTPRequest(method, url, defaultOpts...)
	if err != nil {
		return nil, err
	}
	for _, opt := range options {
		opt(req)
	}
	if req.Body != nil {
		defer req.Body.Close()
	}

	var resp *http.Response
	system.TimeLog(func() {
		resp, err = hc.client.Do(req)
	}).Observe(func(s, e time.Time, d time.Duration) {
		log.Printf("[%s]%s 请求时长 %dms", req.Method, req.URL.String(), d/time.Millisecond)
	})

	if err != nil {
		return nil, err
	}
	return ResponseWrapper(c, req, resp), nil
}