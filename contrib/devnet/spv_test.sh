 #!/bin/bash
 joltify tx spv create-pool 123 1 0.21 300000000000000000000ausdc  --from validator -y
 
joltify tx kyc upload-investor 44 jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8 --from validator  -y

 joltify tx spv add-investors  0xec230857257653d20acf913f2fa4a410667e86471a0515a88f9cf448d65a46e1 44  --from validator -y

 joltify tx spv  active-pool    0xec230857257653d20acf913f2fa4a410667e86471a0515a88f9cf448d65a46e1  --from validator -y

 joltify tx spv deposit 0xec230857257653d20acf913f2fa4a410667e86471a0515a88f9cf448d65a46e1   30000000000000000000ausdc --from user -y

 joltify tx spv  borrow 0xec230857257653d20acf913f2fa4a410667e86471a0515a88f9cf448d65a46e1     10000000000000000000ausdc --from validator

 joltify q nft nfts --owner jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8


