#!/bin/bash

# if $1 and $2 are not empty, then use the $1 and $2
if [ -z "$1" ] || [ -z "$2" ]; then
	ret=$(joltify q spv list-pools --output json)
	# get the index of the pool
	indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
	indexJunior=$(echo $ret | jq -r '.pools_info[1].index')
else
	indexSenior=$1
	indexJunior=$2
fi

# remove the leading 0x in the indexSenior
indexSeniorRemove=$(echo $indexSenior | cut -c3-)
# remove the leading 0x in the indexJunior
indexJuniorRemove=$(echo $indexJunior | cut -c3-)
# get the nft class of the senior pool
id=class-$indexSeniorRemove-0
class1=$(joltify q nft class $id --output json)
#echo $class1

#the jq filter to extract the payment amounts from the class1
paymentsSenior=$(echo $class1 | jq '.class.data.payments[].payment_amount.amount')

paymentsCount=$(echo $(echo $class1 | jq '.class.data.payments|length')-"1" | bc)
echo "payment count:" $paymentsCount

sumPaymentsSenior=$(echo $class1 | jq '[.class.data.payments[].payment_amount.amount | tonumber] | add')
echo "senior payment total:" $sumPaymentsSenior

id=class-$indexJuniorRemove-0
class2=$(joltify q nft class $id --output json)

#the jq filter to extract the payment amounts from the class1
paymentsJunior=$(echo $class2 | jq '.class.data.payments[].payment_amount.amount')
sumPaymentsJunior=$(echo $class2 | jq '[.class.data.payments[].payment_amount.amount | tonumber] | add')
echo "junior payments total:" $sumPaymentsJunior

sum=$(echo $sumPaymentsJunior+$sumPaymentsSenior | bc)
echo "sum of the total payment:" $sum

# I need to have payments in the class1

# for each payment in the payments, add all the payment_amount
#for payment in $payments
#do
#  payment_amount=$(echo $class1| jq -r '.payment_amount.amount')
#  total_payment_amount=$(echo $total_payment_amount+$payment_amount|bc)
#done
