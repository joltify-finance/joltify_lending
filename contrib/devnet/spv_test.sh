 #!/bin/bash
 joltify tx spv create-pool 123 1 0.21 3 10000jolt --from validator -y
 
joltify tx kyc upload-investor 44 jolt1p3jl6udk43vw0cvc5hjqrpnncsqmsz56wd32z8 --from validator  -y

 joltify tx spv add-investors  0x252f479d7a50d5e848c90e95e39ef51d1164f37b97ffefba1eb8b0f745dff689  44  --from validator -y

 joltify tx spv deposit 0x252f479d7a50d5e848c90e95e39ef51d1164f37b97ffefba1eb8b0f745dff689 44  300jolt --from user -y

 joltify tx spv  borrow 0x252f479d7a50d5e848c90e95e39ef51d1164f37b97ffefba1eb8b0f745dff689 200jolt --from validator
