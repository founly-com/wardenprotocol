#!/bin/bash

CONTRACT_ADDRESS="warden14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9srt30us"

# interact with the contract

# send a tx with increased gas limit
wardend tx wasm execute $CONTRACT_ADDRESS '{"match_cred":{"sample_size": 1000}}' --from shulgin --gas auto --gas-adjustment 1.3 -y | wardend q wait-tx

# query the state
wardend q wasm contract-state smart $CONTRACT_ADDRESS '{"get_future_result": {"id": 1}}'
