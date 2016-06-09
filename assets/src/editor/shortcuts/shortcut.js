import Key from "../key.js";

export default class Shortcut {
	constructor(key) {
		this.key = new Key(key);
	}
}
