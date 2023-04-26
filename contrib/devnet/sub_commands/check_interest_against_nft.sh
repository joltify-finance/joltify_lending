#!/bin/bash
# for this check we sum all the investors claimable interest and  compare with the nft paidout interest
# the first parameter is the pool index

i=1
while [ $i -ne $(echo $1 + 1|bc) ]
do


# for the while loop from 1 to 10
# get the interest

s_total=0
j_total=0


while read line; do
	lines+=("$line")
done < <(./check_interest.sh $i)
# get the claimable interest
echo $lines
s_interest=$(printf '%s\n' "${lines[0]}")

j_interest=$(printf '%s\n' "${lines[1]}")
sd_number=$(echo "$s_interest" | sed 's/[^0-9]*//g')
jd_number=$(echo "$s_interest" | sed 's/[^0-9]*//g')

unset lines
#echo "#####################"
#echo $sd_number
#echo $jd_number
#echo "#####################"
#
#
## add all the jd_number and sd_number
#s_total=$(( s_total+$sd_number))
#j_total=$(( j_total+$jd_number ))
#
i=$(($i+1))

# done
done

echo total senior inter is $s_total
echo total junior inter is $j_total



