# Neutral Client Java Demo

This demo shows a simplistic contributor app that converts some amount of
coin (one defined for the NeutralDollar basket) to NUSD at current rate
quoted by the neutral backend service.

Note that the actual settlement is still being worked on and will be included
in an update in a day or so.

## Syntax
The following command line arguments are expected.
coin              - one of TUSD, USDT, PAX and DAI (string)
quantity          - amount of the coins (double)
withdraw-address  - address from which to withdraw the coin (in Hex)
deposit-address   - address to which to deposit NUSD (in Hex)


To run the contributor against the Ganache setup,
`
NTL_ENV=test ./gradlew runContributor --args='TUSD 100 0xf1f0862711fd40259522cabfa4f99e9c80d61f18 0x1a63c94aae69d28ddd810e4e9e3113eeecfcfe2d'
`

To run the contributor against the Ropsten setup,
`
NTL_ENV=ropsten ./gradlew runContributor --args='TUSD 100 0xf1f0862711fd40259522cabfa4f99e9c80d61f18 0x1a63c94aae69d28ddd810e4e9e3113eeecfcfe2d'
`
