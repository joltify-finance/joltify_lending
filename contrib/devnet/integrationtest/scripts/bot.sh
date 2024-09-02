#!/bin/bash
remote=139.180.193.235
key_name=ki
set -x
height=$(curl -s http://$remote:26657/status | jq -r '.result.sync_info.latest_block_height')
heightsubmuit=$((height + 15))
# 1 for buy
# 2 for sell
#joltify tx clob place-order jolt15qdefkmwswysgg4qxgqpqr35k3m49pkxu8ygkq    0 0 0 1 1000000 12000000 $heightsubmuit --from validator -y

#sleep 1
#buy

random_client_id_buy=$((RANDOM % 100))
random_client_id_sell=$((RANDOM % 100))



# last one is price and it is 10^5
# second last is the quantity and it is 10^10
# current is price 4 size 0.001
#joltify tx clob place-order $(joltify keys show -a $key_name) 0 $random_client_id_sell 0 2 10000000 400000 $heightsubmuit true --from $key_name -y

# run the loop of 100 times
for i in {1..30}
do
  random_client_id=$((RANDOM % 1000000))
  random_price=$((RANDOM % 100))
  # get the number that is random_price * 100000
  random_value_sell=$(((random_price+20) * 100000))
  joltify tx clob place-order $(joltify keys show -a $key_name) 0 $random_client_id 0 1 20000000000 $random_value_sell $heightsubmuit true --from $key_name -y
  sleep 2

#    random_client_id=$((RANDOM % 100))
##    random_price=$((RANDOM % 100))
#    # get the number that is random_price * 100000
#    random_value_sell=$(((random_price -20) * 100000))
#    joltify tx clob place-order $(joltify keys show -a $key_name) 0 $random_client_id 0 1 10000000 $random_value_sell $heightsubmuit true --from $key_name -y
#    sleep 2



done




#joltify tx clob place-order  $(joltify keys show -a key_2) 0 $random_value_buy 0 1 1000000 2000000 $heightsubmuit --from key_2 -y




#joltify tx clob place-order jolt18djzfsnq79f3xpfm78ym5evu7dvpnvxq794w3y 0 123 0 1 1000000 10000000 $heightsubmuit --from key_2 -y
