---
sidebar_position: 1
---

# Building a Keychain Service - Part 1: Basics

## Overview

This tutorial will guide you through creating a Keychain service using the Warden Protocol. In this first part, we'll cover the basics, including setting up the project, creating the main structure, and implementing basic configuration.

## Prerequisites

- Go 1.23 or later

## Setting up the project

1. Create a new directory for your project:

   ```bash
   mkdir warden-keychain-service && cd warden-keychain-service
   ```

2. Initialize a new Go module:

   ```bash
   go mod init warden-keychain-service
   ```

3. Install the required dependencies:

   ```bash
   go get github.com/warden-protocol/wardenprotocol/keychain-sdk
   go get github.com/ethereum/go-ethereum/crypto
   ```

## Creating the main application

1. Create a new file named `main.go` in your project directory.

2. Add the following code to `main.go`:

   ```go
   package main

   import (
       "context"
       "log/slog"
       "os"
       "time"

       "github.com/warden-protocol/wardenprotocol/keychain-sdk"
       "github.com/ethereum/go-ethereum/crypto"
   )

   func main() {
       // Set up a logger for debugging
       logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
           Level: slog.LevelDebug,
       }))

       // Create a new Keychain application
       app := keychain.NewApp(keychain.Config{
           Logger:         logger,
           ChainID:        "warden",
           GRPCURL:        "localhost:9090",
           GRPCInsecure:   true,
           KeychainID:     1, // Replace with your actual Keychain ID
           Mnemonic:       "your mnemonic phrase here",
           DerivationPath: "m/44'/118'/0'/0/0",
           GasLimit:       400000,
           BatchInterval:  8 * time.Second,
           BatchSize:      10,
       })

       // Set up handlers for key requests and sign requests
       app.SetKeyRequestHandler(handleKeyRequest)
       app.SetSignRequestHandler(handleSignRequest)

       // Start the application
       if err := app.Start(context.Background()); err != nil {
           logger.Error("application error", "error", err)
           os.Exit(1)
       }
   }

   // handleKeyRequest processes incoming key requests
   func handleKeyRequest(w keychain.KeyResponseWriter, req *keychain.KeyRequest) {
       // To be implemented in Part 2
   }

   // handleSignRequest processes incoming sign requests
   func handleSignRequest(w keychain.SignResponseWriter, req *keychain.SignRequest) {
       // To be implemented in Part 2
   }
   ```

## Error Handling Structure

We'll use the following structure for error handling:

1. Log the error using the `slog` logger.
2. Attempt to reject the request with an appropriate error message.
3. If rejection fails, log this additional error.

Here's a basic example of this structure:

```go
if err := someOperation(); err != nil {
    logger.Error("operation failed", "error", err)
    if rejectErr := w.Reject("Internal error"); rejectErr != nil {
        logger.Error("failed to reject request", "error", rejectErr)
    }
    return
}
```

## Running the App

To run the app:

1. Replace the placeholder values in the `keychain.Config` struct with your actual values.
2. Run the following command:

   ```bash
   go run main.go
   ```

You should see output indicating that the app has started and is connecting to the Warden Protocol node.

## Next Steps

In Part 2 of this tutorial, we'll implement the key and sign request handlers, add detailed error handling, and write tests for our Keychain service.
