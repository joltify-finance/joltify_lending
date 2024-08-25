#!/bin/bash
set -x
height=$(curl -s http://localhost:26657/status | jq -r '.result.sync_info.latest_block_height')
heightsubmuit=$((height + 15))
# 1 for buy
# 2 for sell
#joltify tx clob place-order jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq    0 0 0 1 1000000 12000000 $heightsubmuit --from validator -y

#sleep 1
#buy
joltify tx clob place-order $(joltify keys show -a key_1) 0 101 0 2 1000000 2000000 $heightsubmuit --from key_1 -y

joltify tx clob place-order  $(joltify keys show -a key_2) 0 242 0 1 1000000 2000000 $heightsubmuit --from key_2 -y




#joltify tx clob place-order jolt18djzfsnq79f3xpfm78ym5evu7dvpnvxq794w3y 0 123 0 1 1000000 10000000 $heightsubmuit --from key_2 -y
