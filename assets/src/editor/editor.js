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
		this.on("button", name => this.buttonHandler(name));
		this.on("error", error => this.errorHandler(error));
		// this.on("save-triggered", () => this.emit("save", this.document));

		// Set up external components.
		this.autosaveManager = new AutosaveManager(this);
		this.saveManager = new LocalStorageSaveManager(this);

		// Set up element.
		this.container = $(e).addClass("project-domino-editor");
		this.container.on("keydown", e => this.keyDownListener(e));
		this.container.on("keypress", e => this.keyPressListener(e));
		this.buttons = $("<div>").addClass("project-domino-editor-buttons").append([
			"h1",
			"h2",
			"h3",
			"bold",
			"italic",
			"underline",
		].map(label => {
			return $("<button>").addClass("btn btn-default").text(label).click(() => {
				this.emit("button", label);
			});
		})).appendTo(this.container);
		this.element = $("<div>").addClass("project-domino-editor-content").attr({
			contentEditable: true,
		}).appendTo(this.container);
		this.errors = $("<div>").addClass("project-domino-editor-errors").appendTo(this.container);

		// Load and render.
		return this.saveManager.load().then(note => {
			this.note = note;
			this.emit("update");
			return this;
		}).catch(err => this.emit("error", err));
	}

	buttonHandler(name) {
		// TODO
		this.emit("error", `TODO buttonHandler(${name})`);
	}
	errorHandler(error) {
		alert(error);
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
		// TODO
	}
}

export default Editor;
