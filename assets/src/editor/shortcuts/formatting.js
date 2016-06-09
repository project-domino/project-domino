import Shortcut from "./shortcut.js";

export class FormattingShortcut extends Shortcut {
	constructor(key, format, propagate = false) {
		super(key);
		this.format = format;
		this.propagate = propagate;
	}
	action(editor) {
		editor.emit("formatting", this.format, window.getSelection());
	}
}

export const HeaderOne   = new FormattingShortcut("C-1", "header-one");
export const HeaderTwo   = new FormattingShortcut("C-2", "header-two");
export const HeaderThree = new FormattingShortcut("C-3", "header-three");
export const HeaderFour  = new FormattingShortcut("C-4", "header-four");
export const HeaderFive  = new FormattingShortcut("C-5", "header-five");
export const HeaderSix   = new FormattingShortcut("C-6", "header-six");
export const Clear = new FormattingShortcut("Enter", "clear", true);
