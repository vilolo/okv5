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

	data := [11][2]string{
		{"1","1"},
		{"2","2"},
		{"3","3"},
		{"4","4"},
		{"5","5"},
		{"6","6"},
		{"7","7"},
		{"8","8"},
		{"9","9"},
		{"10","10"},
		{"11","11"},
	}
	var ma5 []float64
	var ma10 []float64
	var t5 int
	var t10 int
	var tempCount float64
	var tempVal float64
	for i := len(data)-1; i>=0; i-- {
		switch reflect.TypeOf(data[i]).Kind() {
			case reflect.Slice, reflect.Array:
				s := reflect.ValueOf(data[i])
				cur,_ := strconv.ParseFloat(s.Index(1).Interface().(string), 32)

				if i == len(data)-1 {
					ma5 = append(ma5, cur)
					ma10 = append(ma10, cur)
				}else{
					var ma5Temp float64 = 0
					var ma10Temp float64 = 0
					tempCount = 0
					t5 = 4
					t10 = 9
					for j := i+1; j < len(data); j++ {
						if t10 <= 0 {
							break
						}

						tempVal, _ = strconv.ParseFloat(reflect.ValueOf(data[j]).Index(1).Interface().(string), 32)
						tempCount = tempCount + tempVal

						ma10Temp = tempCount

						if t5 == 1 {
							ma5Temp = tempCount
						}

						t5 --
						t10 --
					}

					if ma5Temp == 0 {
						ma5Temp = ma10Temp
					}

					ma5Temp = ma5Temp + cur
					ma10Temp = ma10Temp + cur

					ma5 = append(ma5, ma5Temp)
					ma10 = append(ma10, ma10Temp)
				}

			case reflect.String:
				s := reflect.ValueOf(data[i])
				fmt.Println(s.String(), "I am a string type variable.")
			case reflect.Int:
				s := reflect.ValueOf(data[i])
				t := s.Int()
				fmt.Println(t, " I am a int type variable.")
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