#!/usr/bin/bash

output=$(./check_balance.sh $1)
while read line; do
	lines_interest+=("$line")
done < <(./check_interest.sh $1)

senior_interest=$(printf '%s\n' "${lines_interest[0]}")
junior_interest=$(printf '%s\n' "${lines_interest[1]}")
total_interest=$(printf '%s\n' "${lines_interest[2]}")

last_line=$(echo "$output" | tail -n 1)
last_line_num=$(echo "$last_line" | sed 's/[^0-9]*//g')
# Output: 4347645907971020295819735

echo "balance is $last_line_num"
echo "senior interest" $senior_interest
echo "junior interest" $junior_interest
