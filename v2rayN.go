package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// v2rayN 格式分享链接。参考网址: https://github.com/2dust/v2rayN/wiki
//
// 格式 vmess://Base64编码的json
//
//	域名:
//	1. http(tcp)->host中间逗号(,)隔开
//	2. ws->host
//	3. h2->host
//	4. QUIC->security
//
//	路径:
//	1. ws->path
//	2. h2->path
//	3. QUIC->key/Kcp->seed
//	4. grpc->serviceName
type V2rayN struct {
	Version     string `json:"v"`             // 配置文件版本号,主要用来识别当前配置
	Alias       string `json:"ps"`            // 备注或别名
	Addr        string `json:"add"`           // 地址IP或域名
	Port        string `json:"port"`          // 端口号
	Uuid        string `json:"id"`            // UUID
	AlterId     string `json:"aid"`           // alterId
	Security    string `json:"scy,omitempty"` // 加密方式(security),没有时值默认auto
	Net         string `json:"net"`           // 传输协议(tcp\kcp\ws\h2\quic)
	Type        string `json:"type"`          // 伪装类型(none\http\srtp\utp\wechat-video) *tcp or kcp or QUIC
	Host        string `json:"host"`          // 伪装的域名
	Path        string `json:"path"`          // 路径
	Tls         string `json:"tls"`           // 底层传输安全(tls)
	ServerName  string `json:"sni"`           // 服务器名
	Alpn        string `json:"alpn"`
	Fingerprint string `json:"fp"`
}

func (v *V2rayN) String() string {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	s := base64.StdEncoding.EncodeToString(bs)
	return fmt.Sprintf("vmess://%s", s)
}

// 从 v2rayN url 中创建 V2rayN 配置
func V2rayNFromUrl(s string) V2rayN {
	u, err := url.Parse(s)
	panicIf(err)
	if u.Scheme != "vmess" {
		panic(errors.New("协议类型错误"))
	}

	bs, err := base64.StdEncoding.DecodeString(u.Host)
	panicIf(err)

	return V2rayNFromJson(bs)
}

// 从 v2rayN json 中创建 V2rayN 配置
func V2rayNFromJson(bs []byte) (v V2rayN) {
	err := json.Unmarshal(bs, &v)
	panicIf(err)
	return v
}
