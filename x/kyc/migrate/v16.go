package v16

import (
	"encoding/base64"
	"encoding/binary"

	"cosmossdk.io/store/prefix"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/gogoproto/proto"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/joltify-finance/joltify_lending/x/kyc/types"
)

var KeyProjects = []byte("projectInfo")

func setProject(ctx context.Context, storeKey storetypes.StoreKey, p *types.ProjectInfo, cdc codec.BinaryCodec) (int32, error) {
	var currentNum uint32
	projectStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ProjectInfoPrefix))
	projectNum := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ProjectInfoNum))
	data := projectNum.Get(types.KeyPrefix(types.ProjectInfoNum))
	if data == nil {
		currentNum = 1
	} else {
		currentNum = binary.BigEndian.Uint32(data)
		currentNum += 1
	}
	p.Index = int32(currentNum)
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, currentNum)
	projectNum.Set(types.KeyPrefix(types.ProjectInfoNum), bs)
	previousProject := projectStore.Get(types.KeyPrefix(string(p.Index)))
	if previousProject != nil {
		panic("should never happen: two projects with the same index")
	}
	ctx.Logger().Info("setProject", "project index", p.Index)
	projectStore.Set(types.KeyPrefix(string(p.Index)), cdc.MustMarshal(p))
	return int32(currentNum), nil
}

func MigrateStore(ctx context.Context, paramstore paramtypes.Subspace, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	var projectsEncoded string
	paramstore.Get(ctx, KeyProjects, &projectsEncoded)

	val, err := base64.StdEncoding.DecodeString(projectsEncoded)
	if err != nil {
		panic("invalid encoded string")
	}

	var projects types.Projects
	err = proto.Unmarshal(val, &projects)
	if err != nil {
		panic("invalid encoded string")
	}

	for _, el := range projects.Items {
		_, err := setProject(ctx, storeKey, el, cdc)
		ctx.Logger().Info("migrate", "index", el.Index, "err", err)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
