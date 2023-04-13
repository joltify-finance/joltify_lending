#!/usr/bin/env bash
# Path: contrib/devnet/submit_proposal.sh

# Params
submitter_addr="jolt1uvpeue6qtp27p63d6p68rayklpau7y99cjqpyr"
validator_addr="jolt10jghunnwjka54yzvaly4pjcxmarkvevzvq8cvl"
proposal_path="./proposal.json"
proposal_submitter_name="submitter"
proposal_validator_name="validator"
proposal_voter1="voter1"
proposal_voter2="voter2"
proposal_voter3="voter3"
proposal_voter4="voter4"
proposal_voter5="voter5"
vote_yes="yes"
vote_no="no"
vote_veto="no_with_veto"
vote_abstain="abstain"
chain_id="joltifylocalnet_8888-1"
fee="2000ujolt"
deposit_amount="10000000ujolt"


# check and add accounts
#ret=$(joltify keys list --output json)
#ret=$(joltify keys import validator validator_key)
#ret=$(joltify keys import validator submitter_key)


# transfer coins
#if [];then
#ret=$(joltify tx bank send $validator_addr $submitter_addr 100000000ujolt --from $proposal_validator_name --chain-id $chain_id -y --output json)
#txHash=$(echo $ret | jq ".txhash")
#txCode=$(echo $ret | jq ".code")
#echo "transaction code: "$txCode
#echo "proposal submission tx hash: "$txHash
#fi

# submit proposal
echo ">>>> Submit the proposal (sleep for 15s)"
ret=$(joltify tx gov submit-legacy-proposal param-change $proposal_path --from $proposal_submitter_name --chain-id $chain_id --gas=auto --fees $fee -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal submission tx hash: "$txHash
sleep 15

# find the proposal id
echo ""
echo ">>>> Query the submitted proposal ID"
ret=$(joltify q gov proposals --output json)
proposal_id_string=$(echo $ret | jq ".proposals[-1].id")
#proposal_id=${proposal_id_string:1:1}
proposal_id=${proposal_id_string#"\""}
proposal_id=${proposal_id%"\""}
echo "Proposal ID: "$proposal_id

# deposit for proposal
echo ""
echo ">>>> Deposit for the proposal (sleep for 10s)"
ret=$(joltify tx gov deposit $proposal_id $deposit_amount --from $proposal_submitter_name --chain-id $chain_id --gas=auto --fees $fee -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal deposit tx hash: "$txHash
sleep 10

## vote for proposal
echo ""
echo ">>> Vote for the proposal"

echo "vote from "$proposal_voter1
ret=$(joltify tx gov vote $proposal_id $vote_no --from $proposal_voter1 --chain-id $chain_id  -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal vote tx hash: "$txHash
sleep 2

echo "vote from "$proposal_voter2
ret=$(joltify tx gov vote $proposal_id $vote_veto --from $proposal_voter2 --chain-id $chain_id  -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal vote tx hash: "$txHash
sleep 2

echo "vote from "$proposal_voter3
ret=$(joltify tx gov vote $proposal_id $vote_veto --from $proposal_voter3 --chain-id $chain_id  -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal vote tx hash: "$txHash
sleep 2

echo "vote from "$proposal_voter4
ret=$(joltify tx gov vote $proposal_id $vote_abstain --from $proposal_voter4 --chain-id $chain_id  -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal vote tx hash: "$txHash
sleep 2

echo "vote from "$proposal_voter5
ret=$(joltify tx gov vote $proposal_id $vote_abstain --from $proposal_voter5 --chain-id $chain_id  -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal vote tx hash: "$txHash
sleep 2

echo "vote from "$proposal_validator_name
ret=$(joltify tx gov vote $proposal_id $vote_yes --from $proposal_validator_name --chain-id $chain_id  -y --output json)
txHash=$(echo $ret | jq ".txhash")
txCode=$(echo $ret | jq ".code")
echo "transaction code: "$txCode
echo "proposal vote tx hash: "$txHash
sleep 2

# check proposal status
echo ""
echo ">>> Wait for the proposal to be passed (sleep for 60s)"
sleep 60
ret=$(joltify q gov proposal $proposal_id --output json)
status=$(echo $ret | jq ".status")
echo "Proposal Status: "$status