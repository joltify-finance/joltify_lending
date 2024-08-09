package types

import (
	gometrics "github.com/hashicorp/go-metrics"
	"github.com/joltify-finance/joltify_lending/dydx_helper/lib/metrics"
)

// GetMetricLabels returns a slice of gometrics labels for a match perpetual liquidation.
// Currently, the only label is the perpetual id.
func (m MatchPerpetualLiquidation) GetMetricLabels() []gometrics.Label {
	return []gometrics.Label{
		metrics.GetLabelForIntValue(
			metrics.PerpetualId,
			int(m.PerpetualId),
		),
	}
}
