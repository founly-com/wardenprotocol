---
sidebar_position: 2
---

# Quick start

## Overview

This guide will walk you through the basics of interacting with Keychain.

To follow along this tutorial, you will need to run the **warden** chain locally or connect to the buenavista testnet.

### Run a node

- [Run a local chain](/build-an-app/test/run-a-local-chain)
- [Connect to our Buenavista testnet](/operate-a-node/buenavista-testnet/join-buenavista)

For the rest of this guide, we'll assume you have **warden** node running with a local account that has a few WARD tokens. The local account will be used to fund the Keychain and its Writers and referenced as `<your-key>` in the following commands.

1. Check the list of available accounts by running this command:

    ```bash
    wardend keys list
    ```

2. This should return output similar to:

    ```bash
    - address: warden10kmgv5gzygnecf46x092ecfe5xcvvv9r870rq4
    name: shulgin
    pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Atj5FRYRKgo0mUuj6SISUvzKTk6qYnLwwjLxU0JCk3VM"}'
    type: local
    - address: warden1u6lmp3jkke4jl3qr6wd8emvc0ph3asdrgym3dr
    name: val
    pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A3DHjBwGDNSV7Aye1lNK9iDHf9q83AQK5gkyHaIm+LHo"}'
    type: local
    ```

3. Then, you can check the local account balance:

    ```bash
    wardend query bank balances <your-key>
    ```

4. This should return output similar to:

    ```bash
    balances:
    - amount: "10000000000000000000"
        denom: uWARD
    pagination:
        total: "1"
    ```

:::tip

In development genesis files, you'll typically find an account named `shulgin` that is ready to be used.

:::

## Create a Keychain entity

You need to register your Keychain entity on-chain.

1. Initiate a `MsgNewKeychain` transaction by running this command:

    ```bash
    wardend tx warden new-keychain \
      --description 'My Keychain' \
     --keychain-fees "{\"key_req\":[{\"amount\":\"100\",\"denom\":\"uward\"}],\"sig_req\":[{\"amount\":\"1\",\"denom\":\"uward\"}]}" \
      --from <your-key> \
      --chain-id warden
    ```

    Specify the following details:

    - `description` (optional): The Keychain description
    - `keychainFees`(optional):
         - `key_req`: A fee in uWARD for creating a key pair
         - `key_req`: A fee in uWARD for signing a transaction
    - `from`: Your local account
    - `chain-id`: The chain ID – `warden`

2. A new Keychain object will be created on-chain with its dedicated [Keychain ID](/learn/glossary#keychain-id). Get the ID:

    ```bash
    wardend query warden keychains
    ```

    This should return output similar to:

    ```bash
    keychains:
    - admins:
        - warden1uxlhv5wtyuka38pnek3u4ytfglhzu3w5dqp9gh
    creator: warden1uxlhv5wtyuka38pnek3u4ytfglhzu3w5dqp9gh
    description: OCP KMS
    id: "1"
    writers:
        - warden16zhve0rhdf382apymhdcwz2stqzgvx37ds5sck

    - admins:
        - warden1j6jey3j7djvmjw3vykxp2nl0xauashyphua0hu
    creator: warden1j6jey3j7djvmjw3vykxp2nl0xauashyphua0hu
    description: Open Custody Protocol (Fordefi MPC)
    id: "2"
    writers:
        - warden1u289xle6fs4w5dh5phx5xca4w2cq8548ex7l60
        - warden1xzhguvv99pvs725w7reaad58wktryda47j784t
        - warden1mtppkfxc4nzchgmyzd99ptysqfc0y862hf42dj
    ```

    This ID will be used in the next steps, referenced as `<keychain-id>`.

### Add a Keychain Writer

A Keychain Writer is an account that can write Keychain results (public keys and signatures) to the chain. The Keychain Writers list is essentially an allowlist of accounts that can interact on behalf of the Keychain.

To add a Keychain Writer, take these steps:

1. Initiate a `MsgAddKeychainWriter` transaction by running the following command:

    ```bash
    wardend keys add my-keychain-writer
    ```

2. The output will be similar to the following:

    ```bash
    - address: warden18my6wqsrf5ek85znp8x202wwyg8rw4fqhy54k2
      name: my-keychain-writer
      pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2cECb3ziw5/LzUBUZIChyek3bnGQv/PSXHAH28xd9/Q"}'
      type: local
    
    
    **Important** write this mnemonic phrase in a safe place. It is the only way to recover your account if you ever forget your password.
    
    virus boat radio apple pilot ask vault exhaust again state doll stereo slide exhibit scissors miss attack boat budget egg bird mask more trick
    ```

    :::tip
    Only the Keychain Writer address, returned as `address`, will be able to publish signatures and public keys on behalf of the Keychain.
    :::

3. Note down the mnemonic phrase and the address of the new account, it'll be needed to configure the Keychain SDK and interact with the chain using this account.

4. Fund the new account with some tokens (1 WARD in this example):

    ```bash
    wardend tx bank send <your-key> \
      $(wardend keys show -a my-keychain-writer) \
      1000000uward \
      --chain-id warden
    ```

## Conclusion

This guide covered the basics of setting up and interacting with a Keychain on the Warden blockchain. The next sections will walk you through building an application that uses the Keychain.
