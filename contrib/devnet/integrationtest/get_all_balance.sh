#!/bin/bash
all_keys=$1


DATA=/Users/yb/.tmpdisk/ram
HOME=$DATA/joltifydata
total_amount=0
for (( c=1; c<=$all_keys; c++ ))
do
  ret=$(joltify q bank balances  $(joltify keys show key_$c -a)  --output json)
  balance=$(echo $ret | jq -r '.balances[0].amount')
  interest=$(echo $balance - 5000000000000000000000000 | bc)
  total_amount=$(echo $total_amount+ $interest | bc)
  echo ">>$c>$balance"
done


validator_ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
echo $validator_ret

# now we check the payment ?= recevied
ret=$(joltify q spv total-reserve --output json)
coins=$(echo $ret | jq -r '.coins')
amount_reserved=${coins%ausdc}
ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
balance_validator=$(echo $ret | jq -r '.balances[1].amount')

sum_all=$(echo $total_amount+$amount_reserved+$balance_validator| bc)

echo "###########################################################"
echo "recovered total amount of validator is  $sum_all"
echo "total interest to investors $total_amount"
echo "reserved $amount_reserved"
echo "balance validator $balance_validator"
echo "###########################################################"
