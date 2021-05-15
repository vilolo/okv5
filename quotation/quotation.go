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
	// data := utils.Get("/api/v5/market/history-candles?instId=" + setting.InstId)

	data := [11][4]string{
		{"1","1.111","1.111","1"},
		{"2","2.111","2.111","2"},
		{"3","3.111","3.111","3"},
		{"4","4.111","4.111","4"},
		{"5","5.111","5.111","5"},
		{"6","6.111","6.111","6"},
		{"7","7.111","7.111","7"},
		{"8","8.111","8.111","8"},
		{"9","9.111","9.111","9"},
		{"10","10.111","10.111","10"},
		{"11","11.111","11.111","11"},
	}
	
	var ma5 []float64
	var ma10 []float64

	for i := len(data)-1; i>=0; i-- {
		s := reflect.ValueOf(data[i])
		curH,_ := strconv.ParseFloat(s.Index(2).Interface().(string), 32)
		curL,_ := strconv.ParseFloat(s.Index(3).Interface().(string), 32)
		cur := (curH + curL)/2

		if i != len(data)-1 {
			//找前面的几根
			var t5 = 1
			var t10 = 1
			var tempVal5 float64
			var tempVal10 float64
			for j := i+1; j <= len(data)-1; j++ {
				if t10 > 9 {
					break
				}

				tempH, _ := strconv.ParseFloat(reflect.ValueOf(data[j]).Index(2).Interface().(string), 32)
				tempL, _ := strconv.ParseFloat(reflect.ValueOf(data[j]).Index(3).Interface().(string), 32)
				tempVal := (tempH + tempL)/2
				
				if t5 <= 4 {
					t5++
					tempVal5 = tempVal5 + tempVal
				}
				t10++
				tempVal10 = tempVal10 + tempVal
			}

			ma5 = append(ma5, (tempVal5+cur)/float64(t5))
			ma10 = append(ma10, (tempVal10+cur)/float64(t10))
		}else{
			ma5 = append(ma5, cur)
			ma10 = append(ma10, cur)
		}
	}

	fmt.Println(ma5)
	fmt.Println(ma10)
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