package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	path = flag.String("path", "", "yapi-import.json所在文件夹")
	token = flag.String("token", "", "yapi项目token")
)

type config struct {
	Type   string `json:"type"`
	Token  string `json:"token"`
	File   string `json:"file"`
	Merge  string `json:"merge"`
	Server string `json:"server"`
}

// main 上传swagger.json到yapi
func main() {
	flag.Parse()
	if *path == "" {
		fmt.Println("path cannot be empty!")
		return
	}
	configPath := *path + "/yapi-import.json"
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("read file[%s] failed err=%v", configPath, err)
		return
	}
	var config config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Printf("read file[%s] Unmarshal json failed err=%v content=%s", *path, err, string(bytes))
		return
	}
	fmt.Printf("load config:%s \ncontent: %+v \n", configPath, config)
	upload(config)
}

func upload(config config) {
	files := strings.Split(config.File, ",")
	for _, f := range files {
		post(config, f)
	}
}

func post(config config, fileName string) {
	filePath := *path + "/" + fileName
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file[%s] failed err=%v", filePath, err)
		return
	}
	actualToken := config.Token
	if config.Token == "" {
		actualToken = *token
	}
	values := url.Values{
		"type":  {config.Type},
		"json":  {string(bytes)},
		"merge": {config.Merge},
		"token": {actualToken},
	}
	client := &http.Client{}
	uri := config.Server + "/api/open/import_data"
	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Printf("http.NewRequest failed err=%v", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client.Do failed err=%v", err)
		return
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read responde failed err=%v", err)
		return
	}
	fmt.Println(string(result))
}
