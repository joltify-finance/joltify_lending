#!/usr/bin/env bash
# Path: contrib/devnet/check_value.sh
total_investors=2

ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

function queryAllBalances(){
  for (( c=1; c<=$total_investors; c++ ))
  do
    ret=$(joltify q bank balances $(joltify keys show key_$c -a) --output json)
    # get the code from json
    # check whether the return value of the function is 0
      # get the balance
      balance=$(echo $ret | jq -r '.balances[0].amount')
      # store the balance in balances array
      balances[$c]=$balance
  done
}

function queryAllLockedAmount(){

  for (( c=1; c<=$total_investors; c++ ))
  do
    ret=$(joltify q spv depositor $indexSenior   $(joltify keys show key_$c -a)   --output json)
    # get the locked amount from ret
    locked_amount_senior=$(echo $ret | jq -r '.depositor.locked_amount.amount')
    # store the locked amount in locked_amounts array
    locked_amounts_senior[$c]=$locked_amount_senior


    ret=$(joltify q spv depositor $indexJunior $(joltify keys show key_$c -a)   --output json)
    # get the locked amount from ret
    locked_amount_junior=$(echo $ret | jq -r '.depositor.locked_amount.amount')
    # store the locked amount in locked_amounts array
    locked_amounts_junior[$c]=$locked_amount_junior
  done
}

# we check that the total amount is equal to the sum of all the locked amounts
function checkTotalAmountJunior(){
  # get the total amount from pool

  ret=$(joltify q spv query-pool $indexJunior --output json)
  # get the total locked from the pool
 total_locked_junior=$(echo $ret | jq -r '.pool_info.borrowed_amount.amount')


  ret=$(joltify q spv query-pool $indexSenior --output json)
  # get the total locked from the pool
 total_locked_senior=$(echo $ret | jq -r '.pool_info.borrowed_amount.amount')

  total_amount=$(echo $total_locked_junior+$total_locked_senior|bc)

  sum=0
  for i in "${locked_amounts_junior[@]}"
  do
    sum=$(echo $sum+$i|bc)
  done
  echo $total_locked_junior
  echo $sum
  # check whether the total amount is equal to the sum of all the locked amounts
  if [ $total_locked_junior == $sum ]; then
    echo "Total amount is equal to the sum of all the locked amounts"
  else
    echo "Total amount is not equal to the sum of all the locked amounts"
    exit 1
  fi
}

# we check that the total amount is equal to the sum of all the locked amounts
function checkTotalAmountSenior(){
  # get the total amount from pool

  ret=$(joltify q spv query-pool $indexSenior --output json)
  # get the total locked from the pool
 total_locked_senior=$(echo $ret | jq -r '.pool_info.borrowed_amount.amount')


  sum=0
  for i in "${locked_amounts_senior[@]}"
  do
    sum=$(echo $sum+$i|bc)
  done
  echo $total_locked_senior
  echo $sum
  # check whether the total amount is equal to the sum of all the locked amounts
  if [ $total_locked_senior == $sum ]; then
    echo "Total amount is equal to the sum of all the locked amounts"
  else
    echo "Total amount is not equal to the sum of all the locked amounts"
    exit 1
  fi
}



function queryAllInterest(){
  pool_index=$1
  for (( c=1; c<=$total_investors; c++ ))
  do
    ret=$(joltify q spv claimable-interest  $(joltify keys show key_$c -a) $pool_index  --output json)
    interest=$(echo $ret | jq -r '.claimable_interest_amount.amount')
    interests[$c]=$interest
    # get the locked amount from ret
  done
}

function queryAllDepositInfo(){
  pool_index=$1
  for (( c=1; c<=$total_investors; c++ ))
  do
    ret=$(joltify q spv depositor $pool_index $(joltify keys show key_$c -a)  --output json)
    echo $ret

    locked_amount=$(echo $ret | jq -r '.depositor.locked_amount.amount')
    withdrawal_amount=$(echo $ret | jq -r '.depositor.withdrawal_amount.amount')
    locked_amounts[$c]=$locked_amount
    withdrawal_amounts[$c]=$withdrawal_amount
  done
}


# get the locked_amount


queryAllBalances
queryAllLockedAmount
checkTotalAmountJunior
checkTotalAmountSenior


