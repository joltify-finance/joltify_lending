#!/bin/bash
all_keys=5

total_amount=0
for (( c=1; c<=$all_keys; c++ ))
do
  ret=$(joltify q bank balances  $(joltify keys show key_$c -a)  --output json)
  balance=$(echo $ret | jq -r '.balances[0].amount')
  total_amount=$(echo $total_amount+$balance | bc)
  echo $ret
done
echo $total_amount

# now we check the payment ?= recevied
ret=$(joltify q spv total-reserve --output json)
coins=$(echo $ret | jq -r '.coins')
amount_reserved=${coins%ausdc}
ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
balance_validator=$(echo $ret | jq -r '.balances[1].amount')

sum_all=$(echo $total_amount-25000000000000000000000000+$amount_reserved+$balance_validator| bc)


echo "recovered total amount of validator is  $sum_all"
