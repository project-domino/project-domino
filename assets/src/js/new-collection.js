import $ from "jquery";
import _ from "lodash";
import {errorHandler} from "./util/error.js";
import WriterPanelCollectionUtil from "./util/writer-panel-collection-util.js";

const util = new WriterPanelCollectionUtil();

var newCollectionHandler = publish => {
	return () => {
		$.ajax({
			type:     "POST",
			url:      "/api/v1/collection",
			data:     JSON.stringify(_.set(util.getData(), "publish", publish)),
			dataType: "json",
		}).then(data => {
			window.location.assign("/writer-panel/collection/" + data.ID + "/edit");
		}).fail(errorHandler);
	};
};

$(() => {
	$(".save-btn").click(newCollectionHandler(false));
	$(".publish-btn").click(newCollectionHandler(true));
});
