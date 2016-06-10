/**
 * This module contains utilities for debugging.
 * @module debug
 */

import SaveManager from "./save-manager.js";

/**
 * Returns a thenable function that returns a promise that allows a value to
 * pass through after waiting for the given delay.
 * @param {number} delay - The delay to add, in milliseconds.
 * @return {module:debug~thenable}
 */
// eslint-disable-next-line no-unused-vars
const slowPromise = delay => value => new Promise((resolve, reject) => {
	window.setTimeout(() => resolve(value), delay);
});

/**
 * A function that accepts zero or one parameters and returns a Promise.
 * @callback thenable
 * @param {*} [previous] - The result from the previous Promise, if any.
 * @return {Promise}
 */

/**
 * A SlowSaveManager wraps another SaveManager and adds a configurable delay for
 * all operations.
 * @extends module:save-manager~SaveManager
 */
class SlowSaveManager extends SaveManager {
	/**
	 * Constructs a SlowSaveManager.
	 * @param {module:save-manager~SaveManager} inner - The SaveManager to wrap.
	 * @param {number} delay - The delay to add, in milliseconds.
	 */
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
