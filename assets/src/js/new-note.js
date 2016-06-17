import $ from "jquery";
import "select2";
import Editor from "../editor/editor.js";

$(() => {
	const editor = new Editor(document.querySelector(".new-note-editor"));
	window.editor = editor;
	$(".tag-selector").select2({
		placeholder: "Type to search for tags...",
		allowClear:  true,
	});
});
