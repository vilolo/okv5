package trader

import (
	"fmt"
	"../structs"
	"../utils"
)

func Td1(setting structs.Setting){
	order(setting.InstId, "cross", "buy", "market", "1","long")
}

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
// posSide
	// 持仓方向，单向持仓模式下此参数非必填，如果填写仅可以选择net；在双向持仓模式下必填，且仅可选择 long 或 short。
	// 双向持仓模式下，side和posSide需要进行组合
	// 开多：买入开多（side 填写 buy； posSide 填写 long ）
	// 开空：卖出开空（side 填写 sell； posSide 填写 short ）
	// 平多：卖出平多（side 填写 sell；posSide 填写 long ）
	// 平空：买入平空（side 填写 buy； posSide 填写 short ）
func order(instId string, tdMode string, side string, ordType string, sz string, posSide string){
	data := map[string]string {"instId":instId,"tdMode":tdMode,"side":side,"ordType":ordType,"sz":sz, "posSide":posSide}
	res := utils.Post("/api/v5/trade/order", data)
	fmt.Println(res)
}