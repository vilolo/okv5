package structs

type Conf struct {
    APIKey string
	SecretKey string
	Passphrase string
	Timestamp string
}

type Setting struct {
    InstId string		//instId	String	是	产品ID，如 BTC-USD-SWAP
	Sz string		//sz	String	是	委托数量
	TdMode string	//mgnMode	String	是	保证金模式	全仓：cross ； 逐仓： isolated
}

type APIResponse struct {
	Code string                   `json:"code"`
	Msg  string                   `json:"msg"`
	Data []interface{} `json:"data"`
}

//获取单个产品行情信息
type Ticker struct {
    InstType	string	//产品类型
	InstId	string	//产品ID
	Last	string	//最新成交价
	LastSz	string	//最新成交的数量
	AskPx	string	//卖一价
	AskSz	string	//卖一价对应的数量
	BidPx	string	//买一价
	BidSz	string	//买一价对应的数量
	Open24h	string	//24小时开盘价
	High24h	string	//24小时最高价
	Low24h	string	//24小时最低价
	VolCcy24h	string	//24小时成交量，以币为单位
						//如果是衍生品合约，//数值为结算货币的数量。
						//如果是币币/币币杠杆，//数值为计价货币的数量。
	Vol24h	string	//24小时成交量，以张为单位
					//如果是衍生品合约，//数值为合约的张数。
					//如果是币币/币币杠杆，//数值为交易货币的数量。
	SodUtc0	string	//UTC 0 时开盘价
	SodUtc8	string	//UTC+8 时开盘价
	Ts	string	//ticker数据产生时间，Unix时间戳的毫秒数格式，如 1597026383085
}