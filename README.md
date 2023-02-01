# 简易在线词典
## 概述
我选取了彩云小译和火山翻译的两个引擎，输入要查询的单词后返回该单词的音标、词性及意思，并标注来自哪个翻译引擎。

采用了Goroutine，两个引擎并行，提升了翻译速度。
由下图响应速度对比可得，并发后响应速度提升明显。

![uTools_1673798979137](https://www.kkkode.top/wp-content/uploads/2023/01/uTools_1673798979137.png)

**运行样例**
```bash
go run main.go abandon
```

## 具体步骤
-  抓包

进入官网[彩云小译 - 在线翻译 (caiyunapp.com)](https://fanyi.caiyunapp.com/#/)

右键检查-Network-dict (请求 URL: https://api.interpreter.caiyunai.com/v1/dict  请求方法: POST)

![uTools_1673778139893](https://www.kkkode.top/wp-content/uploads/2023/01/uTools_1673778139893.png)

- curl生成Go代码 [Convert curl to Go (curlconverter.com)](https://curlconverter.com/go/)

![uTools_1673778815476](https://www.kkkode.top/wp-content/uploads/2023/01/uTools_1673778815476.png)

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"trans_type":"en2zh","source":"abandon"}`)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
```

- 生成request body

```go
type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}
```

- 解析response body[JSON转Golang Struct - 在线工具 - OKTools](https://oktools.net/json2go)

```go
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}
```

- 输出想要的内容

```go
err = json.Unmarshal(bodyText, &dictResponse)
if err != nil {
    log.Fatal(err)
}
fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
for _, item := range dictResponse.Dictionary.Explanations {
    fmt.Println(item)
}
fmt.Println("By CaiYun")
```
