#!/bin/bash


base=1000000000000000000
amount=$(echo $2*$base | bc)
ret=$(joltify tx spv repay-interest $1 $amount"ausdc" --from validator --gas 20000000 --output json -y)
# get the code from json
hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo " repay interest $3 successful"
	exit 0
else
	echo "repay interest $3 failed with $ret"
	exit 1
fi
