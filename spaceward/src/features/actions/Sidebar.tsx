import clsx from "clsx";
import { hexlify, Transaction } from "ethers";
import { useContext, useEffect } from "react";
import { isDeliverTxSuccess, StargateClient } from "@cosmjs/stargate";
import { useChain, walletContext } from "@cosmos-kit/react-lite";
import { KeyringSnapRpcClient } from "@metamask/keyring-api";
import { cosmos, warden } from "@wardenprotocol/wardenjs";
import { base64FromBytes } from "@wardenprotocol/wardenjs/codegen/helpers";
import { ActionStatus } from "@wardenprotocol/wardenjs/codegen/warden/act/v1beta1/action";
import { useToast } from "@/components/ui/use-toast";

import {
	Popover,
	PopoverContent,
	PopoverTrigger,
} from "@/components/ui/popover";

import { COSMOS_CHAINS } from "@/config/tokens";
import { env } from "@/env";
import { getClient, getSigningClient } from "@/hooks/useClient";
import { useWeb3Wallet } from "@/hooks/useWeb3Wallet";
import { getProvider, isSupportedNetwork } from "@/lib/eth";
import { isUint8Array } from "@/lib/utils";
import "./animate.css";
import { QueuedAction, QueuedActionStatus, useActionsState } from "./hooks";
import { getActionHandler, GetStatus } from "./util";
import { prepareTx } from "../modals/util";
import { TEMP_KEY, useKeySettingsState } from "../keys/state";
import Assets from "../keys/assets";

interface ItemProps extends QueuedAction {
	single?: boolean;
}

function ActionItem({ single, ...item }: ItemProps) {
	const { walletManager } = useContext(walletContext);
	const { data: ks, setData: setKeySettings } = useKeySettingsState();
	const { toast } = useToast()
	const { w } = useWeb3Wallet("wss://relay.walletconnect.org");
	const { setData } = useActionsState();

	const { getOfflineSignerDirect: getOfflineSigner } = useChain(
		env.cosmoskitChainName,
	);

	const type = typeof item.keyThemeIndex !== "undefined" ?
		"key" : (["walletConnectRequestId", "walletConnectTopic"] as const).every(key => typeof item[key] !== "undefined") ?
			"wc" : typeof item.snapRequestId ?
				"snap" : undefined

	useEffect(() => {
		if (item.status === QueuedActionStatus.Signed) {
			const signer = getOfflineSigner();
			let cancel = false;

			getSigningClient(signer)
				.then((client) => {
					if (cancel) {
						return;
					}

					const txRaw = cosmos.tx.v1beta1.TxRaw.encode(item.txRaw);
					return client.broadcastTx(Uint8Array.from(txRaw.finish()));
				})
				.then((res) => {
					if (!res || cancel) {
						return;
					}

					if (isDeliverTxSuccess(res)) {
						setData({
							[item.id]: {
								...item,
								status: QueuedActionStatus.Broadcast,
								response: res,
							},
						});
					} else {
						console.error("Failed to broadcast", res);

						toast({
							title: "Failed",
							description: "Could not broadcast transaction",
							duration: 10000,
						});

						setData({
							[item.id]: {
								...item,
								status: QueuedActionStatus.Failed,
								response: res,
							},
						});
					}
				})
				.catch((err) => {
					toast({
						title: "Failed",
						description: err.message ?? "Unexpected error",
						duration: 10000,
					});

					setData({
						[item.id]: {
							...item,
							status: QueuedActionStatus.Failed,
						},
					});

					console.error(err);
				});

			return () => {
				cancel = true;
			};
		} else if (item.status === QueuedActionStatus.Broadcast) {
			const actionCreatedAny = item.response?.msgResponses[0];
			let actionId: bigint | undefined;

			if (!actionCreatedAny) {
				console.error("no action created");
				return;
			}

			if (item.typeUrl === warden.act.v1beta1.MsgNewAction.typeUrl) {
				const actionCreated =
					warden.act.v1beta1.MsgNewActionResponse.decode(
						actionCreatedAny.value,
					);

				actionId = actionCreated.id;
			} else {
				console.error("unexpected action type", item.typeUrl);
			}

			if (actionId) {
				setData({
					[item.id]: {
						...item,
						actionId,
						status: QueuedActionStatus.AwaitingApprovals,
					},
				});
			}
		} else if (item.status === QueuedActionStatus.AwaitingApprovals) {
			let cancel = false;
			let timeout: number | undefined;

			async function checkAction() {
				if (cancel || !item.actionId) {
					return;
				}

				const client = await getClient();

				const res = await client.warden.act.v1beta1.actionById({
					id: item.actionId,
				});

				if (
					res.action?.status === ActionStatus.ACTION_STATUS_COMPLETED
				) {
					const { action } = res;

					setData({
						[item.id]: {
							...item,
							status: QueuedActionStatus.ActionReady,
							action,
						},
					});
				} else if (
					res.action?.status === ActionStatus.ACTION_STATUS_PENDING
				) {
					timeout = setTimeout(
						checkAction,
						1000,
					) as unknown as number;
				} else {
					console.error("action failed", res);

					toast({
						title: "Failed",
						description: "Unexpected action status",
						duration: 10000,
					});

					setData({
						[item.id]: {
							...item,
							status: QueuedActionStatus.Failed,
						},
					});
				}
			}

			checkAction();

			return () => {
				cancel = true;
				clearTimeout(timeout);
			};
		} else if (item.status === QueuedActionStatus.ActionReady) {
			let cancel = false;
			let timeout: number | undefined;
			let getStatus: GetStatus | undefined;

			try {
				getStatus = getActionHandler(item).getStatus;
			} catch (e) {
				console.error(e);

				toast({
					title: "Failed",
					description: (e as Error)?.message ?? "Action failed",
					duration: 10000,
				});

				setData({
					[item.id]: {
						...item,
						status: QueuedActionStatus.Failed,
					},
				});

				return;
			}

			async function checkResult() {
				if (cancel || !getStatus) {
					return;
				}

				const client = await getClient();
				const status = await getStatus(client);
				// TMP_KEY to key Id
				if (status.error) {
					toast({
						title: "Failed",
						description: "Action failed",
						duration: 10000,
					});
					setData({
						[item.id]: {
							...item,
							status: QueuedActionStatus.Failed,
						},
					});
				} else if (status.pending) {
					setTimeout(checkResult, 1000);
				} else if (status.done) {
					if (!status.next) {
						toast({
							title: "Success",
							description: "Action successful",
							duration: 10000,
						});
					}

					if (type === "key" && typeof status.value === "bigint") {
						const keyId = status.value;
						const settings = { ...ks?.settings, [keyId.toString()]: ks?.settings?.[TEMP_KEY] };
						delete settings[TEMP_KEY];
						setKeySettings({ settings })
					}

					setData({
						[item.id]: status.next
							? {
								...item,
								status: QueuedActionStatus.AwaitingBroadcast,
								networkType: status.next,
								value: status.value,
							}
							: {
								...item,
								status: QueuedActionStatus.Success,
							},
					});
				}
			}

			checkResult();

			return () => {
				cancel = true;
				clearTimeout(timeout);
			};
		} else if (item.status === QueuedActionStatus.AwaitingBroadcast) {
			let promise: Promise</*TransactionReceipt*/ {} | null> | undefined;

			switch (item.networkType) {
				case "eth-raw": {
					const {
						value,
						walletConnectRequestId,
						walletConnectTopic,
						snapRequestId
					} = item;

					if (!value) {
						console.error("missing value");
						return;
					}

					switch (type) {
						case "wc": {
							if (!w) {
								promise = Promise.reject(
									new Error("walletconnect not initialized"),
								);

								break;
							}

							promise = w
								.respondSessionRequest({
									topic: walletConnectTopic!,
									response: {
										jsonrpc: "2.0",
										id: walletConnectRequestId!,
										result: hexlify(value),
									},
								})
								.then(() => true);
							break;
						}

						case "snap": {
							const keyringSnapClient = new KeyringSnapRpcClient(
								env.snapOrigin,
								window.ethereum,
							);

							promise = keyringSnapClient.getRequest(snapRequestId!)
								.then((req) => keyringSnapClient.approveRequest(
									req.id,
									{ result: hexlify(value) })
								).then(() => true)
							break;
						}

						default:
							promise = Promise.reject(`type ${type} not implemented`)
					}

					break;
				}
				case "eth": {
					const {
						chainName,
						tx,
						value,
						walletConnectRequestId,
						walletConnectTopic,
						snapRequestId
					} = item;

					if (!tx || !value) {
						console.error("missing tx or value");
						return;
					}

					if (!isSupportedNetwork(chainName)) {
						console.error("unsupported network", chainName);
						return;
					}

					const signedTx = Transaction.from({ ...tx });
					signedTx.signature = hexlify(value);

					if (snapRequestId) {
						if (!signedTx.signature) {
							promise = Promise.reject(new Error("invalid signature"))
						} else {
							const keyringSnapClient = new KeyringSnapRpcClient(
								env.snapOrigin,
								window.ethereum,
							);

							const { r, s, yParity } = signedTx.signature;

							promise = keyringSnapClient.getRequest(snapRequestId).then(req => {
								return keyringSnapClient.approveRequest(req.id, {
									result: {
										r,
										s,
										v: `0x${yParity}`,

									}
								});
							}).then(() => true);
						}
					} else {
						const { provider } = getProvider(chainName);

						promise = provider
							.broadcastTransaction(signedTx.serialized)
							.then((res) => {
								if (walletConnectRequestId && walletConnectTopic) {
									if (!w) {
										throw new Error(
											"walletconnect not initialized",
										);
									}

									return (
										w
											.respondSessionRequest({
												topic: walletConnectTopic,
												response: {
													jsonrpc: "2.0",
													id: walletConnectRequestId,
													result: res.hash,
												},
											})
											// fixme
											.then(() => true)
									);
								}

								return provider.waitForTransaction(res.hash);
							});
					}

					break;
				}

				case "cosmos": {
					const {
						chainName,
						signDoc,
						value,
						pubkey,
						walletConnectRequestId,
						walletConnectTopic,
					} = item;

					if (!chainName || !signDoc || !pubkey) {
						console.error(
							"missing chainName, signDoc, value, pubkey",
						);
						return;
					}

					const chain = COSMOS_CHAINS.find(
						(item) => item.chainName === chainName,
					);

					if (!chain) {
						console.error("chain not found", chainName);
						return;
					}

					if (!isUint8Array(value)) {
						console.error("value is not Uint8Array");
						return;
					}

					let sig = value;

					if (sig.length === 65) {
						sig = sig.slice(0, -1);
					}

					if (sig.length !== 64) {
						console.error("unexpected signature length");
						return;
					}

					if (walletConnectRequestId && walletConnectTopic) {
						if (!w) {
							promise = Promise.reject(
								new Error("walletconnect not initialized"),
							);
						} else {
							promise = w
								.respondSessionRequest({
									topic: walletConnectTopic,
									response: {
										jsonrpc: "2.0",
										id: walletConnectRequestId,
										result: {
											signed: signDoc,
											signature: {
												signature: base64FromBytes(sig),
												pub_key: {
													type: "tendermint/PubKeySecp256k1",
													value: base64FromBytes(
														pubkey,
													),
												},
											},
										},
									},
								})
								// fixme
								.then(() => true);
						}
					} else {
						const { signedTxBodyBytes, signedAuthInfoBytes } =
							prepareTx(signDoc, pubkey);

						const txRaw = cosmos.tx.v1beta1.TxRaw.fromPartial({
							bodyBytes: signedTxBodyBytes,
							authInfoBytes: signedAuthInfoBytes,
							signatures: [sig],
						});

						let getRpc: Promise<string>;

						if (chain.rpc) {
							getRpc = Promise.resolve(chain.rpc[0]);
						} else {
							const repo = walletManager.getWalletRepo(chainName);
							repo.activate();
							getRpc = repo.getRpcEndpoint().then(endpoint => endpoint ? typeof endpoint === "string" ? endpoint : endpoint.url : "https://rpc.cosmos.directory/" + chainName);
						}

						promise = getRpc.then(rpc => StargateClient.connect(rpc))
							.then((client) => {
								return client.broadcastTx(
									cosmos.tx.v1beta1.TxRaw.encode(
										txRaw,
									).finish(),
								);
							})
							.then((res) => {
								if (!isDeliverTxSuccess(res)) {
									console.error("broadcast failed", res);
									throw new Error("broadcast failed");
								}

								return res as any;
							});
					}

					break;
				}

				default:
					console.error("not implemented", item.networkType);
			}

			if (!promise) {
				return;
			}

			promise
				.then((res) => {
					if (!res) {
						toast({
							title: "Failed",
							description: "Transaction failed",
							duration: 10000,
						});

						setData({
							[item.id]: {
								...item,
								status: QueuedActionStatus.Failed,
							},
						});
						return;
					}

					toast({
						title: "Success",
						description: "Transaction successful",
						duration: 10000,
					});

					setData({
						[item.id]: {
							...item,
							status: QueuedActionStatus.Success,
						},
					});
				})
				.catch((err) => {
					console.error("broadcast failed", err);

					toast({
						title: "Failed",
						description: err.message ?? "Unexpected error",
						duration: 10000,
					});

					setData({
						[item.id]: {
							...item,
							status: QueuedActionStatus.Failed,
						},
					});
				});
		}
	}, [item.status]);

	return (
		<div
			className={clsx("flex flex-col", {
				"mx-2 p-3 rounded-lg": !single,
			})}
		>
			<div className="flex items-center">
				{type === "key" ? <Assets.themeSelector className="w-8 h-5 mr-2" themeIndex={item.keyThemeIndex ?? 0} /> : null}
				<p className="text-lg font-semibold">{item.title ?? "Action"}</p>
			</div>

			<p className="text-sm text-gray-500 mb-2">
				{item.status === QueuedActionStatus.Signed
					? "Awaiting broadcast"
					: item.status === QueuedActionStatus.Broadcast
						? "Awaiting action result"
						: item.status === QueuedActionStatus.AwaitingApprovals
							? "Awaiting approvals"
							: item.status === QueuedActionStatus.ActionReady
								? "Action ready"
								: item.status ===
									QueuedActionStatus.AwaitingBroadcast
									? "Awaiting broadcast"
									: item.status === QueuedActionStatus.Success
										? "Success"
										: item.status ===
											QueuedActionStatus.Failed
											? "Failed"
											: "Unknown"}
			</p>
			{/* progress bar */}
			<div className="flex items-center gap-2">
				<div className="flex-1 h-1 bg-fill-quaternary rounded-lg">
					<div
						className="h-1 bg-accent rounded-lg"
						style={{
							width: `${item.status === QueuedActionStatus.Success
								? 100
								: (item.status /
									(QueuedActionStatus.AwaitingBroadcast +
										1)) *
								100
								}%`,
						}}
					/>
				</div>
			</div>
		</div >
	);
}

export default function ActionSidebar() {
	const { data } = useActionsState();
	const storeIds = Object.keys(data ?? {});

	const filtered = storeIds.filter((id) => {
		const action = data?.[id];
		return (
			action &&
			action.status !== QueuedActionStatus.Failed &&
			action.status !== QueuedActionStatus.Success
		);
	});

	const total = filtered.length;
	const hidden = !total;
	const first = data?.[filtered[0]];

	return (
		<div
			className={clsx(
				"flex flex-col mx-2 p-3 rounded-lg bg-fill-quaternary",
				{
					hidden,
					"border-progress": !hidden,
				},
			)}
		>
			<div>
				{total === 1 ? (
					first ? (
						<ActionItem single {...data?.[filtered[0]]!} />
					) : null
				) : (
					<Popover>
						<PopoverTrigger asChild>
							<div className="flex flex-col relative cursor-pointer">
								<p className="text-lg font-semibold">
									{total} transactions
								</p>

								<p className="text-sm text-gray-500 mb-2">
									In progress..
								</p>
							</div>
						</PopoverTrigger>
						<PopoverContent
							side="left"
							sideOffset={20}
							className="p-0"
						>
							<div className="bg-fill-quaternary">
								{filtered.map((id) => {
									const action = data?.[id];

									return action ? (
										<ActionItem key={id} {...action} />
									) : null;
								})}
							</div>
						</PopoverContent>
					</Popover>
				)}
			</div>
		</div>
	);
}
