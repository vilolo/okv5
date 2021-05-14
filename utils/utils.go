package utils

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"../structs"
	"os"
	"encoding/json"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	// "reflect"
)

var conf = structs.Conf{}
var baseUrl = "https://www.okex.win"

func init(){
	file, _ := os.Open("config.json")
    defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
    if err != nil {
        fmt.Println("Error:", err)
    }
}

func Get(url string) []map[string]interface{} {
	client := &http.Client{}
    reqest, err := http.NewRequest("GET", baseUrl+url, nil)

	if err != nil {
        fmt.Println(err)
    }

	conf.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	sign := hmacSha256(conf.Timestamp + "GET" + url, conf.SecretKey)

	//增加header选项
    reqest.Header.Add("OK-ACCESS-KEY", conf.APIKey)
    reqest.Header.Add("OK-ACCESS-SIGN", sign)
    reqest.Header.Add("OK-ACCESS-TIMESTAMP", conf.Timestamp)
	reqest.Header.Add("OK-ACCESS-PASSPHRASE", conf.Passphrase)

	//处理返回结果
    response, err := client.Do(reqest)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return nil
	}
    defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
    //fmt.Println(string(body))
    // fmt.Printf("Get request result: %s\n", string(body))
	
	resp := structs.APIResponse{}
	err = json.Unmarshal([]byte(string(body)), &resp)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return nil
	}

	return resp.Data
}

func hmacSha256(data string, secret string) string {
    h := hmac.New(sha256.New, []byte(secret))
    h.Write([]byte(data))
	// sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}