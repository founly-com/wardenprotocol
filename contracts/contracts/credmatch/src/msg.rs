use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::Binary;

#[cw_serde]
pub enum ExecuteMsg {
    MatchCred {
        sample_size: u64,
    },
    FutureReady {
        output: Binary,
    },
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    #[returns(FutureResult)]
    GetFutureResult {
        id: u64,
    },
}

#[cw_serde]
pub struct FutureResult {
    pub output: String,
}