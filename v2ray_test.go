package main

import (
	_ "embed"
	"testing"
)

//go:embed v2ray.json
var v2ray string

func TestV2ray(t *testing.T) {
	v := V2ray{
		Tag:      "vmess",
		Protocol: "vmess",
		Settings: Settings{
			VNext: []VNext{
				{
					Address: "www.host.com",
					Port:    443,
					Users: []User{
						{
							Id: "uuid",
						},
					},
				},
			},
		},
		StreamSettings: StreamSettings{
			Network:  "ws",
			Security: "tls",
			WsSettings: WsSettings{
				Path: "/path",
			},
		},
	}

	s := v.String()
	if s != v2ray {
		t.Fatalf("v2ray: format: %s", s)
	}
}
