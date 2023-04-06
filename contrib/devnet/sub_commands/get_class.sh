#!/bin/bash

# if $1 and $2 are not empty, then use the $1 and $2
if [ -z "$2" ] || [ -z "$3" ]; then
	ret=$(joltify q spv list-pools --output json)
	# get the index of the pool
	indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
	indexJunior=$(echo $ret | jq -r '.pools_info[1].index')
else
	indexSenior=$2
	indexJunior=$3
fi

# remove the leading 0x in the indexSenior
indexSeniorRemove=$(echo $indexSenior | cut -c3-)
# remove the leading 0x in the indexJunior
indexJuniorRemove=$(echo $indexJunior | cut -c3-)
# get the nft class of the senior pool
id=class-$indexSeniorRemove-$1
class1=$(joltify q nft class $id --output json)


paymentsCount=$(echo $(echo $class1 | jq '.class.data.payments|length')-"1" | bc)
echo "payment count:" $paymentsCount

#the jq filter to extract the payment amounts from the class1
data=$(echo $class1 | jq '.class.data.payments[].payment_amount.amount')

output_data=$(echo $data | tr -d '"')

# Split the string into an array
readarray -t array <<< "$output_data"

sumPaymentsSenior=0
for payment in  $array
do
  sumPaymentsSenior=$(echo $sumPaymentsSenior+$payment |bc)
done
echo "senior payment total:" $sumPaymentsSenior

id=class-$indexJuniorRemove-$1
class2=$(joltify q nft class $id --output json)

#the jq filter to extract the payment amounts from the class1

#the jq filter to extract the payment amounts from the class1
data=$(echo $class2 | jq '.class.data.payments[].payment_amount.amount')

output_data=$(echo $data | tr -d '"')

# Split the string into an array
readarray -t array <<< "$output_data"

sumPaymentsJunior=0
for payment in  $array
do
  sumPaymentsJunior=$(echo $sumPaymentsJunior+$payment |bc)
done
echo "junior payment total:" $sumPaymentsJunior





sum=$(echo $sumPaymentsJunior+$sumPaymentsSenior | bc)
echo "sum of the total payment:" $sum

# I need to have payments in the class1

# for each payment in the payments, add all the payment_amount
#for payment in $payments
#do
#  payment_amount=$(echo $class1| jq -r '.payment_amount.amount')
#  total_payment_amount=$(echo $total_payment_amount+$payment_amount|bc)
#done
