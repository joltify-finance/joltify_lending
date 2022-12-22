<!--
order: 5
-->

# Parameters

The issuance module has the following parameters:

| Key        | Type           | Example         | Description                                 |
|------------|----------------|-----------------|---------------------------------------------|
| Assets     | array (Asset)  | `[{see below}]` | array of assets created via issuance module |


Each `Asset` has the following parameters

| Key               | Type                   | Example                                         | Description                                           |
|-------------------|------------------------|-------------------------------------------------|-------------------------------------------------------|
| Owner             | sdk.AccAddress         | "jolt1cwhfwfysc4l5guwpffc5ttgvm9dky5snzqt6aa"   | the address that controls the issuance of the asset   |
| Denom             | string                 | "usdtoken"                                      | the denomination or exchange symbol of the asset      |
| BlockedAccounts   | array (sdk.AccAddress) | ["jolt1cwhfwfysc4l5guwpffc5ttgvm9dky5snzqt6aa"] | addresses which are blocked from holding the asset    |
| Paused            | boolean                | false                                           | boolean for if issuance and redemption are paused     |
