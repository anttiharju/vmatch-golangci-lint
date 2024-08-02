import * as vscode from "vscode";
import { workspace, window } from "vscode";
import * as path from "node:path";

export const activate = (context: vscode.ExtensionContext) => {
	const disposable = vscode.commands.registerCommand(
		"bun-vscode-extension.helloworld",
		async () => {
			let workspacePath = "";
			if (workspace.workspaceFolders?.length) {
				workspacePath = path.normalize(
					workspace.workspaceFolders[0].uri.fsPath,
				);
				const files = await workspace.findFiles("**/.golangci-version");

				// Exclude workspacePath and .golangci-version from all found files
				const relativeFiles = files.map((file) => {
					const relativePath = path.relative(workspacePath, file.fsPath);
					const suffixLength = ".golangci-version".length;
					const trimmedPath = relativePath.slice(0, -suffixLength);
					return `/${trimmedPath}`;
				});

				const message = `${workspacePath},${relativeFiles.join(",")}`;
				window.showInformationMessage(message);
			}
		},
	);

	context.subscriptions.push(disposable);
};

export const deactivate = () => {};
