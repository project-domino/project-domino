import Shortcut from "./shortcut.js";

export class SaveShortcut extends Shortcut {
	action(editor) {
		editor.emit("save-triggered");
	}
}

export const Save = new SaveShortcut("C-s");
