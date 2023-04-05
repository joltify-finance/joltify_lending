#!/bin/bash

ret=$(joltify tx spv withdraw-principal $1 $2 --from key_$3 --output json -y --gas 100000000)
# get the code from json
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
	echo " withdraw $1 successful"
else
	echo "withdraw $1 failed with $ret"
	exit 1
fi
