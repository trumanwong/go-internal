package notice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func PushWeChatRobot(level, message, url string) {
	messages, params := make([]string, 0), make(map[string]interface{})
	messages = append(messages, fmt.Sprintf("- 时间：%s", time.Now().Format("2006-01-02 15:04:05")))
	messages = append(messages, fmt.Sprintf("- Level：%s", level))
	messages = append(messages, fmt.Sprintf("- 信息：%s", message))
	markdown := make(map[string]interface{})
	markdown["title"] = "通知"
	markdown["content"] = strings.Join(messages, "\n")
	params["timestamp"] = time.Now().Unix()
	params["msgtype"] = "markdown"
	params["markdown"] = markdown
	data, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		log.Printf("NewRequest fail, %s", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("Request WeChat Api fail, %s", err)
		return
	}
	readAll, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(readAll))
}
