#!/bin/bash
# show the balance of usdc
balance=$(joltify q bank balances $(joltify keys show -a $1) -o json | jq -r '.balances[1].amount')
echo $balance
