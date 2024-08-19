#! /bin/bash
set -x

source "./genesis.sh"
base=1000000000000000000
base2=1000000


validatorMnemonic="equip town gesture square tomorrow volume nephew minute witness beef rich gadget actress egg sing secret pole winter alarm law today check violin uncover"
#jolt1a33x0juy5t8a0zgksfz50yluw8jyvy764p9ych

faucetMnemonic="crash sort dwarf disease change advice attract clump avoid mobile clump right junior axis book fresh mask tube front require until face effort vault"

evmFaucetMnemonic="hundred flash cattle inquiry gorilla quick enact lazy galaxy apple bitter liberty print sun hurdle oak town cash because round chalk marriage response success"
# 0x3C854F92F726A7897C8B23F55B2D6E2C482EF3E0

userMnemonic="news tornado sponsor drastic dolphin awful plastic select true lizard width idle ability pigeon runway lift oppose isolate maple aspect safe jungle author hole"
# 0x7Bbf300890857b8c241b219C6a489431669b3aFA

relayerMnemonic="never reject sniff east arctic funny twin feed upper series stay shoot vivid adapt defense economy pledge fetch invite approve ceiling admit gloom exit"
# 0xa2F728F997f62F47D4262a70947F6c36885dF9fa

DATA=~/.joltify
# remove any old state and config
rm -rf $DATA

BINARY=joltify

# Create new data directory, overwriting any that alread existed
chainID="joltify_1729-1"
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
sed -in-place='' 's/chain-id = ""/chain-id = "joltify_1729-1"/g' $DATA/config/client.toml
sed -in-place='' 's/keyring-backend = "os" = ""/keyring-backend = "test"/g' $DATA/config/client.toml
sed -in-place='' 's/broadcast-mode = "sync"/broadcast-mode = "sync"/g' $DATA/config/client.toml

# avoid having to use password for keys
$BINARY config  set client keyring-backend test

# Create validator keys and add account to genesis
validatorKeyName="validator"
printf "$validatorMnemonic\n" | $BINARY keys add $validatorKeyName --recover --keyring-backend test
$BINARY genesis add-genesis-account $validatorKeyName 200000000000000000ujolt,200000000000000000uoppy,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3,123456usd-ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3
# Create faucet keys and add account to genesis
faucetKeyName="faucet"
printf "$faucetMnemonic\n" | $BINARY keys add $faucetKeyName --recover
$BINARY genesis add-genesis-account $faucetKeyName 2000000000000000ujolt,100000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3

evmFaucetKeyName="evm-faucet"
printf "$evmFaucetMnemonic\n" | $BINARY keys add $evmFaucetKeyName  --recover
$BINARY genesis add-genesis-account $evmFaucetKeyName 100000000000000ujolt

userKeyName="user"
printf "$userMnemonic\n" | $BINARY keys add $userKeyName --recover
$BINARY genesis add-genesis-account $userKeyName  200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3

$BINARY genesis add-genesis-account jolt1kdgjxwdk4w5pexwhtvek009pnp4qw07f4s89ea   200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3


#ibc test account
$BINARY genesis add-genesis-account jolt1nlrlywakama45q59cqfx3sksf4xdkup6d439zk 200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3


# Obtained from `authtypes.NewModuleAddress(subaccounttypes.ModuleName)`.
SUBACCOUNTS_MODACC_ADDR="jolt1v88c3xv9xyv3eetdx0tvcmq7ung3dywph9jkty"
REWARDS_VESTER_ACCOUNT_ADDR="jolt1wtws9xa2v5f4r2zncnlg273mr0nda5xc3qx42k"
BRIDGE_MODACC_ADDR="jolt1zlefkpe3g0vvm9a4h0jf9000lmqutlh93hptrj"


$BINARY genesis add-genesis-account $SUBACCOUNTS_MODACC_ADDR 200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3
$BINARY genesis add-genesis-account $REWARDS_VESTER_ACCOUNT_ADDR 200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3
$BINARY genesis add-genesis-account $BRIDGE_MODACC_ADDR 200000000000000000ujolt,100000000000000000000000000000abnb,100000000000000000000000000000ausdt,100000000000000000ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3



relayerKeyName="relayer"
printf "$relayerMnemonic\n" | $BINARY keys add $relayerKeyName  --recover
$BINARY genesis add-genesis-account $relayerKeyName 1000000000ujolt


for i in {1..7}
do
  a=$(joltify keys add key_$i --keyring-backend test --output json)
  # get the address from the json
  address=$(echo $a | jq -r '.address')

  # transfer amount
  amount=$(echo 5000000*$base2|bc)
  amountAtom=10000000

  $BINARY genesis add-genesis-account $address $amount"ibc/65D0BEC6DAD96C7F5043D1E54E54B6BB5D5B3AEC3FF6CEBB75B9E059F3580EA3",$amountAtom"ujolt"
done

$BINARY config  set app minimum-gas-prices "0ujolt"


# Create a delegation tx for the validator and add to genesis
$BINARY genesis gentx $validatorKeyName 1000000000ujolt --keyring-backend test --chain-id $chainID
$BINARY genesis collect-gentxs

# Replace stake with ujolt
sed -in-place='' 's/stake/ujolt/g' $DATA/config/genesis.json

# Replace the default evm denom of aphoton with ujolt
sed -in-place='' 's/aphoton/ajolt/g' $DATA/config/genesis.json

# Zero out the total supply so it gets recalculated during InitGenesis
jq '.app_state.bank.supply = []' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

# update the vote
jq '.app_state.gov.voting_params.voting_period = "60s"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.gov.params.voting_period = "60s"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
#jq '.app_state.gov.params.expected_voting_period = "60s"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json


jq '.app_state.distribution.params.community_tax= "0"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

jq '.consensus_params.block.max_gas= "8000000000"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json



#jq '.app_state.feemarket.params.base_fee= "100"' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

addr1=$(joltify keys show -a key_1)
addr2=$(joltify keys show -a key_2)
addr3=$(joltify keys show -a key_3)
addr4=$(joltify keys show -a key_4)
addr5=$(joltify keys show -a key_5)
addr6=$(joltify keys show -a key_6)
addr7=$(joltify keys show -a key_7)
# Define all test accounts for the chain.
TEST_ACCOUNTS=(
#	"dydx199tqg4wdlnu4qjlxchpd7seg454937hjrknju4" # alice
#	"dydx10fx7sy6ywd5senxae9dwytf8jxek3t2gcen2vs" # bob
#	"dydx1fjg6zp6vv8t9wvy4lps03r5l4g7tkjw9wvmh70" # carl
#	"dydx1wau5mja7j7zdavtfq9lu7ejef05hm6ffenlcsn" # dave

$addr1
$addr2
$addr3
$addr4


)

FAUCET_ACCOUNTS=(
#	"dydx1nzuttarf5k2j0nug5yzhr6p74t9avehn9hlh8m" # main fauceta
$addr5
)

# Addresses of vaults.
# Can use ../scripts/vault/get_vault.go to generate a vault's address.
VAULT_ACCOUNTS=(
#	"dydx1c0m5x87llaunl5sgv3q5vd7j5uha26d2q2r2q0" # BTC vault
#	"dydx14rplxdyycc6wxmgl8fggppgq4774l70zt6phkw" # ETH vault
$addr6
$addr7
)
# Number of each vault, which for CLOB vaults is the ID of the clob pair it quotes on.
VAULT_NUMBERS=(
	0 # BTC clob pair ID
	1 # ETH clob pair ID
)
VAL_CONFIG_DIR="$DATA/config"
edit_genesis "$VAL_CONFIG_DIR" "" "${FAUCET_ACCOUNTS[*]}" "${VAULT_ACCOUNTS[*]}" "${VAULT_NUMBERS[*]}" "../dydx_exchange_testdata" "../dydx_delaymsg_config" "" ""
