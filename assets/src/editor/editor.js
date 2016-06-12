/** @module editor */

import $ from "jquery";

import EventEmitter    from "./ee.js";
import AutosaveManager from "./autosave-manager.js";

import {LocalStorageSaveManager} from "./save-manager.js";

/**
 * Editor is the central class used for creating an editor component.
 * @extends EventEmitter
 */
class Editor extends EventEmitter {
	constructor(e, saveManager = new LocalStorageSaveManager()) {
		super();

		// Set up internal event handlers.
		this.on("command", name => this.commandHandler(name));
		this.on("formatting", name => this.formattingHandler(name));
		this.on("error", error => this.errorHandler(error));

		// Set up external components.
		this.autosaveManager = new AutosaveManager(this);
		this.saveManager = saveManager;

		// Set up element.
		this.container = $(e).addClass("project-domino-editor");
		this.container.on("keydown", e => this.keyDownListener(e));
		this.container.on("keypress", e => this.keyPressListener(e));
		this.buttons = $("<div>").addClass("project-domino-editor-buttons").append([
			["formatting", "h1"],
			["formatting", "h2"],
			["formatting", "h3"],
			["formatting", "bold"],
			["formatting", "italic"],
			["formatting", "underline"],
			["command", "save"],
		].map(label => {
			let button = $("<button>").addClass("btn btn-default").click(() => {
				this.emit(label[0], label[1]);
			}).addClass(`project-domino-editor-${label[1]}`);
			return button;
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

	commandHandler(command) {
		switch(command) {
		case "save":
			this.save();
			break;
		default:
			this.emit("error", `Unknown command: ${command}`);
			break;
		}
	}
	formattingHandler(formatting) {
		switch(formatting) {
		case "h1":
		case "h2":
		case "h3":
		case "bold":
		case "italic":
		case "underline":
		default:
			this.emit("error", `Unknown command: ${formatting}`);
			break;
		}
	}
	errorHandler(error) {
		console.error(error);
		$("<div>").addClass("project-domino-editor-error").text(error).appendTo(this.errors);
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
		console.log("Saving...", this.saveManager);
		const button = $(".project-domino-editor-save", this.buttons);
		button.attr("disabled", true);
		this.saveManager.save(this.note).then(() => {
			button.attr("disabled", false);
			console.log("Saved!");
		}).catch(err => this.emit("error", err));
	}
}

export default Editor;
