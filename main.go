package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

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

type DictRequest2 struct {
	Source         string   `json:"source"`
	Words          []string `json:"words"`
	SourceLanguage string   `json:"source_language"`
	TargetLanguage string   `json:"target_language"`
}

type DictResponse2 struct {
	Details []struct {
		Detail string `json:"detail"`
		Extra  string `json:"extra"`
	} `json:"details"`
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
}

type Detail struct {
	ErrorCode string `json:"errorCode"`
	RequestID string `json:"requestId"`
	Msg       string `json:"msg"`
	Result    []struct {
		Ec struct {
			ReturnPhrase []string `json:"returnPhrase"`
			Web          []struct {
				Phrase   string   `json:"phrase"`
				Meanings []string `json:"meanings"`
			} `json:"web"`
			Synonyms []struct {
				Pos   string   `json:"pos"`
				Words []string `json:"words"`
				Trans string   `json:"trans"`
			} `json:"synonyms"`
			Etymology struct {
				ZhCHS []struct {
					Description string `json:"description"`
					Detail      string `json:"detail"`
					Source      string `json:"source"`
				} `json:"zh-CHS"`
			} `json:"etymology"`
			MTerminalDict string `json:"mTerminalDict"`
			Dict          string `json:"dict"`
			Basic         struct {
				UsPhonetic string   `json:"usPhonetic"`
				UsSpeech   string   `json:"usSpeech"`
				Phonetic   string   `json:"phonetic"`
				UkSpeech   string   `json:"ukSpeech"`
				ExamType   []string `json:"examType"`
				Explains   []struct {
					Pos   string `json:"pos"`
					Trans string `json:"trans"`
				} `json:"explains"`
				UkPhonetic  string `json:"ukPhonetic"`
				WordFormats []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"wordFormats"`
			} `json:"basic"`
			Phrases []struct {
				Phrase   string   `json:"phrase"`
				Meanings []string `json:"meanings"`
			} `json:"phrases"`
			Lang           string `json:"lang"`
			SentenceSample []struct {
				Sentence     string `json:"sentence"`
				SentenceBold string `json:"sentenceBold"`
				Translation  string `json:"translation"`
			} `json:"sentenceSample"`
			IsWord  bool   `json:"isWord"`
			WebDict string `json:"webDict"`
		} `json:"ec"`
	} `json:"result"`
}

func query(word string) {
	client := &http.Client{}
	request := DictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
	fmt.Println("By CaiYun")
	fmt.Println()
}

func query2(word string) {
	client := &http.Client{}
	request := DictRequest2{Source: "youdao", Words: []string{word}, SourceLanguage: "en", TargetLanguage: "zh"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/detail/v1/?msToken=&X-Bogus=DFSzswVLQDadxLyJSZW76tteJnyE&_signature=_02B4Z6wo00001ljRzmgAAIDCU11wL0iWDHZY0crAAPYG0kiMNFxedoitm5HLGUD0jZ5DcXIetnNY8ZGiALMbzgBDv1KOARzAc9zVX0hBbkTuGlfrQXxOs7iKnzDbkKrVsENB0LBFxbQpr7cx6d", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16737825730907524; i18next=zh-CN; s_v_web_id=verify_lcxaxfl8_82b7bTsx_spXo_4ojy_8eRx_vYijmz0YGEDK; ttcid=53e7468bfdd8463d9035633a3619cd1e27; tt_scid=kFfR0z7TRSKbDMjd5nuEbx.LtFvHzHFpDi1yQWzIA0zpggD470g9zU3L.lxerpe6c31c")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("referer", "https://translate.volcengine.com/mobile?category=&home_language=zh&source_language=detect&target_language=zh&text=hello")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dictResponse2 DictResponse2
	var detail Detail
	err = json.Unmarshal(bodyText, &dictResponse2)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(dictResponse2.Details[0].Detail), &detail)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v", detail)
	fmt.Println(word, "UK:", detail.Result[0].Ec.Basic.Phonetic, "US:", detail.Result[0].Ec.Basic.UsPhonetic)
	for _, item := range detail.Result[0].Ec.Basic.Explains {
		fmt.Println(item.Pos, item.Trans)
	}
	fmt.Println("By HuoShan")
	fmt.Println()
}

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	words := os.Args[1:]
	var word string
	for i, item := range words {
		if i != len(words)-1 {
			word += item + " "
		} else {
			word += item
		}
	}
	go func() {
		query(word)
	}()
	query2(word)
	time.Sleep(time.Second)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
