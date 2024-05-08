package types

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/types"
)

func GenerateTestProjects() []*ProjectInfo {
	acc, err := types.AccAddressFromBech32("jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq")
	if err != nil {
		panic(err)
	}

	allProjects := make([]*ProjectInfo, 100)
	for i := 0; i < 100; i++ {
		b := BasicInfo{
			"This is the test info",
			"empty",
			"ABC",
			"ABC123",
			[]byte("reserved"),
			"This is the Test Project 1",
			"example@example.com",
			"example",
			"empty logo url",
			"empty project Brief",
			"empty project description",
		}
		val, ok := types.NewIntFromString("100000000000000000000")
		if !ok {
			panic("fail to convert")
		}
		pi := ProjectInfo{
			Index:                        int32(i + 1),
			SPVName:                      strconv.Itoa(i) + ":" + "test projects",
			ProjectOwner:                 acc,
			BasicInfo:                    &b,
			ProjectLength:                480, // 5 mins
			SeparatePool:                 true,
			BaseApy:                      types.NewDecWithPrec(10, 2),
			PayFreq:                      "120",
			PoolLockedSeconds:            100,
			PoolTotalBorrowLimit:         100,
			MarketId:                     "aud:usd",
			WithdrawRequestWindowSeconds: 30,
			MinBorrowAmount:              math.NewInt(100),
			MinDepositAmount:             val,
		}
		pi.BasicInfo.ProjectName = fmt.Sprintf("this is the project %v", i)
		allProjects[i] = &pi
	}

	return allProjects
}
