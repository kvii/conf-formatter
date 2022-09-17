package main

import (
	"encoding/json"
)

// v2ray 配置。参考网址: https://www.v2fly.org/config/outbounds.html
type V2ray struct {
	Tag            string         `json:"tag"`            // 配置标识
	Protocol       string         `json:"protocol"`       // 协议名称
	Settings       Settings       `json:"settings"`       // 协议配置
	StreamSettings StreamSettings `json:"streamSettings"` // 底层传输配置
}

func (v *V2ray) String() string {
	bs, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// 协议配置
type Settings struct {
	VNext []VNext `json:"vnext"` // 服务器配置数组
}

// 服务器配置
type VNext struct {
	Address string `json:"address"` // 服务器地址。支持 ip 和域名
	Port    int    `json:"port"`    // 服务器端口号
	Users   []User `json:"users"`   // 用户数组
}

// 用户数组
type User struct {
	Id string `json:"id"` // VMess 用户的主 ID。必须是一个合法的 UUID。
}

// 底层传输配置
type StreamSettings struct {
	Network    string     `json:"network"`    // 数据流所使用的网络类型，默认值为 "tcp"
	Security   string     `json:"security"`   // 是否启用传输层加密.支持的选项有 "none" 表示不加密（默认值），"tls" 表示使用 TLSopen in new window。
	WsSettings WsSettings `json:"wsSettings"` // WebSocket 配置，仅当此连接使用 WebSocket 时有效。
}

// WebSocket 配置，仅当此连接使用 WebSocket 时有效。
type WsSettings struct {
	Path string `json:"path"` // WebSocket 所使用的 HTTP 协议路径，默认值为 "/"。
}
