/**
 * This module contains utilities for debugging.
 * @module debug
 */

import SaveManager from "./save-manager.js";

// eslint-disable-next-line no-unused-vars
const slowPromise = delay => value => new Promise((resolve, reject) => {
	window.setTimeout(() => resolve(value), delay);
});

/**
 * A SlowSaveManager wraps another SaveManager and adds a configurable delay for
 * all operations.
 */
class SlowSaveManager extends SaveManager {
	constructor(inner, delay = 1000) {
		super();
		this.inner = inner;
		this.delay = delay;
	}

	/**
	 * Handles loading.
	 * @return {Promise.<module:note~Note>} A Promise for a Note.
	 */
	load() {
		return this.inner.load()
			.then(slowPromise(this.delay));
	}

	/**
	 * Handles saving.
	 * @param {module:note~Note} note - The Note to save.
	 * @return {Promise} A Promise for the save operation's completion.
	 */
	save(note) {
		return Promise.resolve(note)
			.then(slowPromise(this.delay))
			.then(note => this.inner.save(note));
	}
}

export {
	slowPromise,
	SlowSaveManager,
};
