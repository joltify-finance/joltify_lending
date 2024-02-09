#!/bin/bash

echoerr() { echo "$@" 1>&2; }

ret=$(joltify tx spv claim-interest $1 --from key_$2 --gas 8000000 --output json -y)
# get the code from json
hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
	echo "claim interest $2 successful"
	exit 0
else
	echoerr "claim interest $2 failed with $ret"
	exit 1
fi
