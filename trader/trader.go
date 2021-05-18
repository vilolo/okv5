package trader

import (
	"fmt"
	"../structs"
	"../utils"
)

func Open1(setting structs.Setting, side string, posSide string){
	order(setting.InstId, "cross", side, "market", setting.Sz, posSide, "test_long")
	// orderInfo(setting.InstId, "bb")

	//行情分析后得到是否下单
	//是下单
	//判断是否有相反的订单，有则平仓
	//查询是否有未成交的订单
		//如果有取消订单
	//查询是否已下订单
		//已下了就不再下了
		//没有则下订单
	
	fmt.Println("下单完成")
}

func Close1(setting structs.Setting, posSide string){
	//查看持仓信息接口
	//盈利超过20%，基础条件
	//k线反向突破5日均线，出局
	//或亏损超过20%出局
	
	closePosition(setting.InstId, posSide, setting.TdMode)
	fmt.Println("平仓完成")
}

// tdMode	String	是	交易模式
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
func order(instId string, tdMode string, side string, ordType string, sz string, posSide string, clOrdId string){
	data := map[string]string {"instId":instId,"tdMode":tdMode,"side":side,"ordType":ordType,"sz":sz, "posSide":posSide, "clOrdId":clOrdId}
	utils.Post("/api/v5/trade/order", data)
	// fmt.Println(res)
}

func orderInfo(instId string, clOrdId string){
	res := utils.Get("/api/v5/trade/order?instId="+instId+"&clOrdId="+clOrdId)
	fmt.Println(res)
}

// mgnMode	String	是	保证金模式	全仓：cross ； 逐仓： isolated
func closePosition(instId string,posSide string,mgnMode string){
	data := map[string]string {"instId":instId,"posSide":posSide,"mgnMode":mgnMode}
	utils.Post("/api/v5/trade/close-position", data)
	// fmt.Println(res)
}