package market

import "github.com/mumenma/huobi_golang/pkg/model/base"

type SubscribeTickerResponse struct {
	base.WebSocketResponseBase
	Data *Ticker
	Tick *Ticker
}
