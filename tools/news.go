package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	WEIBO    = "http://api.tianapi.com/weibohot/index?key="
	INTERNET = "http://api.tianapi.com/internet/index?key="
	nCOV     = "http://api.tianapi.com/ncov/index?key="
	HISTORY  = "http://api.tianapi.com/lishi/index?key="
	newDay   = "http://api.tianapi.com/bulletin/index?key="
)

var key string
var urlMap map[string]string

func main() {
	flag.StringVar(&key, "wb", "", "微博热搜")
	flag.StringVar(&key, "net", "", "互联网新闻")
	flag.StringVar(&key, "nCov", "", "疫情信息")
	flag.StringVar(&key, "history", "", "历史的今天")
	flag.StringVar(&key, "newDay", "", "今日简报")
	flag.StringVar(&key, "dy", "", "抖音视频榜")
	flag.Parse()
	tx(key)
}

func tx(apiKey string) {
	args := os.Args
	m := args[1][1:]
	log.Printf("os args[1]/tag is :%s\n", m)
	if m == "history" {
		history()
		return
	}
	val := urlMap[m]
	if val == "" {
		log.Fatalf("input tag: %s is absent\n!", m)
	}
	realUrl := val + apiKey
	req, _ := http.NewRequest("GET", realUrl, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("invoke  fail: %s", err.Error())
	}
	log.Println(resp)
	body, _ := io.ReadAll(resp.Body)
	var prettyJson bytes.Buffer
	parseErr := json.Indent(&prettyJson, body, "", "   ")
	if parseErr != nil {
		log.Printf("Parse tag:%s resp err:%s", m, parseErr.Error())
	} else {
		log.Printf("Fetch [%s] news success!\n%s", m, string(prettyJson.Bytes()))
	}
}

func history() {
	now := time.Now()
	month := now.Format("01")
	day := now.Format("02")
	var buf bytes.Buffer
	buf.WriteString(month)
	buf.WriteString(day)
	today := buf.String()
	fmt.Printf("历史上的今天:%s\n", today)
	postValue := url.Values{"key": {key}, "date": {today}}
	res, _ := http.PostForm(HISTORY, postValue)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var prettyJson bytes.Buffer
	parseErr := json.Indent(&prettyJson, body, "", "   ")
	if parseErr != nil {
		log.Printf("Parse tag:%s resp err:%s", "history", parseErr.Error())
	} else {
		log.Printf("Fetch [%s] news success!\n%s", "history", string(prettyJson.Bytes()))
	}
}

func init() {
	urlMap = make(map[string]string, 4)
	urlMap["wb"] = WEIBO
	urlMap["net"] = INTERNET
	urlMap["nCov"] = nCOV
	urlMap["newDay"] = newDay
	keys := make([]string, 0)
	for k := range urlMap {
		keys = append(keys, k)
	}
	keysJoin := strings.Join(keys, ",")
	log.Printf("urlMap init over, keys: %v", keysJoin)
}
