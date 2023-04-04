#!/bin/bash

cecho() {
	RED="\033[0;31m"
	GREEN="\033[0;32m"  # <-- [0 means not bold
	YELLOW="\033[1;33m" # <-- [1 means bold
	CYAN="\033[1;36m"
	# ... Add more colors if you like

	NC="\033[0m" # No Color

	# printf "${(P)1}${2} ${NC}\n" # <-- zsh
	printf "${!1}${2} ${NC}\n" # <-- bash
}

ret=$(joltify q spv list-pools --output json)
# get the index of the pool
indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

./check_balance_interest.sh $1

while read line; do
	lines_balance_interest_before+=("$line")
done < <(./check_balance_interest.sh $1)

last_balance=$(echo ${lines_balance_interest_before[0]} | awk '{print $NF}')
last_senior_interest=$(echo ${lines_balance_interest_before[1]} | awk '{print $NF}')
last_junior_interest=$(echo ${lines_balance_interest_before[2]} | awk '{print $NF}')

# total is last_balance+last_senior_interest+last_junior_interest
total=$(echo $last_balance+$last_senior_interest+$last_junior_interest | bc)

ret=$(joltify tx spv claim-interest $indexSenior --output json --from key_$1 -y --gas 800000)

# get the code from ret
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
	cecho "GREEN" "Senior pool claim interest successful"
else
	cecho "RED" "Senior pool claim interest failed with $ret"
	exit 1
fi

ret=$(joltify tx spv claim-interest $indexJunior --output json --from key_$1 -y --gas 800000)
# get the code from ret
code=$(echo $ret | jq -r '.code')
# check whether the return value of the function is 0
if [ $code -eq 0 ]; then
	cecho "GREEN" "Junior pool claim interest successful"
else
	cecho "RED" "Junior pool claim interest failed with $ret"
	exit 1
fi

echo "-------------------------------after----------------------------------"
./check_balance_interest.sh $1

while read line; do
	lines_balance_interest+=("$line")
done < <(./check_balance_interest.sh $1)

last_balance=$(echo ${lines_balance_interest[0]} | awk '{print $NF}')
last_senior_interest=$(echo ${lines_balance_interest[1]} | awk '{print $NF}')
last_junior_interest=$(echo ${lines_balance_interest[2]} | awk '{print $NF}')
echo "last balance is $last_balance"
echo "last senior interest" $last_senior_interest
echo "last junior interest" $last_junior_interest

# check whether total is equal to last_balance
if [[ $total -eq $last_balance ]]; then
	cecho "GREEN" "Total is equal to last balance"
else
	cecho "RED" "Total is not equal to last balance"
	exit 1
fi
