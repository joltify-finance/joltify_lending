#!/bin/bash


base=1000000
amount=$(echo $2*$base | bc)
ret=$(joltify tx spv repay-interest $1 $amount"ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3" --from validator --gas 20000000 --output json -y)
# get the code from json
code=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $code
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo " pay interest $2 successful"
	exit 0
else
	echo "pay interest $2 failed with $ret"
	exit 1
fi
