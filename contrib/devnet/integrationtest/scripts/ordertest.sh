#!/bin/bash
set -x
height=$(curl -s http://localhost:26657/status | jq -r '.result.sync_info.latest_block_height')
heightsubmuit=$((height + 15))
# 1 for buy
# 2 for sell
#joltify tx clob place-order jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq    0 0 0 1 1000000 12000000 $heightsubmuit --from validator -y
joltify tx clob place-order jolt1xlgsa2sdvlrwf2w9zlsmvz94vd22g6nlpp8ydg 0 0 0 2 1000000 12000000 $heightsubmuit --from key_2 -y

#sleep 1
#buy
#joltify tx clob place-order jolt1df40uzyqqzehq4n08r3eq9l7anahtalnuemcay 0 0 0 1 10000000 15000000 $heightsubmuit --from key_1 -y

