import EventEmitter    from "./ee.js";
import AutosaveManager from "./autosave-manager.js";
import {Document}      from "./document.js";
import Key             from "./key.js";
import ShortcutManager from "./shortcut-manager.js";

import LocalStorageSaveManager from "./savemanagers/localstorage.js";

export default class Editor extends EventEmitter {
	constructor(e) {
		super();

		// Set up internal event handlers.
		this.on("render", contents => this.render(contents));
		this.on("save-triggered", () => this.emit("save", this.document));

		// Set up external components.
		this.autosaveManager = new AutosaveManager(this);
		this.document = new Document(this);
		this.saveManager = new LocalStorageSaveManager(this);
		this.shortcutManager = new ShortcutManager(this);

		// Set up element.
		this.element = e;
		this.element.classList.add("project-domino-editor");
		this.element.addEventListener("keydown", e => this.keyDownListener(e));
		this.element.addEventListener("keypress", e => this.keyPressListener(e));
		this.element.contentEditable = true;

		// Load and render.
		this.emit("render", this.document.render());
	}
	keyDownListener(event) {
		const key = new Key(event);

		if(this.shortcutManager.isShortcut(key)) {
			this.emit("shortcut", key);
			if(!this.shortcutManager.getShortcut(key).propagate) {
				event.stopPropagation();
				event.preventDefault();
			}
		}
	}
	keyPressListener(event) {
		this.emit("keypress", new Key(event));
	}
	render(contents) {
		for(const node of this.element.childNodes)
			this.element.removeChild(node);
		contents.forEach(e => this.element.appendChild(e));
		console.debug(contents.map(e => e.textContent));
	}
}
