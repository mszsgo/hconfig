package hconfig

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// 加载配置
func LoadConfig(client *http.Client, name string) ([]byte, error) {
	if client == nil {
		client = http.DefaultClient
	}
	// 如果是local本地开发环境，配置服务名加后缀`-local`，仅用于本地开发环境，发布无需配置此环境变量
	if os.Getenv("MS_ENV") == "local" {
		name = name + "-local"
	}
	resp, err := client.Get("http://config/get?name=" + name)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NowConfig(client *http.Client, name string, v interface{}) error {
	bytes, err := LoadConfig(client, name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}
