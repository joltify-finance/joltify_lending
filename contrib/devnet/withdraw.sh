#!/usr/bin/bash



  ret=$(joltify q spv  list-pools --output json)
  # get the index of the pool
  indexSenior=$(echo $ret | jq -r '.pools_info[0].index')
  indexJunior=$(echo $ret | jq -r '.pools_info[1].index')

  ./run_deposit $indexJunior 5 0 300000 true
  ./run_deposit $indexSenior 5 0 300000 true
