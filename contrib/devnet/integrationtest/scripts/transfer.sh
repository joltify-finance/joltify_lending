#!/bin/bash

ret=$(joltify tx spv transfer-ownership  $1 --from key_$2 --gas 80000000 --output json -y)
# get the code from json
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
	echo " Deposit $3 successful"
	exit 0
else
	echo "Deposit $3 failed with $ret"
	exit 1
fi