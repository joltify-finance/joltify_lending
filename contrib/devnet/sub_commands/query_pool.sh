#!/usr/bin/bash

ret=$(joltify q spv list-pools --output json)
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

pooltype=$(echo $ret | jq -r '.pools_info[0].pool_type')
echo "#############$pooltype#############"
borrowed_local=$(echo $ret | jq -r '.pools_info[0].borrowed_amount'|jq -r '.amount')
echo borrowed amount: $(echo $borrowed_local *0.7 | bc)
echo usable_amount: $(echo $ret | jq -r '.pools_info[0].usable_amount')
echo apy: $(echo $ret | jq -r '.pools_info[0].apy')
echo pay freq :  $(echo $ret | jq -r '.pools_info[0].pay_freq')
echo "####################################"


a="{ "denom": "aud-ausdc", "amount": "285714285714285714285714" }"
# get the amoubnt from a
amount=$(echo $a | jq -r '.amount')



pooltype=$(echo $ret | jq -r '.pools_info[1].pool_type')
echo "#############$pooltype#############"
borrowed_local=$(echo $ret | jq -r '.pools_info[1].borrowed_amount'|jq -r '.amount')
echo borrowed amount: $(echo $borrowed_local *0.7 | bc)
echo usable_amount: $(echo $ret | jq -r '.pools_info[1].usable_amount')
echo apy: $(echo $ret | jq -r '.pools_info[1].apy')
echo pay freq :  $(echo $ret | jq -r '.pools_info[1].pay_freq')
echo "####################################"


