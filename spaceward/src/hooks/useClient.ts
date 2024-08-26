import {
	cosmos,
	createRpcQueryHooks,
	getSigningWardenClient,
	useRpcClient,
	warden,
} from "@wardenprotocol/wardenjs";
import { env } from "../env";
import { useChain } from "@cosmos-kit/react";
import { EncodeObject, OfflineSigner } from "@cosmjs/proto-signing";
import { ToasterToast, useToast } from "@/components/ui/use-toast";
import { DeliverTxResponse, StdFee, isDeliverTxSuccess } from "@cosmjs/stargate";

export async function getSigningClient(signer: OfflineSigner) {
	return await getSigningWardenClient({
		signer,
		rpcEndpoint: env.rpcURL,
	});
}

const txRaw = cosmos.tx.v1beta1.TxRaw;

const defaultFee: StdFee = {
	gas: '200000',
	amount: [{ denom: 'uward', amount: '250' }],
};

export interface TxOptions {
	fee?: StdFee | null;
	toast?: Partial<ToasterToast>;
	onSuccess?: (res: DeliverTxResponse) => void;
}

export enum TxStatus {
	Failed = 'Transaction Failed',
	Successful = 'Transaction Successful',
	Broadcasting = 'Transaction confirmation in progress',
}

export function useTx() {
	const { address, getOfflineSignerDirect: getOfflineSigner } = useChain(env.cosmoskitChainName);
	const { toast } = useToast();

	const tx = async (msgs: EncodeObject[], options: TxOptions) => {
		if (!address) {
			toast({
				title: 'Wallet not connected',
				description: 'Please connect the wallet',
			});
			return;
		}

		let signed: Parameters<typeof txRaw.encode>['0'];
		const signer = getOfflineSigner();
		const client = await getSigningClient(signer);

		try {
			const fee = options.fee || defaultFee;
			signed = await client.sign(address, msgs, fee, '');
		} catch (e: unknown) {
			console.error(e);
			toast({
				title: TxStatus.Failed,
				// eslint-disable-next-line @typescript-eslint/no-explicit-any
				description: (e as any)?.message || 'An unexpected error has occured',
			});
			return;
		}

		const { id, update } = toast({
			title: TxStatus.Broadcasting,
			description: 'Waiting for transaction to be included in the block',
			duration: 999999,
		});

		if (client && signed) {
			try {
				const res = await client.broadcastTx(Uint8Array.from(txRaw.encode(signed).finish()));
				if (isDeliverTxSuccess(res)) {
					if (options.onSuccess) options.onSuccess(res);

					update({
						id,
						title: options.toast?.title || TxStatus.Successful,
						description: options.toast?.description,
					});

					return res;
				} else {
					update({
						id,
						title: TxStatus.Failed,
						description: res?.rawLog ?? 'An unexpected error has occured',
						duration: 10000,
					});

					return res;
				}
			} catch (err) {
				update({
					id,
					title: TxStatus.Failed,
					// eslint-disable-next-line @typescript-eslint/ban-ts-comment
					// @ts-ignore
					description: err?.message,
					duration: 10000,
				});
			}
		} else {
			update({
				id,
				title: TxStatus.Failed,
				description: "The transaction could't be signed or broadcasted",
				duration: 10000,
			});
		}
	};

	return { tx };
}

export function useQueryHooks() {
	const rpcQuery = useRpcClient({ rpcEndpoint: env.rpcURL });
	const rpc = rpcQuery.data;
	const isReady = !!rpcQuery.data;

	return {
		isReady,
		...createRpcQueryHooks({ rpc }),
	};
}

export function getClient() {
	return warden.ClientFactory.createRPCQueryClient({
		rpcEndpoint: env.rpcURL,
	});
}
