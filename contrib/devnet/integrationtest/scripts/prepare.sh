#!/bin/bash

DATA=/Users/yb/.tmpdisk/ram
HOME=$DATA/joltifydata
bash -x
# amount is value*base
base=1000000000000000000
junior=$(echo 200000*$base|bc)
senior=$(echo 800000*$base|bc)
all_keys=100

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




for (( c=1; c<=$all_keys; c++ ))
do
  ret=$(joltify keys show key_$c --keyring-backend test --output json)
  # get the address from the json
  address=$(echo $ret | jq -r '.address')
  # put the address in the array
  allInvestors="$allInvestors,$address"
done

# remove the first comma of all investors
allInvestors=${allInvestors:1}

ret=$(joltify tx spv create-pool $2 $1 0.15 0.0875 $junior"ausdc" $senior"ausdc"  --from validator --gas 8000000 --output json -y)
# get the code from the json
tx_hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $tx_hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
cecho "GREEN" "Pool creation successful"
else
cecho "READ" "Pool creation failed with $tx_hash"
exit 1
fi
return


ret=$(joltify tx kyc upload-investor 44 $allInvestors --from validator  --gas 8000000 --output json -y)
tx_hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $tx_hash
# get the code from json
#get the code value from json
# remove leading empty spaces
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
cecho "GREEN" "KYC uplad investor successful"
else
cecho "READ" "KYC upload investor failed $ret"
fi

# get the index of the pool
indexJunior=$3

#ret=$(joltify tx spv active-pool $indexSenior --from validator --output json -y)
#code=$(echo $ret | jq -r '.code')
## check whether the return value of the function is 0
#if [ $code -eq 0 ]; then
#cecho  "GREEN" "Junior pool activation successful"
#else
#cecho "READ" "Senior pool activation failed $ret"
#exit 1
#fi

ret=$(joltify tx spv active-pool $indexJunior --from validator --gas 8000000 --output json -y)
tx_hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $tx_hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
cecho "GREEN" "junior pool activation successful"
else
cecho "READ" "junior pool activation failed with $ret"
exit 1
fi

ret=$(joltify tx spv add-investors  $indexJunior 44  --from validator --gas 8000000 --output json -y)

tx_hash=$(echo $ret | jq -r '.txhash')
./scripts/checktx.sh $tx_hash
# check whether the return value of the function is 0
if [ $? -eq 0 ]; then
cecho "GREEN" "Add investors to junior pool successful"
else
cecho "READ" "Add investors  to junior pool failed with $ret"
exit 1
fi

#ret=$(joltify tx spv add-investors  $indexSenior 44  --from validator --gas 20000000 --output json -y)
#code=$(echo $ret | jq -r '.code')
## check whether the return value of the function is 0
#if [ $code -eq 0 ]; then
#cecho "GREEN" "Add investors to senior pool successful"
#else
#cecho "READ" "Add investors to junior pool failed with $ret"
#exit 1
#fi
#
