package quotation

import (
	"../utils"
	"reflect"
	"fmt"
	"../structs"
	"strconv"
)

func Main(setting structs.Setting)  {
	history(setting)
	
}

//获取历史100数据
func history(setting structs.Setting) {
	data := utils.Get("/api/v5/market/history-candles?instId=" + setting.InstId)

	var ma5 []float64
	var ma10 []float64

	for _, origin := range data{
		switch reflect.TypeOf(origin).Kind() {
			case reflect.Slice, reflect.Array:
				s := reflect.ValueOf(origin)
				// for i := 0; i < s.Len(); i++ {
				// 	fmt.Println(s.Index(i))
				// }

				tt,_ := strconv.ParseFloat(s.Index(1).Interface().(string), 32)
				ma5 = append(ma5, tt)
				ma10 = append(ma10, tt)

			case reflect.String:
				s := reflect.ValueOf(origin)
				fmt.Println(s.String(), "I am a string type variable.")
			case reflect.Int:
				s := reflect.ValueOf(origin)
				t := s.Int()
				fmt.Println(t, " I am a int type variable.")
		}
	}

	fmt.Println(ma5)
}

//获取当前数据
func ticker(setting structs.Setting) []interface{} {
	return utils.Get("/api/v5/market/ticker?instId=" + setting.InstId)

	// fmt.Println(resp.Data[0]["instType"])
	// fmt.Println(reflect.TypeOf(resp.Data[0]))
}

//深度
func books(setting structs.Setting) []interface{} {
	return utils.Get("/api/v5/market/books?instId=" + setting.InstId)
}