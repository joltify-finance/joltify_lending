#!/bin/bash

indexJunior="0x43ce7e072884180e125328e727911ad83fcaba1cc487ece1ccc3e19376f51118"




cecho(){
    RED="\033[0;31m"
    GREEN="\033[0;32m"  # <-- [0 means not bold
    YELLOW="\033[1;33m" # <-- [1 means bold
    CYAN="\033[1;36m"
    # ... Add more colors if you like

    NC="\033[0m" # No Color

    # printf "${(P)1}${2} ${NC}\n" # <-- zsh
    printf "${!1}${2} ${NC}\n" # <-- bash
}


# call function prepare investment
while true; do
	ret=$(./window.sh |tail -n 1)
	echo $ret
	# split the ret with comma
	IFS=',' read -ra ADDR <<< "$ret"
	# get the second element
	second=${ADDR[1]}
	third=${ADDR[2]}
	echo "need to wait for $second to submit withdraw request"
	#if second is less than 10
	#if [ $second -lt 300 ]; then
		#submit the withdrawal request
#		while true; do
		# repay the  interest of junior pool
		ret=$(joltify tx spv submit-withdrawal-proposal $indexJunior  --from key_$1 --output json -y)
		# check the return code of ret
		code=$(echo $ret | jq -r '.code')
		if [ $code -eq 0 ]; then
  		cecho "GREEN" "submit withdraw request successful"
	done=1
  		break
		else
			# get raw log from ret
			rawlog=$(echo $ret| jq -r '.raw_log')
  		cecho "RED" "fail with submit the withdraw request with $rawlog"
  		sleep 1
		fi
#		done
	#fi


	if [ $done -eq 1 ]; then
		break
	fi

	sleep 5
done



base=1000000000000000000

junior=$(echo 200000*$base|bc)
senior=$(echo 800000*$base|bc)
interest=$(echo 8000*$base|bc)



# call function prepare investment
while true; do

	ret=$(./window.sh |tail -n 1)
	echo $ret

	# split the ret with comma
	IFS=',' read -ra ADDR <<< "$ret"
	# get the second element
	third=${ADDR[2]}

	echo "need to wait for $third to submit pay principal request"
	#if second is less than 10
	#if [ $third -lt 300 ]; then
		# repay the  interest of junior pool

#  			while true; do
  		ret=$(joltify tx spv  pay-principal-partial $indexJunior $junior"ausdc" --from validator --output json -y)
		# check the return code of ret
		code=$(echo $ret | jq -r '.code')
		if [ $code -eq 0 ]; then
  			cecho "GREEN" "pay principal junior successful"
  			break
		else
  			cecho "RED" "pay principal junior failed with $ret"
		fi

#	done
	#fi

	sleep 5
done

echo "done@!!!"
