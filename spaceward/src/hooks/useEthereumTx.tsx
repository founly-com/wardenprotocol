import { ethers } from "ethers";
import { env } from "@/env";
import { useNewAction } from "./useAction";
import { warden } from "@wardenprotocol/wardenjs";
import { useEnqueueAction } from "@/features/actions/hooks";

export function useEthereumTx() {
	const { getMessage, authority } = useNewAction(
		warden.warden.v1beta3.MsgNewSignRequest,
	);

	const { addAction } = useEnqueueAction(getMessage);

	const signRaw = async (
		keyId: bigint,
		input: Uint8Array,

		wc?: {
			requestId: number;
			topic: string;
		},
		snap?: {
			requestId: string;
		}
	) => {
		if (!authority) {
			throw new Error("no authority");
		}

		return await addAction(
			{
				authority,
				keyId,
				analyzers: [],
				input,
				// @ts-expect-error telescope generated code doesn't handle empty array correctly, use `undefined` instead of `[]`
				encryptionKey: undefined,
			},
			{
				walletConnectRequestId: wc?.requestId,
				walletConnectTopic: wc?.topic,
				snapRequestId: snap?.requestId
			},
		);
	};

	const signEthereumTx = async (
		keyId: bigint,
		_tx: ethers.TransactionLike,
		chainName: string,
		wc?: {
			requestId: number;
			topic: string;
		},
		snap?: {
			requestId: string;
		}
	) => {
		if (!authority) {
			throw new Error("no authority");
		}

		if (!env.ethereumAnalyzerContract) {
			console.warn(
				"Missing ethereumAnalyzerContract. Can't use Ethereum transactions.",
			);

			return;
		}

		const tx = ethers.Transaction.from(_tx);

		return await addAction(
			{
				analyzers: [env.ethereumAnalyzerContract],
				authority,
				input: ethers.getBytes(tx.unsignedSerialized),
				keyId,
				// @ts-expect-error telescope generated code doesn't handle empty array correctly, use `undefined` instead of `[]`
				encryptionKey: undefined,
			},
			{
				tx: _tx,
				chainName,
				walletConnectRequestId: wc?.requestId,
				walletConnectTopic: wc?.topic,
				snapRequestId: snap?.requestId
			},
		);
	};

	return {
		signRaw,
		signEthereumTx,
	};
}
