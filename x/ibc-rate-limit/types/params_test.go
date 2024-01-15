package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
)

func TestValidateQuota(t *testing.T) {
	testCases := map[string]struct {
		tokenQuota interface{}
		expected   bool
	}{
		// ToDo: Why do tests expect the bech32 prefix to be cosmos?
		"valid_quota": {
			tokenQuota: "100uatom,1000000ujolt",
			expected:   true,
		},
		"invalid_quota": {
			tokenQuota: "cosmos1234",
			expected:   false,
		},
		"valid_quota_2": {
			tokenQuota: "1000000ujolt,100uatom",
			expected:   true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateQuota(tc.tokenQuota)

			// Assertions.
			if !tc.expected {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func SetupBech32Prefix() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("jolt", "joltpub")
	config.SetBech32PrefixForValidator("joltvaloper", "joltvpub")
	config.SetBech32PrefixForConsensusNode("joltvalcons", "joltcpub")
}

func TestValidateParams(t *testing.T) {
	SetupBech32Prefix()

	testCases := map[string]struct {
		tokenQuota interface{}
		whitelist  interface{}
		expected   bool
	}{
		// ToDo: Why do tests expect the bech32 prefix to be cosmos?
		"valid_params": {
			tokenQuota: "1000uatom,1000ujolt",
			whitelist:  []string{},
			expected:   true,
		},
		"invalid_white_list": {
			tokenQuota: "1000uatom,1000ujolt",
			whitelist:  []string{"cosmos1234"},
			expected:   false,
		},
		"valid params": {
			tokenQuota: "1000uatom,1000ujolt",
			whitelist:  []string{"jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"},
			expected:   true,
		},
		"empty quota": {
			tokenQuota: "",
			whitelist:  []string{"jolt1txtsnx4gr4effr8542778fsxc20j5vzqxet7t0"},
			expected:   false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tokenQuota, ok := tc.tokenQuota.(string)
			require.True(t, ok, "unexpected type of sorted token")

			whitelist, ok := tc.whitelist.([]string)
			require.True(t, ok, "unexpected type of whitelist")

			params := Params{
				TokenQuota: tokenQuota,
				Whitelist:  whitelist,
			}

			err := params.Validate()

			// Assertions.
			if !tc.expected {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
