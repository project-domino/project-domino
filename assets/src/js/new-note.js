import $ from "jquery";
import _ from "lodash";
import WriterPanelNoteUtil from "./util/writer-panel-note-util.js";
import getModal from "./util/modal.js";

const util = new WriterPanelNoteUtil();
const modal = getModal();

var newNoteHandler = publish => {
	return () => {
		$.ajax({
			type:     "POST",
			url:      "/api/v1/note",
			data:     JSON.stringify(_.set(util.getData(), "publish", publish)),
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
	$(".save-btn").click(newNoteHandler(false));
	$(".publish-btn").click(newNoteHandler(true));
});
