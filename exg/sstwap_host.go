package exg

import (
	"context"

	"github.com/banbox/banexg/errs"
	"github.com/banbox/banexg/sstwap"
)

// unwrapSST returns the inner SSTWAP exchange when Default/BotExchange wraps it.
// Strategies assert *sstwap.Exchange against exg.Default — that fails on *BotExchange
// and silently disables pending-tx / laneState / live_execution gates.
func (e *BotExchange) unwrapSST() (*sstwap.Exchange, bool) {
	if e == nil || e.BanExchange == nil {
		return nil, false
	}
	se, ok := e.BanExchange.(*sstwap.Exchange)
	return se, ok
}

func (e *BotExchange) LiveExecution() bool {
	se, ok := e.unwrapSST()
	if !ok {
		return false
	}
	return se.LiveExecution()
}

func (e *BotExchange) LaneHasPendingTx(laneID uint64) bool {
	se, ok := e.unwrapSST()
	if !ok {
		return false
	}
	return se.LaneHasPendingTx(laneID)
}

func (e *BotExchange) LanePendingTxHash(laneID uint64) (string, bool) {
	se, ok := e.unwrapSST()
	if !ok {
		return "", false
	}
	return se.LanePendingTxHash(laneID)
}

func (e *BotExchange) LaneState(ctx context.Context, laneID uint64) (*sstwap.LaneState, *errs.Error) {
	se, ok := e.unwrapSST()
	if !ok {
		return nil, errs.NewMsg(errs.CodeParamInvalid, "exchange is not sstwap")
	}
	return se.LaneState(ctx, laneID)
}
