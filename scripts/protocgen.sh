#!/usr/bin/env bash

# Adapted from https://github.com/cosmos/cosmos-sdk/blob/master/scripts/protocgen.sh

set -eo pipefail

#protoc_gen_gocosmos() {
##  if ! grep "github.com/cosmos/gogoproto => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
##    echo -e "\tPlease run this command from somewhere inside the root folder."
##    return 1
##  fi
#
#  #go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@latest 2>/dev/null
#  go install github.com/regen-network/cosmos-proto
#}

protoc_gen_doc() {
  go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc 2>/dev/null
  go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc 2>/dev/null
}

#protoc_gen_gocosmos
#protoc_gen_doc

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./joltify -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &>/dev/null; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done
cd ..


# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.

cp -r github.com/joltify-finance/joltify_lending/* ./
rm -rf github.com

#go mod tidy -compat=1.18
echo "proto gen done"
exit
