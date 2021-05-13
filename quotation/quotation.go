package quotation

import (
	"../utils"
	"reflect"
	"fmt"
	"../structs"
	"encoding/json"
)

func Main()  {
	body := utils.Get("/api/v5/market/ticker?instId=BTC-USD-SWAP")
	fmt.Println(body)
	resp := structs.APIResponse{}
	err := json.Unmarshal([]byte(body), &resp)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Println(resp.Data[0]["instType"])
	fmt.Println(reflect.TypeOf(resp.Data[0]))
}