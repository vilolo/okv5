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
	"log"
	"strings"
	// "reflect"
)

var conf = structs.Conf{}
// var baseUrl = "https://www.okex.win"
var baseUrl = "https://www.ouyi.cc"

func init(){
	file, _ := os.Open("config.json")
    defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
    if err != nil {
        fmt.Println("Error:", err)
    }
}

func Get(url string) []interface{} {
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
	fmt.Println(string(body), "%%%%%%%%%")
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

func Post(url string, data map[string]string) []interface{}{
	client := &http.Client{}

	jsons,_:=json.Marshal(data)
	result :=string(jsons)
	jsoninfo :=strings.NewReader(result)
    reqest, err := http.NewRequest("POST", baseUrl+url, jsoninfo)

	if err != nil {
        fmt.Println(err)
    }
	
	conf.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	sign := hmacSha256(conf.Timestamp + "POST" + url + result, conf.SecretKey)

	//增加header选项
    reqest.Header.Add("OK-ACCESS-KEY", conf.APIKey)
    reqest.Header.Add("OK-ACCESS-SIGN", sign)
    reqest.Header.Add("OK-ACCESS-TIMESTAMP", conf.Timestamp)
	reqest.Header.Add("OK-ACCESS-PASSPHRASE", conf.Passphrase)
	reqest.Header.Add("Content-Type", "application/json")

	//处理返回结果
    response, err := client.Do(reqest)

	if err != nil {
		fmt.Printf("err:%v\n", err)
		return nil
	}
    defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body))
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

func WriteLog(msg interface{}) {
	var path = "./log/"+time.Now().Format("20060102")
	if _,err := os.Stat(path); !(err == nil || os.IsExist(err)) {
		err = CreateDir(path)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	logfile,err:=os.OpenFile(path+"/test.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
    if err!=nil{
        fmt.Printf("%s\r\n",err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    logger:=log.New(logfile,"",log.Ldate|log.Ltime|log.Llongfile)
    logger.Println(msg)
}

//CreateDir  文件夹创建
func CreateDir(path string) error {
    err := os.MkdirAll(path, os.ModePerm)
    if err != nil {
        return err
    }
    os.Chmod(path, os.ModePerm)
    return nil
}