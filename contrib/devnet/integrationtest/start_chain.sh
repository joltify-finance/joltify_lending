#!/bin/bash
# generate n cosmos accounts
# loop function to generate n cosmos accounts
base=1000000000000000000
all_keys=100
project_index=1
cecho(){
    RED="\033[0;31m"
    GREEN="\033[0;32m"  # <-- [0 means not bold
    YELLOW="\033[1;33m" # <-- [1 means bold
    CYAN="\033[1;36m"
    # ... Add more colors if you like

    NC="\033[0m" # No Color

    # printf "${(P)1}${2} ${NC}\n" # <-- zsh
    printf "${!1}${2} ${NC}\n" # <-- bash
}

rm -rf ~/.joltify
bash init-new-chain-run.sh


# call function prepare investment
while true; do
  # your command goes here
  ret=$(joltify status)
  if [ $? -eq 0 ]; then
#    blockHeight=$($ret | jq -r '.sync_info.latest_block_height')
  height=$(echo $ret | jq -r '.sync_info.latest_block_height')
    if [ $height -gt 1 ]; then
      cecho "CYAN" "chain is ready"
      break
    fi
  fi
  cecho "YELLOW" "wait for chain to be ready"
  sleep 1 # wait for a second before running the command again
done

ret=$(joltify tx pricefeed postprice aud:usd 0.7 253402300799 -y --from validator --output json)
 code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "submit price done"
  else
    cecho "READ" "submit price failed with $ret"
    exit 1
  fi
