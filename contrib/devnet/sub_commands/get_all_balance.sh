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
