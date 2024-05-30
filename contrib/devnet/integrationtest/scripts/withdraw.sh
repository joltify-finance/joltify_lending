#!/bin/bash
set -x

base=1000000
amount=$(echo $2*$base | bc)
ret=$(joltify tx spv withdraw-principal  $1 $amount"ausdc" --from key_$3 --gas 80000 --output json -y)
# get the code from json
hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo " withdrawal $3 successful"
	exit 0
else
	echo "withdrawal $3 failed with $ret"
	exit 1
fi
