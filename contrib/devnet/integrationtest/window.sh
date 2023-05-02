#!/bin/bash
# generate n cosmos accounts
# loop function to generate n cosmos accounts
base=1000000000000000000
total_investors=2
all_keys=5
project_index=1
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



allpoolindex="0x7e92c40bccd50c240fc4bca3262102dc85f20b4b40b1733470285f0645ad0f48,0x43ce7e072884180e125328e727911ad83fcaba1cc487ece1ccc3e19376f51118,0x2b8bc05f5ae783adf4d4e0bb2ae6d25e81c1cf0e5a88c35c47f972a4f7577e79,0x50650186dca8ad51702dd1b99a5920c3fdd3e06ea611624f9afb046e2484fd7a,0x90e88af9988dce50bb93c2a9455121c72c3e9fa96c4dcd8d65946bd7ad9d642f,0xc7f66b53c927b81e1e7ce4e258c63ade57f642736a7da94ef69faf5ea8ef17e5,0x056a9bb86bc07e564a6cc6d03a6731581210c8755bbe7b291149ad8580fca8a5,0xcd7fcfce5d9cd88848e25dc9b02580a505b82856357060d648677b27668ec7ff,0x52ff805f28612937a69afbe70c066805034edbc35c9963ac43b7c63c36d57854,0xeb5b95144c1b3410aee69327b63a6d2b6284ed241e874c5178f75a0705093597,0x1e73a606de1e752295b0d870def94b80b9c9ed8b949b88b3f2687d9033251ec5,0x33aa147f2c7fb43587fb950ba87eecec17668a5da603fcac32f370e39196fba8,0x0805921f9be9b80e696399a3809be44eac1ad2501749c850d8475aaaf8ddcc0c,0xf01c9820eecde4ffbbbe88ffaa06aab7005dbe81c70e5b2e9cbfb5dba76bacce,0x849a05be4f25844afb302172c3ff50d0e135e42dffd7c163145eab068ff4c225,0xa6b8427640f1fcef5ec9ba9c759758dcd40518bf3e3c32c7ab6f8f5b17d64ce1,0x7f2d390045d08d20d794dc63142c896c5074e1f2792b5d150fc7dd7de34405a7,0x724f8f909386c8b875bda036bf84d7f0b50423f4d3588223c2d0c345cb9872e2,0xf456ba8abd9f7b197f289d5e324f232b795251216ebff6c67e7f6a0bfeb1375a,0x562aded3eb6fd6cc727bccf03f814078f3170c64a069fec9f1d864b492a8c9cc,0x46cc7460c2660bdff15d2d421205b2a8d254ede177492ed4a326a8b963311139,0x1262a52672a1400d77f1327238a65cb10feee41954654920110f18762ed892eb,0xc08add325ea9dfc36f45a824a6a752e9f90302bd7f08cfb53b32231b22a53fee,0x001ef59b9a96f6b53fa8ca1638c07e759b2d011eb5f749f9b581a454055495c6,0x1e00ebd829117d1184a5205c236af57396f499df70b4120a93da36deb03a61b3,0x0673362d9eceab7008812a3c504f7bf6ead94ea054be94b134f8ce31294fd788,0x36593d2b6f2adad0f1f2187aa2a89d2a31271dd11d7128cb49b588ad950f13c8,0x638757b1fd0075415be8e6ddf5cf12d029d70dc9c0b47c155b868bbdf8f24cf1,0xb767cb86d5b311f839f83a0e498f567d09abc839515a0672f137fc021a65cd50,0x0b801183a4db584e37cc09b9d0522ef749723dae7e9403f8dd7e56ba8646e3a4,0xe275df88b205790ba4a583c271b6c55f7174583ec61a36a1aaf2056c93e914fb,0x762175d35bec048291f35f3633c8417c912fe0b9ae82835e00ce6b8d070c60c8,0x5f4ce3560108cb8e9ecd04645d86e348d7ea96da7db8bebcbbbd1ac9eb8ea523,0x8cb513e2b25e588c5b8103ef93205b0ff75971c9968c3623002c090f1d0e06af,0x4cf10757dcdc54a4d85d66733734f081475385b604792323eb8658bbbf42f3f4,0x8f88fabf27b113f84c34f08c0547ad1bf966dba580f8849a779c17b23eed0886,0x17aa6a560b50a90684ae89a0a13c2165cee3e4adb8a2a862eb48dd1a4f7702b3,0x1d8399f904359a39932af6158c7872839e2429bba75ce47c108483e4f7c8bf9d,0x9faa154d7f021eecf42bf06fcd48b740e5abc22cdba245fcfd512f804ae6f05b,0x01cde14ad1801527ab4a45ce3760ead371c98934dc981bc26bb9e50dd4ea7a39,0x7ea12c90777fc600cca55da8098d47f7c8d5a482e10f323a4931d71c6c9748c6,0x0fd2f9bed59e7c17f914fd2e354dbdd1eb049f469246c675cac6674549bccd17,0x67ee2915a7465197849dc2e1ef3666a90397d492e32435be0b8de88558a896e3,0xa9db866ee0b074b3ea4a9f19f3076b2322854ac46fc0e263b957be7c128f581a,0x803d1918eb3951f0e50e3bab1700ea3c71bcbe42509de48d269a60b7182fd47e,0xd66868a16f863fa3964e0a2bf5ed3a7113cc753a3c6d3a6850431f15db196fe8,0x76d9b313d79510669ee5ae945b39e212bf04f8af14f956e667c497d5410960c3,0xb73a4870dd394191464d96b3f715a60864cbd66e4a13dc8dda1f28c092904e1f,0x997525e690e8a930b6e5d12a64dae2c60d6e1bf70654ed3d2a523759a4ea6f50,0x5b488025e000a2a63ec2ac4db70e56475f875e66be7864babafd4452126cb2fc,0x7d52eb1ba8f30fe47dbf9a5f6081d6d3bc3f96555922e8d1d2f6eb1e15c547ba,0xeefa350aef7b1910824e00eac61ea00a5f0cf7f409efade5828ad952285043f1,0x3804a7712444132c5e921cd927cd753e932fce106b08501267bcc3e1f56c2442,0xaceb104621cad2e757b2c2b53d03a4d413d403cb0d768cab2c8ecc5d4e370149,0xd2d424c9766f26565bd60fc3744dccaf6bda4198b8e020a3c4dc69fde7d040b2,0xbb9fc61e0d3eb83d9d3325e6b363371f7ae3260d0522c021c2b8dc38de65bcd7,0x570448a144edf763207930e5670b582f67070114fcef74b74a237cc18840b318,0xf2e5b6be5625af98a9534b89442c0aeedee490806ad60d4509581f03427533c7,0xf1417cd97de2898ba17c4954ac0cb30e40fdead761a507fbb86482e307a7b0cf,0x8b88961b8b7419724eba01b537f4fd33f9f3c06a7f3997a36a59fad50ee67a09,0x5dd207860029f86043330d9cb9629967290b71927607e534179be3612399acbe,0xcf1fa2d383ca3df71edf84e1b0fedd85deb206477f864e74a4d73c9ada2b01c4,0x15e2ad29f9f41e503c10909dd204ba920d72e17807ae84afccb2494c4bd01dc2,0x877d0ba5666745e11f88f5e26d9692b14852fb1ced7c4fff8da195ba0ccfcbad,0xe9550e1014172fba76d9bc4732290e91212ca03d6ccac7ee35d8c1ad9424b0ae,0xef2dd54dbce900e82942d08ef625321f4da6f16fcd0228c8327b1f81f5a5a149,0x3a9dcc2d15731c7adb7958a9603c27fd87b5e9f25d63fb2e376e8b8c38b79282,0x9f4e078383c8390521dc24c67aaeca663857413d8f76f895ab7ebf7912914a57,0xfa534ff3f3c303b5cb9a13d60c60ff5fe3f3dad9ad3e666860432acabf1ece5c,0x6023b2c0cefc570286f0ade389efc04a0f801f701424055c5c1eaa46dfb65dab,0xf7af78862c7578b7f766e9e31a4de29a1b3dd97bfbd71e25bf69d375d79bcb15,0x00a318a1cef107cff462893fd23f23cb75971d164d0ed3fdc54f531f87804e4a,0x8a230005ae906deaa127464ff1a5cf86ae5d81359b83249b9c28508db2aef6b9,0x2186b708ff2d4c4cedab7251f096d178f886e30285f6802a443e39f46ac7b3d2,0x281a8a740b8120b7ebe6fcf3f3b27dcdff2537acd4a80c7d136d6a8ebe88f634,0xc80fa0f3008a254eb910c2b49620874afe47e129eda7d27a0437f97b771df697,0x463d326a0ebd881bde4fa3cc06630f982e6c2f3b29c15fcf697473954bb7bb11,0x2b9f432baab06935514cd180efee7cafb3ab7fa65b296b272543b946244f422f,0x00debf42f633a719ebdc0f13e5ca41567f239f1751be2f981c6236d97a05e58e,0xbde36d8e000ee5422cac02cf8b00de5fe89a82185d2238af3c7f353ab9549d53,0x02da84cddf364f5c0a64a7b02eab5d0ac0326954ffb0f00d4a6cf1e53c6dc9aa,0x0038fe031282e57ecec80a6a93b63143b5e83f719964e55429841cae5605ab43,0x4e7dc9b4a19936fae085a71ad584639df5ef4c76a64dddf58a73e9ef09378f2d,0x94f4613f03ff73f52c27852f7a2b7eceed47680c7ddded792adc48e5a7d65499,0xa189b20a8398008d8853ffa396791f7d949d4f31d30aae00695c5b5a8582733e,0x654d3160e532deac6a6d49c5387d96a504fbb0aac1a0cde5e7c5f06fe638d1af,0xeccbfa4fd1d1911b626128f22b694635d19666a1908801e3fec7fe89025381d8,0x12fe4fc2440ba35654233f3e5e0a564e2911405773e7f5a0eb74c88aa1fef8dc,0x1da0cbd6c075115032b11866a5028ba1b5f328f87e12898bb1ecbb1ce56e07d0,0x385d8fe2a2a04eb8a8b14f3e8351c7e7120c80d17c7a7c75354e8cadabd9d4e6,0x32451728a5daa22ecaec38e3c7b7098ba508251222a07a1666de564b1f0c906e,0xe73ac37fa691fbc3eb66f1cb3fff2d8d014bd16050adf88611216ad8e41fcfc7,0x0eb704623455266b2bc409bb64e663b2aa1a1f14ec46fb8ae019120819111ef8,0xfcb6c27ac54eb7922faf77f326b9c4dfd6cb0bd71e2aecf690f137b52a1bff08,0xf117c97a745920e9d3b4df0e44ec3d9f501ea21db57cef663e30475aa7ed5a41,0xa5395cac3b4b694e41df1b2b99678205eff4360116f738a6aeae57811a1d5c4f,0xfa057e035bd894d7636a93ad55eef3133f044d7f75b816d0ad3a8ef9aa0bd1b3,0xe897026ea4a357a76c28add2f031da7d58e431655f01811a502dbd10dfa50b6d,0x33f5ebea7e2fdf224d6aae17ac7b759be4e2f210898409a7fa52f70c474cead0,0x881173228e51d5d09771ea298584cc1d20c88df464e15a60d806b252b467e1c5"
IFS=',' read -ra arr <<< "$allpoolindex"
thispool=${arr[1]}
indexJunior=$1

# we get the pool pay freq
ret=$(joltify query spv query-pool $indexJunior --output json)
pool_pay_freq=$(echo $ret |   jq -r '.pool_info.pay_freq')
project_length=$(echo $ret |   jq -r '.pool_info.project_length')
project_due_time=$(echo $ret |   jq -r '.pool_info.project_due_time')
withdraw_window=$(echo $ret |   jq -r '.pool_info.withdraw_request_window_seconds')
last_payment_time=$(echo $ret |   jq -r '.pool_info.last_payment_time')


#echo "pool pay freq $pool_pay_freq"
#echo "project due time $project_due_time"
#echo "#### last payment time $last_payment_time"
#echo "#### project length $project_length"
# Input duration in seconds
duration=$pool_pay_freq

# Convert last_payment_time and duration to Unix timestamps
last_payment_timestamp_seconds=$(date -u -d "$last_payment_time" +%s)
duration_seconds=$duration

#echo "******************************"
#echo $duration_seconds
#echo "******************************"

# Add duration to last_payment_time
next_payment_seconds=$((last_payment_timestamp_seconds + duration_seconds))

# Convert result back to ISO 8601 format
result=$(date -u -d "@$next_payment_seconds" +"%Y-%m-%dT%H:%M:%SZ")

# Output result
#echo "Next payment time: $result"


# check the time difference between result and current time
# Input duration in seconds
current_time=$(date -u +%Y-%m-%dT%H:%M:%SZ)
#echo "current time $current_time"

current_time_seconds=$(date -u -d "$current_time" +%s)

gap=$((next_payment_seconds - current_time_seconds))


project_due_time_seconds=$(date -u -d "$project_due_time" +%s)
withdrawal_start_duration=$(echo $withdraw_window*3 | bc)
pay_partial_start_duration=$(echo $withdraw_window*2 | bc)
proposal_start_time=$((project_due_time_seconds - withdrawal_start_duration))
pay_partial_start_time=$((project_due_time_seconds - pay_partial_start_duration))


gap2=$(echo $current_time_seconds - $proposal_start_time |bc)

gap3=$(( current_time_seconds - pay_partial_start_time ))

#echo "project due time: $project_due_time"

result=$(date -u -d "@$proposal_start_time" +"%Y-%m-%dT%H:%M:%SZ")
#echo "#### proposal start time $result"
result=$(date -u -d "@$pay_partial_start_time" +"%Y-%m-%dT%H:%M:%SZ")
#echo "#### pay partial start time $result"


#cecho "GREEN" "Next interest payment in : $gap seconds"
#cecho "GREEN" "Next submit principal request in : $gap2 seconds (+ means: already passed)"
#cecho "GREEN" "Next pay partial in : $gap3 seconds (+ means: already passed)"
echo "$gap,$gap2,$gap3"



# check the time difference

