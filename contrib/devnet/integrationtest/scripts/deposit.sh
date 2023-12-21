#!/bin/bash

base=1000000000000000000
amount=$(echo $2*$base | bc)
ret=$(joltify tx spv deposit $1 $amount"ausdc" --from key_$3 --output json -y)
# get the code from json
code=$(echo $ret | jq -r '.txhash')
./scripts/query_tx.sh $code
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo " Deposit $3 successful"
else
	echo "Deposit $3 failed with $ret"
	exit 1
fi
