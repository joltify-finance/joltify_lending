#!/bin/bash
ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')


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




while true; do
# repay the  interest of junior pool
ret=$(joltify tx spv submit-withdrawal-proposal $indexJunior  --from key_1 --output json -y)
# check the return code of ret
code=$(echo $ret | jq -r '.code')
if [ $code -eq 0 ]; then
  cecho "GREEN" "submit withdraw request successful"
  break
else
	# get raw log from ret
	rawlog=$(echo $ret| jq -r '.raw_log')
  cecho "RED" "fail with submit the withdraw request with $rawlog"
  sleep 1
fi

done


while true; do
# repay the  interest of junior pool
ret=$(joltify tx spv submit-withdrawal-proposal $indexJunior  --from key_2 --output json -y)
# check the return code of ret
code=$(echo $ret | jq -r '.code')
if [ $code -eq 0 ]; then
  cecho "GREEN" "submit withdraw request successful"
  break
else
	rawlog=$(echo $ret| jq -r '.raw_log')
  cecho "RED" "fail with submit the withdraw request with $rawlog"
  sleep 1
fi

done

