#!/bin/bash

set -x
ret=$(joltify tx spv withdraw-principal $1 $2 --from key_$3 --output json -y --gas 80000)
set +x
# get the code from json
hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo "withdraw principal $1 successful"
else
	echo "withdraw principal $1 failed with $ret"
	exit 1
fi
