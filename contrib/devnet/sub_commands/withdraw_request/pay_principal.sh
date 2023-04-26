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

ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

base=1000000000000000000

junior=$(echo 200000*$base|bc)
senior=$(echo 800000*$base|bc)
interest=$(echo 8000*$base|bc)


while true; do

  # repay the  principal of junior pool
  ret=$(joltify tx spv  pay-principal-partial $indexJunior $junior"ausdc" --from validator --output json -y)
  # check the return code of ret
  code=$(echo $ret | jq -r '.code')
  if [ $code -eq 0 ]; then
	cecho "GREEN" "pay partial principal junior successful"
	break
  else
  	rawlog=$(echo $ret| jq -r '.raw_log')
	cecho "RED" "pay partial principal junior failed with $rawlog"
  fi

done

exit 0
while true; do
  # repay the  principal of senior pool
  ret=$(joltify tx spv  pay-principal-partial $indexSenior $senior"ausdc" --from validator --output json -y)
  # check the return code of ret
  code=$(echo $ret | jq -r '.code')
  if [ $code -eq 0 ]; then
	cecho "GREEN" "Pay partial principal senior successful"
	break
  else
  	rawlog=$(echo $ret| jq -r '.raw_log')
	cecho "RED" "Pay partial principal senior failed with $rawlog"
  fi
done
