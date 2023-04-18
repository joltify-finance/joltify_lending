#!/bin/bash


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

junior=$(echo 200001*$base|bc)
senior=$(echo 800001*$base|bc)
interest=$(echo 100*$base|bc)

echo "do you need to pay principal? (y/n)"
read confirm
set -x
# check whether confirm is y
if [ $confirm == "y" ]; then
  # repay the  principal of junior pool
  ret=$(joltify tx spv pay-principal $indexJunior $junior"ausdc" --from validator --output json -y)
  # check the return code of ret
  code=$(echo $ret | jq -r '.code')
  if [ $code -eq 0 ]; then
	cecho "GREEN" "pay principal junior successful"
  else
	cecho "RED" "pay principal junior failed with $ret"
	exit 1
  fi

  # repay the  principal of senior pool
  ret=$(joltify tx spv pay-principal $indexSenior $senior"ausdc" --from validator --output json -y)
  # check the return code of ret
  code=$(echo $ret | jq -r '.code')
  if [ $code -eq 0 ]; then
	cecho "GREEN" "Pay principal senior successful"
  else
	cecho "RED" "Pay principal senior failed with $ret"
	exit 1
  fi
fi

