#!/bin/bash
all_keys=$1

total_amount=0
for (( c=1; c<=$all_keys; c++ ))
do
  ret=$(joltify q bank balances  $(joltify keys show key_$c -a)  --output json)
  balance=$(echo $ret | jq -r '.balances[0].amount')
  interest=$(echo $balance - 5000000000000 | bc)
  total_amount=$(echo $total_amount+ $interest | bc)
  echo ">>$c>$balance"
done


validator_ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
echo $validator_ret

# now we check the payment ?= recevied
ret=$(joltify q spv total-reserve --output json)
coins=$(echo $ret | jq -r '.coins')
amount_reserved=${coins%ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3}
ret=$(joltify q bank balances  $(joltify keys show validator -a)  --output json)
balance_validator=$(echo $ret | jq -r '.balances[1].amount')

sum_all=$(echo $total_amount+$amount_reserved+$balance_validator| bc)

echo "###########################################################"
echo "recovered total amount of validator is  $sum_all"
echo "total interest to investors $total_amount"
echo "reserved $amount_reserved"
echo "balance validator $balance_validator"
echo "###########################################################"
