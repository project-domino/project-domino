import $ from "jquery";
import WriterPanelUtil from "./writer-panel-util.js";
import getModal from "./modal.js";

const util = new WriterPanelUtil();
const modal = getModal();

$(() => {
	// Parse note JSON
	var noteJSON = JSON.parse($("#note-data").text());

	// Set up quill and tag selector
	util.initQuill();
	util.initTagSelector();
	if(noteJSON.Tags) {
		$(".tag-selector").val(noteJSON.Tags.map(e => {
			return e.ID;
		})).trigger("change");
	}
	window.quill.setHTML(noteJSON.Body);

	// Wire up buttons
	$(".save-btn").click(() => {
		$.ajax({
			type: "PUT",
			url:  "/api/v1/note/" + noteJSON.ID,
			data: JSON.stringify({
				title:   $(".new-note-title").val(),
				body:    window.quill.getHTML(),
				tags:    $(".tag-selector").val().map(e => {return parseFloat(e);}),
				publish: noteJSON.Published,
			}),
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
			type: "PUT",
			url:  "/api/v1/note/" + noteJSON.ID,
			data: JSON.stringify({
				title:   $(".new-note-title").val(),
				body:    window.quill.getHTML(),
				tags:    $(".tag-selector").val().map(e => {return parseFloat(e);}),
				publish: true,
			}),
			dataType: "json",
		}).then(() => {
			window.location.reload();
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
