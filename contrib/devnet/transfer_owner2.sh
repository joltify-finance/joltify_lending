#!/usr/bin/bash


ret=$(joltify q spv  list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')


payfreq=300
delay=5

total_investors=2

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
./run_deposit $indexSenior $number_investors $offset $total_amount false



echo "we need to wait for 5 seconds to have the deposit confirmed"
sleep 5

total_withdrawal_amount_senior=0
for (( c=1; c<=$number_investors; c++ ))
do
  keyindex=$(echo $c+$offset|bc)
ret=$(joltify q spv depositor $indexSenior   $(joltify keys show key_$keyindex -a)   --output json)
withdrawal_senior=$(echo $ret |jq -r '.depositor.withdrawal_amount.amount')
## add the lockled_amount to the total locked amount
total_withdrawal_amount_senior=$(echo $total_withdrawal_amount_senior+$withdrawal_senior|bc)
done
echo $total_withdrawal_amount_senior


# call function prepare investment
while true; do
	# call the funciton of get_class
	while read line; do
		payments_number=$line
		break
	done < <(./sub_commands/get_class.sh)
	payment_count=$(echo "$payments_number" | grep -oE '[0-9]+' | awk '{print $1}')
  # your command goes here
  if [[ $payment_count -gt 2 ]]; then
      echo  "chain is ready for transfer owner"
      break
  fi
  echo  "wait for transfer owner time, current payment round is $payment_count"
  sleep 30 # wait for a second before running the command again
done

echo "we need to wait for 10 seconds to submit the transfer command"
sleep 10


for (( c=1; c<=$total_investors; c++ ))
do
ret=$( joltify tx spv transfer-ownership  $indexSenior   --from key_$c  -y --output json)
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
  echo " transfer owner $1 successful"
else
  echo "transfer owner $1 failed with $ret"
fi
done


