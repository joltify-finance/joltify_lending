#!/bin/bash
all_keys=$1

total_amount=0
for (( c=1; c<=$all_keys; c++ ))
do
  ret=$(joltify q bank balances  $(joltify keys show key_$c -a)  --output json)
  balance=$(echo $ret | jq -r '.balances[0].amount')
  total_amount=$(echo $total_amount+$balance | bc)
  echo $ret
done

echo $total_amount

validator_ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
echo $validator_ret

all_balance=$(echo 5000000000000000000000000*all_keys | bc)
echo $all_balance
# now we check the payment ?= recevied
ret=$(joltify q spv total-reserve --output json)
coins=$(echo $ret | jq -r '.coins')
amount_reserved=${coins%ausdc}
ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
balance_validator=$(echo $ret | jq -r '.balances[1].amount')

sum_all=$(echo $total_amount-$all_balance+$amount_reserved+$balance_validator| bc)


echo "recovered total amount of validator is  $sum_all"
