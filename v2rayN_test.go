package main

import (
	_ "embed"
	"testing"
)

//go:embed v2rayN.json
var v2rayN []byte

func TestV2rayN(t *testing.T) {
	want := V2rayN{
		Version:    "2",
		Alias:      "vmess",
		Addr:       "www.host.com",
		Port:       "443",
		Uuid:       "uuid",
		AlterId:    "0",
		Security:   "",
		Net:        "ws",
		Type:       "none",
		Host:       "www.host.com",
		Path:       "/path",
		Tls:        "tls",
		ServerName: "",
	}

	u := "vmess://eyJ2IjoiMiIsInBzIjoidm1lc3MiLCJhZGQiOiJ3d3cuaG9zdC5jb20iLCJwb3J0IjoiNDQzIiwiaWQiOiJ1dWlkIiwiYWlkIjoiMCIsIm5ldCI6IndzIiwidHlwZSI6Im5vbmUiLCJob3N0Ijoid3d3Lmhvc3QuY29tIiwicGF0aCI6Ii9wYXRoIiwidGxzIjoidGxzIiwic25pIjoiIn0="
	v := V2rayNFromUrl(u)
	if v != want {
		t.Fatalf("v2rayN: parse: %#v", v)
	}

	j := V2rayNFromJson(v2rayN)
	if j != want {
		t.Fatalf("v2rayN: parse json: %#v", j)
	}

	s := want.String()
	if s != u {
		t.Fatalf("v2rayN: format: %s", s)
	}
}
