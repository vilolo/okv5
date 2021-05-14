package quotation

import (
	"../utils"
	// "reflect"
	"fmt"
	"../structs"
)

func Main(setting structs.Setting)  {
	historyData := history(setting)
	fmt.Println(historyData)
}

//获取历史100数据
func history(setting structs.Setting) []map[string]interface{} {
	return utils.Get("/api/v5/market/history-candles?instId=" + setting.InstId)
}

//获取当前数据
func ticker(setting structs.Setting) []map[string]interface{} {
	return utils.Get("/api/v5/market/ticker?instId=" + setting.InstId)

	// fmt.Println(resp.Data[0]["instType"])
	// fmt.Println(reflect.TypeOf(resp.Data[0]))
}

//深度
func books(setting structs.Setting) []map[string]interface{} {
	return utils.Get("/api/v5/market/books?instId=" + setting.InstId)
}