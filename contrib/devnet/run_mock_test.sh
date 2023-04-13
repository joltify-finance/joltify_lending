#!/bin/bash
# generate n cosmos accounts
# loop function to generate n cosmos accounts
base=1000000000000000000
total_investors=2
all_keys=5
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

function prepareInvestment() {
  # amount is value*base
  junior=$(echo 200000*$base|bc)
  senior=$(echo 800000*$base|bc)
 ret=$(joltify tx spv create-pool "testpool" 1 0.15 $junior"ausdc"  --from validator -y --output json)

 # get the code from the json
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "Pool creation successful"
  else
    cecho "READ" "Pool creation failed"
    #exit 1
  fi


 ret=$(joltify tx kyc upload-investor 44 $allInvestors --from validator -y --output json)
 # get the code from json
 #get the code value from json
 code=$(echo $ret | jq -r '.code')
 # remove leading empty spaces
 # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "KYC upload investor successful"
  else
    cecho "READ" "KYC upload investor failed"
    exit 1
  fi


  ret=$(joltify q spv  list-pools --output json)
  # get the index of the pool
  indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
  indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

  ret=$(joltify tx spv active-pool $indexSenior --from validator --output json -y)
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho  "GREEN" "Junior pool activation successful"
  else
    cecho "READ" "Senior pool activation failed $ret"
    exit 1
  fi


  ret=$(joltify tx spv active-pool $indexJunior --from validator --output json -y)
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "Senior pool activation successful"
  else
    cecho "READ" "Senior pool activation failed with $ret"
    exit 1
  fi

  ret=$(joltify tx spv add-investors  $indexJunior 44  --from validator --output json -y)
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "Add investors to junior pool successful"
  else
    cecho "READ" "Add investors  to junior pool failed with $ret"
    exit 1
  fi

  ret=$(joltify tx spv add-investors  $indexSenior 44  --from validator --output json -y)
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "Add investors to senior pool successful"
  else
    cecho "READ" "Add investors to junior pool failed with $ret"
    exit 1
  fi

}

function deposit() {
  ret=$(joltify q spv  list-pools --output json)
  # get the index of the pool
  indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
  indexJunior=$(echo $ret | jq -r '.pools_info[1].index')
                # poolindex total_investors offset amount whether withdraw
  ./run_deposit $indexJunior 2 0 300000 false
  ./run_deposit $indexSenior 2 0 1000000 false
}

#function deposit() {
#  ret=$(joltify q spv  list-pools --output json)
#  # get the index of the pool
#  indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
#  indexJunior=$(echo $ret | jq -r '.pools_info[1].index')
#
#  for (( c=1; c<=$total_investors; c++ ))
#  do
#    # get the random value from 2000 to 8000  and multiply it with base
#    amount_deposit_junior=$(echo $(( ( RANDOM % 6000 )  + 2000 ))*$base|bc)
#    all_amounts_deposits_junior="$all_amounts_deposits_junior,$amount_deposit_joltify"
#    ret=$(joltify tx spv deposit $indexSenior $amount_deposit_junior"ausdc" --from key_$c --output json -y)
#    # get the code from json
#    code=$(echo $ret | jq -r '.code')
#    # check whether the return value of the function is 0
#    if [ $code -eq 0 ]; then
#      cecho "GREEN" "Deposit junior successful"
#    else
#      cecho "READ" "Deposit junior failed with $ret"
#      exit 1
#    fi
#
#    # get the random value from 2000 to 8000  and multiply it with base
#    amount_deposit_senior=$(echo $(( ( RANDOM % 6000 )  + 2000 ))*$base|bc)
#    all_amounts_deposits_senior="$all_amounts_deposits_senior,$amount_deposit_senior"
#    ret=$(joltify tx spv deposit $indexSenior $amount_deposit_senior"ausdc" --from key_$c --output json -y)
#    # get the code from json
#    code=$(echo $ret | jq -r '.code')
#    # check whether the return value of the function is 0
#    if [ $code -eq 0 ]; then
#      cecho "GREEN" "Deposit senior successful"
#    else
#      cecho "READ" "Deposit senior failed with $ret"
#      exit 1
#    fi
#
#  done
#}

function borrow() {

  ret=$(joltify q spv  list-pools --output json)
  # get the index of the pool
  indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
  indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

  amount=$(echo 800000*$base|bc)
   ret=$(joltify tx spv  borrow $indexSenior $amount"ausdc" --from validator -y --output json --gas 800000)
  # get the code from json
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "Borrow senior successful"
  else
    cecho "READ" "Borrow senior failed with $ret"
    exit 1
  fi

  amount=$(echo 200000*$base|bc)
  # run the borrow for junior
  ret=$(joltify tx spv  borrow $indexJunior $amount"ausdc" --from validator -y --output json --gas 800000)
  # get the code from json
  code=$(echo $ret | jq -r '.code')
  # check whether the return value of the function is 0
  if [ $code -eq 0 ]; then
    cecho "GREEN" "Borrow junior successful"
  else
    cecho "READ" "Borrow junior failed with $ret"
    exit 1
  fi
}


# call function prepare investment
while true; do
  # your command goes here
  ret=$(joltify status)
  if [ $? -eq 0 ]; then
#    blockHeight=$($ret | jq -r '.sync_info.latest_block_height')
  height=$(echo $ret | jq -r '.sync_info.latest_block_height')
    if [ $height -gt 3 ]; then
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


prepareInvestment
deposit
borrow
