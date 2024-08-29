---
sidebar_position: 1
---

# Introduction

## What are Keychains?

## Keychain application in Warden

Every [Omnichain Application](/learn/glossary#omnichain-application) has at least one Keychain – a custodian that generates and stores [keys](/learn/glossary#key) and signs transactions. Keychains contribute to Warden's [Modular Key Management](/learn/glossary#modular-key-management) and [Modular Security](/learn/glossary#modular-security).

## Keychain operators

The Warden Protocol allows users or external organizations become Keychain operators. They can onboard their own Keychains and charge fees for [key requests](/learn/glossary#key-request) and [signature requests](/learn/glossary#signature-request). Note that Keychain operators typically use MPC networks to generate keys and signatures. To learn more, see [Request flow](/learn/request-flow).

## Get started

To get started, follow the steps in the [Quick start](quick-start) guide. This guide will walk you through the basics of interacting with Keychain.

In the next section, [build an app](build-an-app)  you will walk through building an Application that uses features of Keychain. The first part of the tutorial will help you get started with basic understanding. In the second part of the tutorial, you will learn more advanced implementations of Keychain.

If you are an advanced user, you would also like to read our [Implementation](examples) guide. This guide features details about `cli-chain`, `warden-kms` and `keychain-sdk`.

In the future, we're going to provide off-chain infrastructure to facilitate Keychain deployment.
