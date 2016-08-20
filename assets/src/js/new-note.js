import $ from "jquery";
import _ from "lodash";
import {errorHandler} from "./util/error.js";
import WriterPanelNoteUtil from "./util/writer-panel-note-util.js";

const util = new WriterPanelNoteUtil();

var newNoteHandler = publish => {
	return () => {
		$.ajax({
			type:     "POST",
			url:      "/api/v1/note",
			data:     JSON.stringify(_.set(util.getData(), "publish", publish)),
			dataType: "json",
		}).then(data => {
			window.location.assign("/writer-panel/note/" + data.ID + "/edit");
		}).fail(errorHandler);
	};
};

$(() => {
	$(".save-btn").click(newNoteHandler(false));
	$(".publish-btn").click(newNoteHandler(true));
});
