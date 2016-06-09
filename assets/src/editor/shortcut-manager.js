import {
	HeaderOne,
	HeaderTwo,
	HeaderThree,
	HeaderFour,
	HeaderFive,
	HeaderSix,
	Clear,
}             from "./shortcuts/formatting.js";
import {Save} from "./shortcuts/save.js";

export default class ShortcutManager {
	constructor(parent) {
		this.parent = parent;

		this.shortcuts = [
			HeaderOne,
			HeaderTwo,
			HeaderThree,
			HeaderFour,
			HeaderFive,
			HeaderSix,
			Clear,
			Save,
		];
		this.parent.on("shortcut", key => this.shortcutHandler(key));
	}
	getShortcut(key) {
		const shortcuts = this.shortcuts.filter(shortcut => shortcut.key.equals(key));
		if(shortcuts.length === 0)
			return null;
		return shortcuts[0];
	}
	isShortcut(key) {
		return this.getShortcut(key) !== null;
	}
	shortcutHandler(key) {
		this.shortcuts
			.filter(shortcut => shortcut.key.equals(key))
			.forEach(shortcut => shortcut.action(this.parent));
	}
}
