package main

import (
	_ "embed"
	"testing"
)

//go:embed v2rayN.json
var v2rayN []byte

func TestV2rayN(t *testing.T) {
	want := V2rayN{
		Version:     "2",
		Alias:       "vmess",
		Addr:        "www.host.com",
		Port:        "443",
		Uuid:        "uuid",
		AlterId:     "0",
		Security:    "auto",
		Net:         "ws",
		Type:        "none",
		Host:        "www.host.com",
		Path:        "/path",
		Tls:         "tls",
		ServerName:  "",
		Alpn:        "",
		Fingerprint: "",
	}

	u := "vmess://eyJ2IjoiMiIsInBzIjoidm1lc3MiLCJhZGQiOiJ3d3cuaG9zdC5jb20iLCJwb3J0IjoiNDQzIiwiaWQiOiJ1dWlkIiwiYWlkIjoiMCIsInNjeSI6ImF1dG8iLCJuZXQiOiJ3cyIsInR5cGUiOiJub25lIiwiaG9zdCI6Ind3dy5ob3N0LmNvbSIsInBhdGgiOiIvcGF0aCIsInRscyI6InRscyIsInNuaSI6IiIsImFscG4iOiIiLCJmcCI6IiJ9"
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
