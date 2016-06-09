/** @module editor */

import EventEmitter    from "./ee.js";
import AutosaveManager from "./autosave-manager.js";

import {LocalStorageSaveManager} from "./save-manager.js";

/**
 * Editor is the central class used for creating an editor component.
 * @extends EventEmitter
 */
class Editor extends EventEmitter {
	constructor(e) {
		super();

		// Set up internal event handlers.
		// this.on("save-triggered", () => this.emit("save", this.document));

		// Set up external components.
		this.autosaveManager = new AutosaveManager(this);
		this.saveManager = new LocalStorageSaveManager(this);

		// Set up element.
		this.element = $(e).addClass("project-domino-editor");
		this.element.on("keydown", e => this.keyDownListener(e));
		this.element.on("keypress", e => this.keyPressListener(e));
		this.element.attr("contentEditable", true);

		// Load and render.
		return this.saveManager.load().then(note => {
			this.note = note;
			this.update();
			return this;
		}).catch(err => this.emit("error", err));
	}

	keyDownListener(event) {
		if(event.ctrlKey && event.key === "s") {
			event.stopPropagation();
			event.preventDefault();
			this.save();
		}
	}
	keyPressListener(event) {
		this.emit("keypress", event);
	}

	/**
	 * Handles all saving-related "things".
	 */
	save() {
		this.saveManager
	}
}

export default Editor;
