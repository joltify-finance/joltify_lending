#!/usr/bin/env bash
# Path: contrib/devnet/check_balance_spv.sh

echo ">>>> spv balance"
spv1="jolt10hvukzunn0whud5ye5pu6cf4w0sncxp9l2vl4s"
spv2="jolt1fq4g96f7gf68ezz0wfj64xt25us57ynaamnwe4"
spv3="jolt1ez2kqt9t3q786dhl6pjg3a2rn9r0tmqgy6z87m"
spv4="jolt1gyusevw8s4ttayzn6mqewprjhjdtk5cs044zs6"
spv5="jolt1hcrvt8sf6fnhx2y4lgntx3dlmkw0c9g9anl00n"
spv6="jolt1zm7vkj2f5mdsgdz3dedrsamytuqdv8n0qhxgcm"
spv7="jolt1v8uy4rskwueqrapcungl6d9g0vjag02qkqp6rd"
spv8="jolt1snnjxyy6en6u40r538r8q0twmwqnu53ljdhj9c"
spv9="jolt1ym20el8fhsmjmgf53wj64zuduyschwn39e0wfn"
spv10="jolt1mqwc5qcy29zhp66j4kwwzfzlhpenqr98gvazxt"
spv11="jolt1jvgdx789dgls6tlvm3jd37dg5072yhx65ffjhl"
spv12="jolt19haysjznlysv67wqdv4xsj35qfprv9fg4lrwgp"
spv13="jolt1e8jpqqpdrm9h8ekqacp9xc7hp4zxwku2ymgez6"
spv14="jolt1zknnsut5tk9wuqjjzphue4p4x2a9w3kstaep89"
spv15="jolt1gj74n3jxtduluqy4askxpmf5kqnzc6w9lg8pn3"
spvs=($spv1 $spv2 $spv3 $spv4 $spv5 $spv6 $spv7 $spv8 $spv9 $spv10 $spv11 $spv12 $spv13)
count=1
for spv in ${spvs[@]};
do
  echo "SPV"$count": "$spv
  ret=$(joltify q bank balances $spv --output json)
  balances_spv=$(echo $ret | jq ".balances")
  echo $balances_spv
  count=$((count+1))
done

echo ""
echo ">>>> alice balance"
ret=$(joltify q bank balances jolt1p2g5say9w86mzahl2ap6tz6vymk7jmk5wn9skr --output json)
balances_alice=$(echo $ret | jq ".balances")
echo $balances_alice

echo ""
echo ">>>> bob balance"
ret=$(joltify q bank balances jolt1x6kj4pngxvdwqsr6se0h60p8dur05q2dhmsmt7 --output json)
balances_bob=$(echo $ret | jq ".balances")
echo $balances_bob

echo ""
echo ">>>> pool reserve"
ret=$(joltify q spv total-reserve --output json)
reserve=$(echo $ret | jq ".coins")
echo $reserve
