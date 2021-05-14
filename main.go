package main

import (
	"fmt"
	"./quotation"
	"./structs"
	// "errors"
)

var setting = structs.Setting{}

func main()  {

	//抓取 panic 抛出的异常
	// defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
    //     if err:=recover();err!=nil{
    //         fmt.Println(err) // 这里的err其实就是panic传入的内容，55
    //     }
    // }()

	fmt.Println("Start >>>")
	initSetting()

	quotation.Main(setting)

	fmt.Println("<<< End")
}

func initSetting()  {
	setting.InstId = "BTC-USD-SWAP"
}