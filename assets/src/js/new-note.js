import $ from "jquery";
import WriterPanelUtil from "./writer-panel-util.js";
import getModal from "./modal.js";

const util = new WriterPanelUtil();
const modal = getModal();

var newNoteHandler = request => {
	return () => {
		$.ajax({
			type:     "POST",
			url:      "/api/v1/note",
			data:     JSON.stringify(request()),
			dataType: "json",
		}).then(data => {
			window.location.assign("/writer-panel/note/" + data.ID + "/edit");
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	};
};

$(() => {
	util.initQuill();
	util.initTagSelector();
	$(".save-btn").click(newNoteHandler(() => {
		return {
			title: $(".new-note-title").val(),
			body:  window.quill.getHTML(),
			tags:  $(".tag-selector").val().map(e => {return parseFloat(e);}),
		};
	}));
	$(".publish-btn").click(newNoteHandler(() => {
		return {
			title:   $(".new-note-title").val(),
			body:    window.quill.getHTML(),
			tags:    $(".tag-selector").val().map(e => {return parseFloat(e);}),
			publish: true,
		};
	}));
});
