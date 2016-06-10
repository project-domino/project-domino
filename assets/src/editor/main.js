import Editor from "./editor.js";
import {
	Note,
	HeaderNode,
	ParagraphNode,
	TextNode,
} from "./note.js";

const editor = new Editor(document.getElementById("editor"));
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
