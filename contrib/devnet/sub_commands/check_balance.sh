#!/bin/bash


# http://localhost:1317/cosmos/nft/v1beta1/nfts?owner=jolt1ur5k3vmwt2dw0r0f4zu4lz08c8p6rmwv0f3xkh

ret=$(joltify q spv list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

index=$1
ret=$(joltify q bank balances $(joltify keys show key_$index -a) --output json)
# get the code from json
# check whether the return value of the function is 0
# get the balance
balance=$(echo $ret | jq -r '.balances[0].amount')

ret=$(joltify q spv depositor $indexSenior $(joltify keys show key_$index -a) --output json)
# get the locked amount from ret
locked_amount_senior=$(echo $ret | jq -r '.depositor.locked_amount.amount')
withdrawal_amount_senior=$(echo $ret | jq -r '.depositor.withdrawal_amount.amount')

ret=$(joltify q spv depositor $indexJunior $(joltify keys show key_$index -a) --output json)
# get the locked amount from ret
locked_amount_junior=$(echo $ret | jq -r '.depositor.locked_amount.amount')
withdrawal_amount_junior=$(echo $ret | jq -r '.depositor.withdrawal_amount.amount')

# add all the values together
total_amount=$(echo $balance+$locked_amount_senior+$locked_amount_junior+$withdrawal_amount_senior+$withdrawal_amount_junior | bc)
echo $total_amount
expected=5000000000000000000000000

# check the total_amount is equal to the expected value
#if [[ $total_amount -eq $expected ]]; then
#  echo "Total amount is equal to the expected value"
#else
#  echo "Total amount is not equal to the expected value"
#  exit 1
#fi

# calculate total locked amount
total_locked=$(echo $locked_amount_senior+$locked_amount_junior | bc)
# print the total locked amount
echo "total locked amount is $total_locked"
echo "senior locked is" $locked_amount_senior
echo "senior withdrawal_amount is" $withdrawal_amount_senior
echo "junior withdrawal_amount is" $withdrawal_amount_junior
echo "junior locked is" $locked_amount_junior
echo "balance of $1 is" $balance

echo "$1:" $balance >>balance.txt
