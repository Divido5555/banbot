package exg

import (
	"strconv"

	"github.com/banbox/banbot/config"
	"github.com/banbox/banexg/sstwap"
)

// AttachSstwapOrderParams copies lane routing hints from banbot order metadata into
// banexg CreateOrder params. SSTWAP adapter only — no change to strategy logic.
func AttachSstwapOrderParams(enterTag string, getInfo func(string) string, params map[string]interface{}) {
	if config.Exchange == nil || config.Exchange.Name != sstwap.ExgID {
		return
	}
	if params == nil {
		return
	}
	if enterTag != "" {
		params[sstwap.ParamEnterTag] = enterTag
	}
	if v := getInfo("laneId"); v != "" {
		params[sstwap.ParamLaneID] = v
	}
	// Hub decision mark (DAI/WPLS) for minOut lock — buys (maxPrice) and sells (stop mark).
	if v := getInfo("decisionMarkDAI"); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil && f > 0 {
			params[sstwap.ParamDecisionMark] = f
		}
	}
	sstwap.EnrichOrderParams(params)
}
