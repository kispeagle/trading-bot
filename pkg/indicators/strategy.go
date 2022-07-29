package indicator

type StrategyIndicator struct {
	profitRate float64
	lossRate   float64
}

func NewStrategyIndicator(profitRate, lossRate float64) StrategyIndicator {
	i := StrategyIndicator{profitRate, lossRate}
	return i
}

func (i *StrategyIndicator) SetProfit(r float64) {
	i.profitRate = r
}

func (i *StrategyIndicator) SetLoss(r float64) {
	i.lossRate = r
}

func (i StrategyIndicator) ShowStratege(price float64) (float64, float64) {

	profit := price * (1 + i.profitRate)
	loss := price * (1 - i.lossRate)
	return profit, loss
}
