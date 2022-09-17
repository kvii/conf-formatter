package main

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// clash 配置。参考网址: https://github.com/Dreamacro/clash/wiki/Configuration
type Clash struct {
	Name       string `yaml:"name"`       // 代理名称
	Type       string `yaml:"type"`       // 协议类型
	Server     string `yaml:"server"`     // 服务器地址
	Port       int    `yaml:"port"`       // 端口
	Uuid       string `yaml:"uuid"`       // uuid
	AlterId    int    `yaml:"alterId"`    // 额外 id
	Cipher     string `yaml:"cipher"`     // 加密类型
	Tls        bool   `yaml:"tls"`        // 是否开启 tls
	ServerName string `yaml:"servername"` // 服务器名称
	Network    string `yaml:"network"`    // 网络类型
	WsOpts     WsOpts `yaml:"ws-opts"`    // websocket 配置
}

func (v *Clash) String() string {
	bs, err := YamlMarshalIndent(v, 2)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// websocket 配置
type WsOpts struct {
	Path string `yaml:"path"` // 路径
}

// yaml 格式化
func YamlMarshalIndent(v any, indent int) ([]byte, error) {
	buf := new(bytes.Buffer)
	e := yaml.NewEncoder(buf)
	e.SetIndent(indent)

	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
