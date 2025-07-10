package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func query(word string) {
	// 创建HTTP客户端实例（用于发送请求）
	client := &http.Client{}

	// 构造API请求结构体（指定翻译类型为英→中，传入待查询单词）
	request := DictRequest{TransType: "en2zh", Source: word}

	// 将请求结构体序列化为JSON格式（用于请求体）
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err) // 序列化失败时终止程序并输出错误
	}

	// 将JSON字节数组转换为io.Reader接口（满足http.NewRequest的参数要求）
	var data = bytes.NewReader(buf)

	// 创建POST请求（目标为彩云翻译词典API）
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err) // 请求创建失败时终止程序
	}

	// 设置请求头（模拟浏览器请求，满足API鉴权和格式要求）
	req.Header.Set("Connection", "keep-alive")                                                                                                               // 保持长连接
	req.Header.Set("DNT", "1")                                                                                                                               // 不跟踪标识
	req.Header.Set("os-version", "")                                                                                                                         // 操作系统版本（空值）
	req.Header.Set("sec-ch-ua-mobile", "?0")                                                                                                                 // 非移动设备标识
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36") // 浏览器UA标识
	req.Header.Set("app-name", "xy")                                                                                                                         // 应用名称
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")                                                                                         // 请求体格式
	req.Header.Set("Accept", "application/json, text/plain, */*")                                                                                            // 接受的响应格式
	req.Header.Set("device-id", "")                                                                                                                          // 设备ID（空值）
	req.Header.Set("os-type", "web")                                                                                                                         // 系统类型（网页端）
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")                                                                                          // API认证令牌
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")                                                                                                  // 请求来源域名
	req.Header.Set("Sec-Fetch-Site", "cross-site")                                                                                                           // 跨站请求标识
	req.Header.Set("Sec-Fetch-Mode", "cors")                                                                                                                 // CORS模式
	req.Header.Set("Sec-Fetch-Dest", "empty")                                                                                                                // 请求目标类型
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")                                                                                                // 跳转来源
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")                                                                                                      // 接受的语言
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")                                                                               // 会话Cookie

	// 发送HTTP请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err) // 请求发送失败时终止程序
	}
	// 延迟关闭响应体（防止资源泄露）
	defer func(Body io.ReadCloser) {
		_ = Body.Close() // 忽略关闭错误（实际生产环境需处理）
	}(resp.Body)

	// 读取响应体全部内容（转换为字节数组）
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err) // 读取响应失败时终止程序
	}

	// 检查HTTP状态码（非200表示请求异常）
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText)) // 输出异常信息
	}

	// 将响应JSON反序列化为自定义结构体（便于解析数据）
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err) // 反序列化失败时终止程序
	}

	// 打印单词的英式/美式发音（从响应结构体中提取）
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)

	// 遍历并打印单词释义列表（每个释义为字符串）
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	fmt.Println("请输入要查询的单词：")
	var word string
	_, err := fmt.Scanln(&word)
	if err != nil {
		return
	}
	query(word)
}
