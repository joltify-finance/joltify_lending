#!/usr/bin/bash


ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')


payfreq=300
delay=5

total_transfers=2

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




for (( c=1; c<=$total_transfers; c++ ))
do
ret=$(joltify q spv depositor $indexSenior   $(joltify keys show key_$c -a)   --output json)
# get the locked amount from ret
locked_amount_senior=$(echo $ret | jq -r '.depositor.locked_amount.amount')

# add the lockled_amount to the total locked amount
total_locked_amount_senior=$(echo $total_locked_amount_senior+$locked_amount_senior|bc)
done
echo $total_locked_amount_senior




