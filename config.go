package resock

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"sync"
)

type config struct {
	Server   string
	Client   string
	Protocol string
	SNI      string
	Key      []byte
}

var (
	cfg *config = &config{}
	so  sync.Once
)

func GetCfg() *config {
	so.Do(func() {
		if file, err := os.ReadFile("cfg.json"); err == nil {
			json.Unmarshal(file, cfg)
			fmt.Println("Init config")
		}
	})
	return cfg
}

func GenCfgExample() {
	//example
	b, _ := json.MarshalIndent(&config{
		Server:   "0.0.0.0:443",
		Client:   "0.0.0.0:1080",
		Protocol: "tcp",
		SNI:      "http://mirror.centos.org/",
		Key:      GenKey("🕳")},
		" ", " ")
	os.WriteFile("cfg.json", b, fs.ModePerm)
}
