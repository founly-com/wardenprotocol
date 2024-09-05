import clsx from "clsx";
import AddOwnerModal from "./AddOwner";
import ApproveModal from "./ApproveActionModal";
import ApproveSnapModal from "./ApproveSnapModal";
import CreateKeyModal from "./CreateKey";
import ConnectedModal from "./ConnectedModal";
import ReceiveAssetsModal from "./ReceiveAssets";
import SelectAssetModal from "./SelectAsset";
import SendAssetsModal from "./SendAssets";
import WalletConnectModal from "./WalletConnect";
import { useModalState } from "./state";
import type {
	CreateKeyParams,
	SelectAssetParams,
	TransferParams,
} from "./types";
import { useMediaQuery } from "@uidotdev/usehooks";
import { useEffect } from "react";

export default function ModalRoot() {
	const { data, setData } = useModalState();
	const isDesktop = useMediaQuery("(min-width: 768px)");

	useEffect(() => {
		const handleKeyDown = (e: KeyboardEvent) => {
			if (e.key === "Escape") {
				setData({ type: undefined, params: undefined });
			}
		};

		window.addEventListener("keydown", handleKeyDown);
		return () => window.removeEventListener("keydown", handleKeyDown);
	}, []);

	if (!data) {
		return null;
	}

	return (
		<div
			className={clsx(
				"absolute left-0 top-0 w-full h-full flex flex-col items-center justify-center min-h-[600px] isolate z-[999]",
				{ hidden: !data.type },
			)}
		>
			<div className="bg-overlay fixed left-0 top-0 w-full h-full backdrop-blur-[20px] -z-10"></div>
			{!isDesktop ? (
				<button
					onClick={() => setData({ type: undefined, params: undefined })}
					className="absolute top-8 invert dark:invert-0 right-8 opacity-[0.5] hover:opacity-[100%] transition-all"
				>
					<img src="/images/button-close.svg" alt="" />
				</button>
			) : null}

			<div className="flex flex-col relative">
				{isDesktop ? (
					<button
						onClick={() => setData({ type: undefined, params: undefined })}
						className="absolute -top-16 invert dark:invert-0 -right-16 opacity-[0.5] hover:opacity-[100%] transition-all"
					>
						<img src="/images/button-close.svg" alt="" />
					</button>
				) : null}

				{data.type === "approve-action" ||
					data.background["approve-action"] ? (
					<ApproveModal hidden={data.type !== "approve-action"} />
				) : null}
				{data.type === "approve-snap" ? (
					<ApproveSnapModal hidden={data.type !== "approve-snap"} />
				) : null}
				{data.type === "add-owner" || data.background["add-owner"] ? (
					<AddOwnerModal hidden={data.type !== "add-owner"} />
				) : null}
				{data.type === "create-key" || data.background["create-key"] ? (
					<CreateKeyModal
						hidden={data.type !== "create-key"}
						{...(data.background["create-key"] ??
							(data.params as CreateKeyParams))}
					/>
				) : null}
				{data.type === "walletconnect" ||
					data.background["walletconnect"] ? (
					<WalletConnectModal hidden={data.type !== "walletconnect"} />
				) : null}
				{data.type === "dapps-modal" ? (
					<ConnectedModal />
				) : !data.params ? (
					<></>
				) : data.type === "receive" ? (
					<ReceiveAssetsModal {...(data.params as TransferParams)} />
				) : data.type === "send" ? (
					<SendAssetsModal {...(data.params as TransferParams)} />
				) : data.type === "select-asset" ? (
					<SelectAssetModal {...(data.params as SelectAssetParams)} />
				) : (
					<></>
				)}
			</div>
		</div>
	);
}
