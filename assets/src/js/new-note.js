import $ from "jquery";
import WriterPanelUtil from "./writer-panel-util.js";
import getModal from "./modal.js";

const util = new WriterPanelUtil();
const modal = getModal();

$(() => {
	util.initQuill();
	util.initTagSelector();
	$(".save-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/api/v1/note",
			data: JSON.stringify({
				title: $(".new-note-title").val(),
				body:  window.quill.getHTML(),
				tags:  $(".tag-selector").val().map(e => {return parseFloat(e);}),
			}),
			dataType: "json",
		}).then(data => {
			console.log(data);
			modal.alert("Note Saved", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
