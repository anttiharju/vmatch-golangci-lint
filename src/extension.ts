import * as vscode from "vscode";
import { workspace, window } from "vscode";
import * as path from "node:path";

export const activate = (context: vscode.ExtensionContext) => {
	const onSaveDisposable = vscode.workspace.onDidSaveTextDocument(
		async (document) => {
			let workspacePath = "";
			if (workspace.workspaceFolders?.length) {
				workspacePath = path.normalize(
					workspace.workspaceFolders[0].uri.fsPath,
				);
				const files = await workspace.findFiles("**/.golangci-version");

				let message = `${workspacePath}`;

				for (const file of files) {
					const fileContent = await vscode.workspace.fs.readFile(file);
					const suffixLength = "/.golangci-version".length;
					const trimmedPath = file.fsPath.slice(0, -suffixLength);
					message += `\n${trimmedPath}`;
					message += `\n${fileContent.toString()}`;
				}

				window.showInformationMessage(message);
			}
		},
	);

	context.subscriptions.push(onSaveDisposable);
};

export const deactivate = () => {};
