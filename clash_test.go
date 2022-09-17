package main

import (
	_ "embed"
	"testing"
)

//go:embed clash.yaml
var clash string

func TestClash(t *testing.T) {
	v := Clash{
		Name:       "vmess",
		Type:       "vmess",
		Server:     "www.host.com",
		Port:       443,
		Uuid:       "uuid",
		AlterId:    0,
		Cipher:     "auto",
		Tls:        true,
		ServerName: "www.host.com",
		Network:    "ws",
		WsOpts: WsOpts{
			Path: "/path",
		},
	}
	s := v.String()
	if s != clash {
		t.Fatalf("clash: format: %s", s)
	}
}
