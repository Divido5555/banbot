package exg

import (
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
	sstwap.EnrichOrderParams(params)
}
