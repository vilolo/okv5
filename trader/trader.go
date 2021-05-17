package trader

import (
	"fmt"
	"../structs"
	"../utils"
)

//tdMode	String	是	交易模式
						// 保证金模式：isolated：逐仓 ；cross：全仓
						// 非保证金模式：cash：非保证金
// side	String	是	订单方向 buy：买 sell：卖
// ordType	String	是	订单类型
						// market：市价单
						// limit：限价单
						// post_only：只做maker单
						// fok：全部成交或立即取消
						// ioc：立即成交并取消剩余
// sz	String	是	委托数量
func Td1(setting structs.Setting){
	order(setting.InstId, "cross", "buy", "market", "1")
}

func order(instId string, tdMode string, side string, ordType string, sz string){
	data := map[string]string {"instId":instId,}
	res := utils.Post("/api/v5/trade/order", data)
	fmt.Println(res)
}