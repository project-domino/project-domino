import $ from "jquery";

import Editor from "./editor.js";
import {LocalStorageSaveManager} from "./save-manager.js";
import {SlowSaveManager} from "./debug.js";
import {
	Note,
	HeaderNode,
	ParagraphNode,
	TextNode,
} from "./note.js";

const saveManager = new SlowSaveManager(new LocalStorageSaveManager());
const editor = new Editor(document.getElementById("editor"), saveManager);
window.editor = editor;
export default editor;

let note = new Note([
	new HeaderNode("Header", 1),
	new ParagraphNode([
		new TextNode("Hello, world!"),
	]),
]);
window.note = note;

$("#editor > .project-domino-editor-content").append(note.render());
