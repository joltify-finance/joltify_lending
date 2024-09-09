package cmd_test

import (
	"testing"

	"github.com/joltify-finance/joltify_lending/cmd/joltify/cmd"
	"github.com/stretchr/testify/require"
)

func TestMinGasPrice(t *testing.T) {
	require.Equal(t,
		"0.025ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3,25000000000adv4tnt",
		cmd.MinGasPrice,
	)
}
