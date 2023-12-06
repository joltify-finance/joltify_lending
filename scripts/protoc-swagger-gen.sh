#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen

echo "cloning cosmos v0.47.4"
git clone --depth 1 --branch v0.47.4 https://github.com/cosmos/cosmos-sdk.git tmp_repo

cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk)

cd proto
proto_dirs=$(find ./joltify ../tmp_repo/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 2 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ../
# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

ls ./tmp-swagger-gen
# clean swagger files
rm -rf ./tmp-swagger-gen
rm -rf ./tmp_repo
