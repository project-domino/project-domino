/** @module save-manager */

import {Note} from "./note.js";

/**
 * SaveManager is an abstract class that handles saving and loading data.
 */
class SaveManager {
	/**
	 * Attaches a SaveManager to an EventEmitter.
	 * @deprecated
	 * @param {module:ee~EventEmitter} ee - The EventEmitter to attach to.
	 * @param {string} [saveEvent="save"] - The event to save on.
	 * @param {string} [doneEvent="saved"] - The event to emit when finished saving.
	 * @param {string} [errorEvent="error"] - The event to emit on error.
	 */
	attach(ee, saveEvent = "save", doneEvent = "saved", errorEvent = "error") {
		ee.on(saveEvent, note => {
			this.save(note).then(() => {
				ee.emit(doneEvent);
			}).catch(err => {
				ee.emit(errorEvent, err);
			});
		});
	}

	/**
	 * Handles loading.
	 * @return {Promise.<module:note~Note>} A Promise for a Note.
	 */
	load() {
		return Promise.reject(new TypeError("A SaveManager subclass must implement load()."));
	}

	/**
	 * Handles saving.
	 * @param {module:note~Note} note - The Note to save.
	 * @return {Promise} A Promise for the save operation's completion.
	 */
	save(note) { // eslint-disable-line no-unused-vars
		return Promise.reject(new TypeError("A SaveManager subclass must implement save()."));
	}
}

/**
 * A SaveManager that saves data to LocalStorage.
 * @extends module:save-manager~SaveManager
 */
class LocalStorageSaveManager extends SaveManager {
	/**
	 * @param {string} key - The key to save notes under.
	 */
	constructor(key = "pdeditor-save") {
		super();
		this.key = key;
	}

	/**
	 * Handles loading.
	 * @return {Promise} A Promise for a Note.
	 */
	load() {
		return new Promise((resolve, reject) => {
			let note;
			try {
				note = new Note(JSON.parse(localStorage.getItem(this.key)));
			} catch(err) {
				return reject(err);
			}
			resolve(note);
		});
	}

	/**
	 * Handles saving.
	 * @param {Note} note - The Note to save.
	 * @return {Promise} A Promise for the save operation's completion.
	 */
	save(note) {
		return new Promise((resolve, reject) => {
			try {
				localStorage.setItem(this.key, JSON.stringify(note));
			} catch(err) {
				return reject(err);
			}
			resolve();
		});
	}
}

export default SaveManager;
export {
	LocalStorageSaveManager,
};
