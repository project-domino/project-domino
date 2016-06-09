/** @module autosave-manager */

/**
 * AutosaveManager is a class to set up autosaving.
 */
class AutosaveManager {
	/**
	 * @param {module:editor~Editor} parent - The instance of Editor to attach to.
	 * @param {number} [timeMax=5000] - The maximum time without typing between autosaves.
	 * @param {number} [keyMax=50] - The maximum number of keypresses between autosaves.
	 */
	constructor(parent, timeMax = 5000, keyMax = 50) {
		this.parent = parent;
		this.keyMax = keyMax;
		this.timeMax = timeMax;

		this.timer = 0;
		this.keys = 0;

		this.parent.on("keypress", () => this.keypressHandler());
	}
	keypressHandler() {
		clearTimeout(this.timer);
		if(this.keys++ > this.keyMax)
			this.save();
		else
			this.timer = setTimeout(() => this.save(), this.timeMax);
	}
	save() {
		this.keys = 0;
		this.timer = 0;

		this.parent.save();
	}
}

export default AutosaveManager;
