#!/usr/bin/bash

ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

# check the claimable interest of the senior pool
ret=$(joltify q spv claimable-interest  $(joltify keys show key_$1 -a) $indexSenior  --output json)
# get the claimable interest
claimable_interest_senior=$(echo $ret | jq -r '.claimable_interest_amount.amount')
# check the claimable interest of the junior pool
ret=$(joltify q spv claimable-interest  $(joltify keys show key_$1 -a) $indexJunior  --output json)
# get the claimable interest
claimable_interest_junior=$(echo $ret | jq -r '.claimable_interest_amount.amount')



# print the claimable interest
echo "claimable interest of the senior pool is $claimable_interest_senior"
echo "claimable interest of the junior pool is $claimable_interest_junior"

# calculate the total claimable interest
total_claimable_interest=$(echo $claimable_interest_senior+$claimable_interest_junior|bc)
echo "total claimable interest is $total_claimable_interest"

