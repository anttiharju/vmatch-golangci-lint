import * as vscode from "vscode";
import { workspace, window } from "vscode";
import * as path from "node:path";

export const activate = (context: vscode.ExtensionContext) => {
	const disposable = vscode.commands.registerCommand(
		"bun-vscode-extension.helloworld",
		() => {
			let workspacePath = "";
			if (workspace.workspaceFolders?.length) {
				workspacePath = workspace.workspaceFolders[0].uri.fsPath;
				workspacePath = path.normalize(workspacePath);
				window.showInformationMessage(workspacePath);
			}
		},
	);

	context.subscriptions.push(disposable);
};

export const deactivate = () => {};
