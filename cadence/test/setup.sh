#!/bin/bash

#
# 0x01cf0e2f2f715450 (admin for testing)
#
# Private Key 16bce0227d2221004a4c1e34d5f73530b33b795059abbf2cde196164ab658dc5
# Public Key  3df9b7d537c6409a95e260d57703c5742785127ed3746bf1023e42a15a80665c26ffaf843a23a3947105c1cc561085da17eaf06ae95cf24638765a07acefbb4d 
flow accounts create \
    --key 3df9b7d537c6409a95e260d57703c5742785127ed3746bf1023e42a15a80665c26ffaf843a23a3947105c1cc561085da17eaf06ae95cf24638765a07acefbb4d \
    --signer emulator-account

flow transactions send ./transactions/create_dit.cdc --signer admin "dc=example"

#
# 0x179b6b1cb6755e31 (org0)
#
# Private Key 917fedc7e1808eac8e0193da5f2b12ee675ab365c570a3fbf7c4b172ca29633d
# Public Key  2b150d20df1b1185614a643916fe6478e0512e75db252462c393b1844032b93cc1139a4b5ced2d09426152e5ecf6466ee360a9dc3a7323d7c5bde0f7aff66ec5 
flow accounts create \
    --key 2b150d20df1b1185614a643916fe6478e0512e75db252462c393b1844032b93cc1139a4b5ced2d09426152e5ecf6466ee360a9dc3a7323d7c5bde0f7aff66ec5 \
    --signer emulator-account

flow transactions send ./transactions/create_dit.cdc --signer org0 "dc=org0,dc=example"

#
# 0xf3fcd2c1a78f5eee (org1)
#
# Private Key aa21c21b460670e44b5ee067bcc8fe86321795e60672f9fe16351770c7b04239
# Public Key  fe8164c961e6790ed8e98893ccf2da55f4dbd85ee5c9e7966dbd69609329ee29417743035fa7a731ebaf4ea684dce9a2c7f63b6924338a5025ee89052e56918b 
flow accounts create \
    --key fe8164c961e6790ed8e98893ccf2da55f4dbd85ee5c9e7966dbd69609329ee29417743035fa7a731ebaf4ea684dce9a2c7f63b6924338a5025ee89052e56918b \
    --signer emulator-account

flow transactions send ./transactions/create_dit.cdc --signer org1 "dc=org1,dc=example"

#
# 0xe03daebed8ca0615 (user0)
#
# Private Key beb4c294bd27844618ef1c193aed63cba4144091aafad92f0047aaff4a682fdb 
# Public Key  b9d9ef88f329ce9a489ac1eb72efe6a4f8debb82368e76fd92269d942ceeec85b27ba3df5d839303d61f4969626910cd7223062c84a6ff9cca305f8c1636dd83
flow accounts create \
    --key b9d9ef88f329ce9a489ac1eb72efe6a4f8debb82368e76fd92269d942ceeec85b27ba3df5d839303d61f4969626910cd7223062c84a6ff9cca305f8c1636dd83 \
    --signer emulator-account

flow transactions send ./transactions/create_dit.cdc --signer user0 "cn=taro.yamada,dc=person,dc=example"

#
# 0x045a1763c93006ca (user1)
#
# Private Key b5b4c87417d61b9e314930b778246ab28facfff93fdc33d7ab27745e60d3a4a4
# Public Key  583501f7c72b66b22c4cfb1dc484147c0e9c536f064f6b78784b0bb9ab32b6bce929817a61ff65400f3b2df6aa435ebf3b05de27505cbba51a608bb15b875406
flow accounts create \
    --key 583501f7c72b66b22c4cfb1dc484147c0e9c536f064f6b78784b0bb9ab32b6bce929817a61ff65400f3b2df6aa435ebf3b05de27505cbba51a608bb15b875406 \
    --signer emulator-account

flow transactions send ./transactions/create_dit.cdc --signer user1 "cn=jiro.sato,dc=person,dc=example"
