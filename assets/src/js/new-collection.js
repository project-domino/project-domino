import $ from "jquery";
import _ from "lodash";
import WriterPanelCollectionUtil from "./writer-panel-collection-util.js";
import getModal from "./modal.js";

const util = new WriterPanelCollectionUtil();
const modal = getModal();

var newCollectionHandler = publish => {
	return () => {
		$.ajax({
			type:     "POST",
			url:      "/api/v1/collection",
			data:     JSON.stringify(_.set(util.getData(), "publish", publish)),
			dataType: "json",
		}).then(data => {
			window.location.assign("/writer-panel/collection/" + data.ID + "/edit");
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	};
};

$(() => {
	$(".save-btn").click(newCollectionHandler(false));
	$(".publish-btn").click(newCollectionHandler(true));
});
