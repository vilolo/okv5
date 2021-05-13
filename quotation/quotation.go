package quotation

import (
	"../utils"
	// "reflect"
	"fmt"
	"../structs"
	"encoding/json"
)

func Main()  {
	body := utils.Get("/api/v5/account/account-position-risk")
	fmt.Println(body)
	ticker := structs.Ticker{}
	err := json.Unmarshal([]byte(body), &ticker)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Println(ticker)
}