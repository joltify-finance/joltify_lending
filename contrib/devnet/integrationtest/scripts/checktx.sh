#!/bin/bash

# Run the loop three times
for ((i=1; i<=5; i++)); do
    echo "This is loop iteration number $i for checking tx"
    ret=$(joltify q tx --type hash $1  -o json)
    # get the code from json
    code=$(echo $ret | jq -r '.code')


    if [ $? -eq 0 ]; then
    # Add your commands or script logic here
    # Check a condition to determine if the loop should break
    if [ $code -eq 0 ]; then
        echo " Transaction $1 successful"
        exit 0
    fi
    fi
    sleep 3
done
exit 1
