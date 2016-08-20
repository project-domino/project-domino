import $ from "jquery";
import _ from "lodash";
import {errorHandler} from "./util/error.js";
import WriterPanelNoteUtil from "./util/writer-panel-note-util.js";
import getModal from "./util/modal.js";

const modal = getModal();

$(() => {
	// Parse noteJSON
	var note = JSON.parse($("#note-data").text());

	// Initialize page
	const util = new WriterPanelNoteUtil(note);

	// Wire up buttons
	$(".save-btn").click(() => {
		$.ajax({
			type:     "PUT",
			url:      "/api/v1/note/" + note.ID,
			data:     JSON.stringify(_.set(util.getData(), "publish", note.Published)),
			dataType: "json",
		}).then(() => {
			modal.alert("Note Saved");
		}).fail(errorHandler);
	});
	$(".publish-btn").click(() => {
		$.ajax({
			type:     "PUT",
			url:      "/api/v1/note/" + note.ID,
			data:     JSON.stringify(_.set(util.getData(), "publish", true)),
			dataType: "json",
		}).then(() => {
			window.location.reload();
		}).fail(errorHandler);
	});
});
