---
sidebar_position: 4
---

# Build a Keychain app

## Overview

In this second part of the tutorial, we'll build upon the basic structure created in Part 1. We'll implement key and signature request handlers, add detailed error handling, and write tests for our Keychain service.

## Prerequisites

Before you start, complete the following prerequisites:

- [Set up a basic Go app](set-up-a-basic-go-app).

## 1. Implement a key storage

First, let's implement a simple in-memory storage for our keys:

```go
import (
    "crypto/ecdsa"
    "sync"
)

type Store struct {
    mutex sync.Mutex
    keys  map[uint64]*ecdsa.PrivateKey
}

func NewStore() *Store {
    return &Store{
        keys: make(map[uint64]*ecdsa.PrivateKey),
    }
}

func (s *Store) Save(id uint64, key *ecdsa.PrivateKey) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.keys[id] = key
}

func (s *Store) Get(id uint64) *ecdsa.PrivateKey {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    return s.keys[id]
}
```

## 2. Implement a key request handler

Now, let's implement the `handleKeyRequest` function:

```go
import (
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta2"
)

func handleKeyRequest(w keychain.KeyResponseWriter, req *keychain.KeyRequest) {
    logger := slog.Default()
    logger.Info("received a key request", "id", req.Id, "key_type", req.KeyType)

    if req.KeyType != v1beta2.KeyType_KEY_TYPE_ECDSA_SECP256K1 {
        logger.Error("unsupported key type", "type", req.KeyType)
        if err := w.Reject("unsupported key type"); err != nil {
            logger.Error("failed to reject the key request", "error", err)
        }
        return
    }

    key, err := crypto.GenerateKey()
    if err != nil {
        logger.Error("failed to generate a key", "error", err)
        if rejectErr := w.Reject("failed to generate a key"); rejectErr != nil {
            logger.Error("failed to reject the key request", "error", rejectErr)
        }
        return
    }

    store.Save(req.Id, key)

    pubKey := crypto.CompressPubkey(&key.PublicKey)

    if err := w.Fulfil(pubKey); err != nil {
        logger.Error("failed to fulfill the key request", "error", err)
        // Note: We can't reject here as we've already generated and stored the key
    }
}
```

## 3. Implement a signature request handler

Next, let's implement the `handleSignRequest` function:

```go
func handleSignRequest(w keychain.SignResponseWriter, req *keychain.SignRequest) {
    logger := slog.Default()
    logger.Info("received a signature request", "id", req.Id, "key_id", req.KeyId)

    key := store.Get(req.KeyId)
    if key == nil {
        logger.Error("key not found", "id", req.KeyId)
        if err := w.Reject("key not found"); err != nil {
            logger.Error("failed to reject the signature request", "error", err)
        }
        return
    }

    sig, err := crypto.Sign(req.DataForSigning, key)
    if err != nil {
        logger.Error("failed to sign", "error", err)
        if rejectErr := w.Reject("failed to sign"); rejectErr != nil {
            logger.Error("failed to reject the signature request", "error", rejectErr)
        }
        return
    }

    if err := w.Fulfil(sig); err != nil {
        logger.Error("failed to fulfil the signature request", "error", err)
    }
}
```

## 4. Update the `main()` function

Update the `main()` function to use the new store:

```go
func main() {
    // ... (previous code remains the same)

    store := NewStore()

    app.SetKeyRequestHandler(func(w keychain.KeyResponseWriter, req *keychain.KeyRequest) {
        handleKeyRequest(store, w, req)
    })
    app.SetSignRequestHandler(func(w keychain.SignResponseWriter, req *keychain.SignRequest) {
        handleSignRequest(store, w, req)
    })

    // ... (rest of the code remains the same)
}
```

## 5. Run the app

To run the complete app:

1. Ensure all the new code is integrated into your `main.go` file.
2. Run the following command:

   ```bash
   go run main.go
   ```

You should see output indicating that the app has started, is connecting to the Warden Protocol node, and is ready to handle key and signature requests.

## 6. Test the app

Let's add some tests to ensure our Keychain service is working correctly. Create a new file named `keychain_test.go` with the following content:

```go
package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/warden-protocol/wardenprotocol/keychain-sdk"
    "github.com/warden-protocol/wardenprotocol/warden/x/warden/types/v1beta2"
)

type mockResponseWriter struct {
    fulfilled bool
    rejected  bool
    data      []byte
    reason    string
}

func (m *mockResponseWriter) Fulfil(data []byte) error {
    m.fulfilled = true
    m.data = data
    return nil
}

func (m *mockResponseWriter) Reject(reason string) error {
    m.rejected = true
    m.reason = reason
    return nil
}

func TestHandleKeyRequest(t *testing.T) {
    store := NewStore()
    writer := &mockResponseWriter{}
    req := &keychain.KeyRequest{
        Id:      1,
        KeyType: v1beta2.KeyType_KEY_TYPE_ECDSA_SECP256K1,
    }

    handleKeyRequest(store, writer, req)

    assert.True(t, writer.fulfilled)
    assert.NotEmpty(t, writer.data)
    assert.NotNil(t, store.Get(1))
}

func TestHandleSignRequest(t *testing.T) {
    store := NewStore()
    key, _ := crypto.GenerateKey()
    store.Save(1, key)

    writer := &mockResponseWriter{}
    req := &keychain.SignRequest{
        Id:             1,
        KeyId:          1,
        DataForSigning: []byte("test data"),
    }

    handleSignRequest(store, writer, req)

    assert.True(t, writer.fulfilled)
    assert.NotEmpty(t, writer.data)
}
```

To run the tests:

```bash
go test -v
```

You should see output indicating that the tests have passed.

## Conclusion

Congratulations! You've now built a fully functional Keychain service using the Warden Protocol. This service can generate ECDSA keys, store them securely in memory, and use them to sign data. Remember to implement proper key storage and management in a production environment, as in-memory storage is not suitable for long-term or secure key management.
