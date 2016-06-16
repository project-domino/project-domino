/** @module editor */

import $ from "jquery";
import StackTrace from "stacktrace-js";

import getModal from "../js/modal.js";

import AutosaveManager from "./autosave-manager.js";
import EventEmitter    from "./ee.js";

import {LocalStorageSaveManager} from "./save-manager.js";

const modal = getModal();

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
		this.on("render", () => this.renderHandler());

		// Set up external components.
		this.autosaveManager = new AutosaveManager(this);
		this.saveManager = saveManager;

		// Set up element.
		this.container = $(e).addClass("project-domino-editor");
		this.container.on("keydown", e => this.keyDownListener(e));
		this.container.on("keypress", e => this.keyPressListener(e));
		this.buttons = $("<div>").addClass("project-domino-editor-buttons").append([
			["formatting", "h1", "header", "1"],
			["formatting", "h2", "header", "2"],
			["formatting", "h3", "header", "3"],
			["formatting", "bold"],
			["formatting", "italic"],
			["formatting", "underline"],
			["command", "save"],
			["command", "debug-export", "terminal", " export"],
		].map(label => {
			const eventType = label[0];
			const eventName = label[1];
			const labelIcon = label[2] || eventName;
			const labelText = label[3];

			return $("<button>").attr("disabled", true).click(() => {
				this.emit(eventType, eventName);
			}).addClass(`project-domino-editor-${eventName}`).append([
				$("<span>").addClass(`fa fa-${labelIcon}`),
				$("<span>").text(labelText),
			]);
		})).appendTo(this.container);
		this.element = $("<div>").addClass("project-domino-editor-content").appendTo(this.container);

		// Load and render.
		this.saveManager.load().then(note => {
			this.note = note;
			this.emit("render");
			this.element.attr("contentEditable", true);
			$("button", this.container).attr("disabled", false);
		}).catch(err => this.emit("error", err));
	}

	commandHandler(command) {
		switch(command) {
		case "debug-export":
			$("#editor-debug-out").append($("<pre>").append(
				$("<code>").text(this.note.toJSON())
			));
			break;
		case "save":
			this.save();
			break;
		default:
			this.emit("error", new Error(`Unknown command: ${command}`));
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
			this.emit("error", new Error(`Unknown command: ${formatting}`));
			break;
		}
	}
	errorHandler(error) {
		modal.alert(error, -1);
		console.error(error);
		StackTrace.fromError(error).then(trace => {
			console.log(trace.join("\n"));
		}).catch(err => console.error(err));
	}
	renderHandler() {
		this.element.empty().append(this.note.render());
	}

	keyDownListener(event) {
		if(event.ctrlKey && event.key === "s") {
			event.stopPropagation();
			event.preventDefault();
			this.emit("command", "save");
		}
	}
	keyPressListener(event) {
		this.emit("keypress", event);
		setTimeout(() => {
			// Wait for the DOM to be updated.
			this.note.update(this.element);
		}, 0);
	}

	/**
	 * Handles all saving-related "things".
	 */
	save() {
		console.log("Saving...");
		const button = $(".project-domino-editor-save", this.buttons);
		button.attr("disabled", true);
		this.saveManager.save(this.note).then(() => {
			button.attr("disabled", false);
			console.log("Saved!");
		}).catch(err => this.emit("error", err));
	}
}

export default Editor;
