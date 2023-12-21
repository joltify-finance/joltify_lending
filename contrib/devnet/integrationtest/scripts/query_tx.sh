#!/bin/bash
max_attempts=8
sleep 5
for ((i=1; i<=$max_attempts; i++)); do
    echo "Attempt $i:"
	ret=$(joltify q tx --type hash $1 --output json)
 	code=$(echo $ret | jq -r '.code')
    # Check the return code
    if [ $code -eq 0 ]; then
        echo "Command succeeded. Exiting."
        exit 0
    else
        echo "Command failed. Retrying... $ret"
        sleep 3
    fi
done
echo "Maximum attempts reached. Exiting."
exit 1
