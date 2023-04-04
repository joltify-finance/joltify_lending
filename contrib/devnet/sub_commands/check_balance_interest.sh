#!/usr/bin/bash

while read line; do
	lines+=("$line")
done < <(./check_balance.sh $1)

senior_locked=$(printf '%s\n' "${lines[2]}")
junior_locked=$(printf '%s\n' "${lines[3]}")
balance=$(printf '%s\n' "${lines[6]}")

while read line; do
	lines_interest+=("$line")
done < <(./check_interest.sh $1)

senior_interest=$(printf '%s\n' "${lines_interest[0]}")
junior_interest=$(printf '%s\n' "${lines_interest[1]}")
total_interest=$(printf '%s\n' "${lines_interest[2]}")

echo "balance is $balance"
echo "senior interest" $senior_interest
echo "junior interest" $junior_interest
