#!/bin/bash

  ret=$(joltify q spv  list-pools --output json)
  # get the index of the pool
  indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
  indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

  echo "withdraw $indexJunior"
  ./run_deposit $indexJunior $1 0 300000 true

  echo "withdraw $indexSenior"
  ./run_deposit $indexSenior $1 0 300000 true
