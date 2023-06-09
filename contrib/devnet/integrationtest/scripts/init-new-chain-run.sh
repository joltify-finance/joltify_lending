#! /bin/bash
base=1000000000000000000

validatorMnemonic="equip town gesture square tomorrow volume nephew minute witness beef rich gadget actress egg sing secret pole winter alarm law today check violin uncover"
#jolt10jghunnwjka54yzvaly4pjcxmarkvevzvq8cvl

faucetMnemonic="crash sort dwarf disease change advice attract clump avoid mobile clump right junior axis book fresh mask tube front require until face effort vault"
# jolt1g6q7rff5jdyny0ph6gm6nc9x540gs02tlu7nsp

evmFaucetMnemonic="hundred flash cattle inquiry gorilla quick enact lazy galaxy apple bitter liberty print sun hurdle oak town cash because round chalk marriage response success"
# 0x3C854F92F726A7897C8B23F55B2D6E2C482EF3E0
# jolt18jz5lyhhy6ncjlyty064kttw93yzaulqgkkqwj

userMnemonic="news tornado sponsor drastic dolphin awful plastic select true lizard width idle ability pigeon runway lift oppose isolate maple aspect safe jungle author hole"
# 0x7Bbf300890857b8c241b219C6a489431669b3aFA
# jolt10wlnqzyss4accfqmyxwx5jy5x9nfkwh6kwhjs9

relayerMnemonic="never reject sniff east arctic funny twin feed upper series stay shoot vivid adapt defense economy pledge fetch invite approve ceiling admit gloom exit"
# 0xa2F728F997f62F47D4262a70947F6c36885dF9fa

DATA=~/.joltify
# remove any old state and config
rm -rf $DATA

BINARY=joltify

# Create new data directory, overwriting any that alread existed
chainID="joltifylocalnet_8888-1"
$BINARY init validator --chain-id $chainID --trace
#cp  $DATA/config/genesis.json /home/yb/development/joltify/joltify_lending/contrib/gen_raw.json

# hacky enable of rest api
sed -in-place='' 's/enable = false/enable = true/g' $DATA/config/app.toml

# change port to 0.0.0.0
sed -in-place='' 's/enable = false/enable = true/g' $DATA/config/app.toml

sed -i -E 's|enable = false|enable = true|g' $DATA/config/app.toml

sed -i -E 's|0uatom|0ujolt|g' $DATA/config/app.toml

sed -i -E 's|enabled-unsafe-cors = false|enabled-unsafe-cors = true |g' $DATA/config/app.toml

# enable swagger
sed -i -E 's|swagger = false|swagger = true |g' $DATA/config/app.toml

sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' $DATA/config/config.toml

sed -i -E 's|max_subscription_clients = 100|max_subscription_clients = 1000|g' $DATA/config/config.toml

sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"*\"\]|g' $DATA/config/config.toml


# Set evm tracer to json
sed -in-place='' 's/tracer = ""/tracer = "json"/g' $DATA/config/app.toml

# Set client chain id
sed -in-place='' 's/chain-id = ""/chain-id = "joltifylocalnet_8888-1"/g' $DATA/config/client.toml
sed -in-place='' 's/broadcast-mode = "sync"/broadcast-mode = "block"/g' $DATA/config/client.toml

# avoid having to use password for keys
$BINARY config keyring-backend test

# Create validator keys and add account to genesis
validatorKeyName="validator"
printf "$validatorMnemonic\n" | $BINARY keys add $validatorKeyName --recover
$BINARY add-genesis-account $validatorKeyName 200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000000000000000ausdc

# Create faucet keys and add account to genesis
faucetKeyName="faucet"
printf "$faucetMnemonic\n" | $BINARY keys add $faucetKeyName --recover
$BINARY add-genesis-account $faucetKeyName 2000000000000000ujolt,100000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000000000000000ausdc

evmFaucetKeyName="evm-faucet"
printf "$evmFaucetMnemonic\n" | $BINARY keys add $evmFaucetKeyName  --recover
$BINARY add-genesis-account $evmFaucetKeyName 100000000000000ujolt

userKeyName="user"
printf "$userMnemonic\n" | $BINARY keys add $userKeyName --recover
$BINARY add-genesis-account $userKeyName  200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000000000000000ausdc

$BINARY add-genesis-account jolt1kdgjxwdk4w5pexwhtvek009pnp4qw07f4s89ea   200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000000000000000ausdc

relayerKeyName="relayer"
printf "$relayerMnemonic\n" | $BINARY keys add $relayerKeyName  --recover
$BINARY add-genesis-account $relayerKeyName 1000000000ujolt


for i in {1..100}
do
  a=$(joltify keys add key_$i --keyring-backend test --output json)
  # get the address from the json
  address=$(echo $a | jq -r '.address')

  # transfer amount
  amount=$(echo 5000000*$base|bc)
  amountAtom=10000000

  $BINARY add-genesis-account $address $amount"ausdc",$amountAtom"ujolt"
done



# Create a delegation tx for the validator and add to genesis
$BINARY gentx $validatorKeyName 1000000000ujolt --keyring-backend test --chain-id $chainID
$BINARY collect-gentxs

# Replace stake with ujolt
sed -in-place='' 's/stake/ujolt/g' $DATA/config/genesis.json

# Replace the default evm denom of aphoton with ujolt
sed -in-place='' 's/aphoton/ajolt/g' $DATA/config/genesis.json

# Zero out the total supply so it gets recalculated during InitGenesis
jq '.app_state.bank.supply = []' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

# update the vote
jq '.app_state.gov.voting_params.voting_period = "60s"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

jq '.app_state.distribution.params.community_tax= "0"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
