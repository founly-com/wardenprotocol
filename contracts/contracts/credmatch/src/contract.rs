use cosmwasm_schema::cw_serde;
#[cfg(not(feature = "library"))]
use cosmwasm_std::entry_point;
use cosmwasm_std::{
    from_json, to_json_binary, Binary, Deps, DepsMut, Empty, Env, MessageInfo, Response, StdResult,
    Timestamp,
};
use cw2::set_contract_version;
use cw_storage_plus::{Item, Map};

use bindings::msg::WardenMsg;
use bindings::{WardenProtocolMsg, WardenProtocolQuery};

use crate::error::ContractError;
use crate::msg::{ExecuteMsg, FutureResult, QueryMsg};

// version info for migration info
const CONTRACT_NAME: &str = "crates.io:credmatch";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[cw_serde]
pub struct AuditResponse {
    audit_id: u64,
    sample_size: u64,
    clear_count: u64,
    consider_count: u64,
    concern_count: u64,
    timestamp: Timestamp,
}

impl Default for AuditResponse {
    fn default() -> Self {
        AuditResponse {
            audit_id: 0,
            sample_size: 0,
            clear_count: 0,
            consider_count: 0,
            concern_count: 0,
            timestamp: Timestamp::default(),
        }
    }
}

pub const AUDIT_RESPONSE: Item<AuditResponse> = Item::new("audit_response");
pub const FUTURES_MAP: Map<u64, String> = Map::new("futures_map");

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut<WardenProtocolQuery>,
    _env: Env,
    _info: MessageInfo,
    _msg: Empty,
) -> StdResult<Response> {
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    Ok(Response::default())
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut<WardenProtocolQuery>,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response<WardenProtocolMsg>, ContractError> {
    match msg {
        ExecuteMsg::MatchCred { sample_size } => execute_match_cred(deps, env, info, sample_size),
        ExecuteMsg::FutureReady { output } => execute_callback(deps, env, info, output),
    }
}

#[cw_serde]
struct FutureInput {
    audit_response: AuditResponse,
}

#[cw_serde]
struct FutureOutput {
    id: u64,
    output: String,
}

fn execute_match_cred(
    deps: DepsMut<WardenProtocolQuery>,
    env: Env,
    _info: MessageInfo,
    sample_size: u64,
) -> Result<Response<WardenProtocolMsg>, ContractError> {
    let mut audit_response = AUDIT_RESPONSE.may_load(deps.storage)?.unwrap_or_default();

    // Update audit response
    audit_response.audit_id += 1;
    audit_response.sample_size = sample_size;
    audit_response.timestamp = env.block.time;

    // Prepare the future to be executed
    let msg = WardenMsg::ExecuteFuture {
        input: to_json_binary(&FutureInput {
            audit_response: audit_response.clone(),
        })
        .unwrap(),
        output: to_json_binary(&FutureOutput {
            id: audit_response.audit_id,
            output: format!(
                "AI Audit completed for {} samples on {}",
                sample_size, audit_response.timestamp
            ),
        })
        .unwrap(),
    };

    // Save the updated audit response
    AUDIT_RESPONSE.save(deps.storage, &audit_response)?;

    let res = Response::new().add_message(msg);
    Ok(res)
}

fn execute_callback(
    deps: DepsMut<WardenProtocolQuery>,
    _env: Env,
    _info: MessageInfo,
    output: Binary,
) -> Result<Response<WardenProtocolMsg>, ContractError> {
    let result: FutureOutput = from_json(output)?;

    // Parse the AI audit results
    let parts: Vec<&str> = result.output.split_whitespace().collect();
    let clear_count: u64 = parts
        .iter()
        .position(|&r| r == "Clear")
        .and_then(|i| parts.get(i + 1))
        .and_then(|s| s.parse().ok())
        .unwrap_or(0);
    let consider_count: u64 = parts
        .iter()
        .position(|&r| r == "Consider")
        .and_then(|i| parts.get(i + 1))
        .and_then(|s| s.parse().ok())
        .unwrap_or(0);
    let concern_count: u64 = parts
        .iter()
        .position(|&r| r == "Concern")
        .and_then(|i| parts.get(i + 1))
        .and_then(|s| s.parse().ok())
        .unwrap_or(0);

    // Update the audit response with results
    let mut audit_response = AUDIT_RESPONSE.load(deps.storage)?;
    audit_response.clear_count = clear_count;
    audit_response.consider_count = consider_count;
    audit_response.concern_count = concern_count;

    // Save the updated audit response
    AUDIT_RESPONSE.save(deps.storage, &audit_response)?;

    // Save the future result
    FUTURES_MAP.save(deps.storage, result.id, &result.output)?;

    Ok(Response::default())
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(_deps: Deps<WardenProtocolQuery>, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::GetFutureResult { id } => {
            let output = FUTURES_MAP.load(_deps.storage, id)?;
            to_json_binary(&FutureResult { output })
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{mock_env, mock_info, MockApi, MockQuerier, MockStorage};
    use cosmwasm_std::{from_json, OwnedDeps};

    // Add this function to create mock dependencies with WardenProtocolQuery
    fn mock_dependencies_with_warden(
    ) -> OwnedDeps<MockStorage, MockApi, MockQuerier, WardenProtocolQuery> {
        OwnedDeps {
            storage: MockStorage::default(),
            api: MockApi::default(),
            querier: MockQuerier::default(),
            custom_query_type: std::marker::PhantomData,
        }
    }

    #[test]
    fn proper_initialization() {
        let mut deps = mock_dependencies_with_warden();
        let env = mock_env();
        let info = mock_info("creator", &[]);
        let msg = Empty {};

        // Instantiate the contract
        let res = instantiate(deps.as_mut(), env, info, msg).unwrap();
        assert_eq!(0, res.messages.len());

        // Check if the contract version is set correctly
        let version = cw2::get_contract_version(deps.as_ref().storage).unwrap();
        assert_eq!(CONTRACT_NAME, version.contract);
        assert_eq!(CONTRACT_VERSION, version.version);
    }

    #[test]
    fn execute_match_cred_works() {
        let mut deps = mock_dependencies_with_warden();
        let env = mock_env();
        let info = mock_info("user", &[]);

        // Execute match_cred
        let msg = ExecuteMsg::MatchCred { sample_size: 1000 };
        let res = execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // Check if a message was added
        assert_eq!(1, res.messages.len());

        // Check if the audit response was saved
        let audit_response = AUDIT_RESPONSE.load(deps.as_ref().storage).unwrap();
        assert_eq!(1, audit_response.audit_id);
        assert_eq!(1000, audit_response.sample_size);
        assert_eq!(env.block.time, audit_response.timestamp);
    }

    #[test]
    fn execute_callback_works() {
        let mut deps = mock_dependencies_with_warden();
        let env = mock_env();
        let info = mock_info("warden", &[]);

        // First, execute match_cred to set up initial state
        let msg = ExecuteMsg::MatchCred { sample_size: 1000 };
        execute(deps.as_mut(), env.clone(), info.clone(), msg).unwrap();

        // Now execute the callback
        let output = to_json_binary(&FutureOutput {
            id: 1,
            output: "AI Audit completed for 1000 samples: Clear 800 Consider 150 Concern 50"
                .to_string(),
        })
        .unwrap();
        let msg = ExecuteMsg::FutureReady { output };
        let res = execute(deps.as_mut(), env, info, msg).unwrap();

        // Check if the response is empty (as expected)
        assert_eq!(0, res.messages.len());

        // Check if the audit response was updated correctly
        let audit_response = AUDIT_RESPONSE.load(deps.as_ref().storage).unwrap();
        assert_eq!(1, audit_response.audit_id);
        assert_eq!(1000, audit_response.sample_size);
        assert_eq!(800, audit_response.clear_count);
        assert_eq!(150, audit_response.consider_count);
        assert_eq!(50, audit_response.concern_count);

        // Check if the future result was saved
        let future_result = FUTURES_MAP.load(deps.as_ref().storage, 1).unwrap();
        assert!(future_result.contains("AI Audit completed for 1000 samples"));
    }

    #[test]
    fn query_future_result_works() {
        let mut deps = mock_dependencies_with_warden();
        let env = mock_env();
        let info = mock_info("user", &[]);

        // Set up initial state
        let msg = ExecuteMsg::MatchCred { sample_size: 1000 };
        execute(deps.as_mut(), env.clone(), info.clone(), msg).unwrap();

        // Execute callback to save a future result
        let output = to_json_binary(&FutureOutput {
            id: 1,
            output: "AI Audit completed for 1000 samples: 800 Clear 150 Consider 50 Concern"
                .to_string(),
        })
        .unwrap();
        let msg = ExecuteMsg::FutureReady { output };
        execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // Query the future result
        let query_msg = QueryMsg::GetFutureResult { id: 1 };
        let res = query(deps.as_ref(), env, query_msg).unwrap();
        let future_result: FutureResult = from_json(&res).unwrap();

        assert!(future_result
            .output
            .contains("AI Audit completed for 1000 samples"));
    }
}
