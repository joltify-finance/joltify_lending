#!/bin/bash

ret=$(joltify tx spv transfer-ownership  $1 --from key_$2 --gas 80000 --output json -y)
# get the code from json
hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo " transfer $3 successful"
	exit 0
else
	echo "transfer $3 failed with $ret"
	exit 1
fi
