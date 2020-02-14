package hconfig

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// 加载配置
func LoadConfig(client *http.Client, name string) ([]byte, error) {
	if client == nil {
		client = http.DefaultClient
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
