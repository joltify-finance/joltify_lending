#!/usr/bin/bash


ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')


payfreq=300
delay=5

total_investors=2

ret=$(joltify status)
latestBlockTime=$(echo $ret | jq -r '.sync_info.latest_block_time')
latestBlockTimeSeconds=$(date -d "$latestBlockTime" +%s)

# get the earlist block time
earlistblockTime=$(echo $ret |jq -r '.sync_info.earliest_block_time')

earlistBlockTimeSeconds=$(date -d "$earlistBlockTime" +%s)

# calculate the time difference
timedelta=$(echo $latestBlockTimeSeconds-$earlistBlockTimeSeconds|bc)

timeadj=$(echo $timedelta+$payfreq|bc)

nexttime=$(echo $timeadj/$payfreq*$payfreq+$delay|bc)

expectedtime=$(echo $nexttime+$earlistBlockTimeSeconds|bc)




# for loop in total_investors
total_locked_amount_senior=0
for (( c=1; c<=$total_investors; c++ ))
do
ret=$(joltify q spv depositor $indexSenior   $(joltify keys show key_$c -a)   --output json)
locked_amount_senior=$(echo $ret | jq -r '.depositor.locked_amount.amount')
## add the lockled_amount to the total locked amount
total_locked_amount_senior=$(echo $total_locked_amount_senior+$locked_amount_senior|bc)
done
echo $total_locked_amount_senior

echo "the total amount that locked for the give investors is $total_locked_amount_senior"

echo "how many investors try to invest for transfer?"
read number_investors

echo "total amount want to invest"
read total_amount


offset=$total_investors
./run_deposit $indexSenior $number_investors $offset $total_amount



echo "we need to wait for 5 seconds to have the deposit confirmed"
sleep 5

total_withdrawal_amount_senior=0
for (( c=1; c<=$number_investors; c++ ))
do
  keyindex=$(echo $c+$offset|bc)
  set -x
ret=$(joltify q spv depositor $indexSenior   $(joltify keys show key_$keyindex -a)   --output json)
set +x
withdrawal_senior=$(echo $ret |jq -r '.depositor.withdrawal_amount.amount')
## add the lockled_amount to the total locked amount
total_withdrawal_amount_senior=$(echo $total_withdrawal_amount_senior+$withdrawal_senior|bc)
done
echo $total_withdrawal_amount_senior

# call function prepare investment
while true; do
  # your command goes here
  ret=$(joltify status)
  if [ $? -eq 0 ]; then
    latestBlockTime=$(echo $ret | jq -r '.sync_info.latest_block_time')
    latestBlockTimeSeconds=$(date -d "$latestBlockTime" +%s)
    echo "***********************************"
    echo $latestBlockTimeSeconds
    echo $expectedtime
    echo "***********************************"
    if [[ $latestBlockTimeSeconds -gt $expectedtime ]]; then
      echo  "chain is ready for transfer owner"
      break
    fi
  fi
  echo  "wait for transfer owner time"
  sleep 1 # wait for a second before running the command again
  break
done

#now we call the transfer owner function

for (( c=1; c<=$total_investors; c++ ))
do
set -x
ret=$( joltify tx spv transfer-ownership  $indexSenior   --from key_$c  -y --output json)
set +x
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
  echo " transfer owner $1 successful"
else
  echo "transfer owner $1 failed with $ret"
fi
done


