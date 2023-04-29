#!/bin/bash

base=1000000000000000000
amount=$(echo $2*$base | bc)
set -x
ret=$(joltify tx spv deposit $1 $amount"ausdc" --from key_$3 --output json -y)
set +x
# get the code from json
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
	echo " Deposit $1 successful"
else
	echo "Deposit $1 failed with $ret"
	exit 1
fi
