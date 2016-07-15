import $ from "jquery";
import _ from "lodash";
import WriterPanelNoteUtil from "./writer-panel-note-util.js";
import getModal from "./modal.js";

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
			modal.alert("Note Saved", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
	$(".publish-btn").click(() => {
		$.ajax({
			type:     "PUT",
			url:      "/api/v1/note/" + note.ID,
			data:     JSON.stringify(_.set(util.getData(), "publish", true)),
			dataType: "json",
		}).then(() => {
			window.location.reload();
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
