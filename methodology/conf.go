package methodology

import (
	_ "embed"

	"github.com/BurntSushi/toml"
)

//go:embed default.toml
var defaultConfEmbedded string

func mustLoadDefaultConf() map[string]*Methodology {
	conf := make(map[string]*Methodology)
	// 解析toml文件
	_, err := toml.Decode(defaultConfEmbedded, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
