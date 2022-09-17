package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"text/template"
)

//go:embed out.tpl
var out string

func main() {
	v2ray := v2rayFromR(os.Stdin)
	clash := clashFromV2ray(v2ray)
	v2rayN := v2rayNFromV2ray(v2ray)

	data := map[string]string{
		"v2ray":  v2ray.String(),
		"clash":  clash.String(),
		"v2rayN": v2rayN.String(),
	}

	t := template.New("out")
	t = template.Must(t.Parse(out))
	err := t.Execute(os.Stdout, data)
	panicIf(err)
}

func v2rayFromR(r io.Reader) (v V2ray) {
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	panicIf(err)
	return v
}

func clashFromV2ray(v V2ray) Clash {
	return Clash{
		Name:       v.Tag,
		Type:       v.Protocol,
		Server:     v.Settings.VNext[0].Address,
		Port:       v.Settings.VNext[0].Port,
		Uuid:       v.Settings.VNext[0].Users[0].Id,
		AlterId:    0,
		Cipher:     "auto",
		Tls:        v.StreamSettings.Security == "tls",
		ServerName: v.Settings.VNext[0].Address,
		Network:    v.StreamSettings.Network,
		WsOpts: WsOpts{
			Path: v.StreamSettings.WsSettings.Path,
		},
	}
}

func v2rayNFromV2ray(v V2ray) V2rayN {
	return V2rayN{
		Version:    "2",
		Alias:      v.Tag,
		Addr:       v.Settings.VNext[0].Address,
		Port:       strconv.Itoa(v.Settings.VNext[0].Port),
		Uuid:       v.Settings.VNext[0].Users[0].Id,
		AlterId:    "0",
		Security:   v.StreamSettings.Security,
		Net:        v.StreamSettings.Network,
		Type:       "none",
		Host:       v.Settings.VNext[0].Address,
		Path:       v.StreamSettings.WsSettings.Path,
		Tls:        v.StreamSettings.Security,
		ServerName: "",
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
