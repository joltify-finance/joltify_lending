#!/bin/bash

ret=$(joltify tx spv submit-withdrawal-proposal  $1 --from key_$2 --gas 20000000 --output json -y)
# get the code from json
hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo " Partial Withdraw submit $3 successful"
	exit 0
else
	echo "Partial withdraw submit $3 failed with $ret"
	exit 1
fi
