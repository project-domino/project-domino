import $ from "jquery";
import _ from "lodash";
import {errorHandler} from "./util/error.js";
import WriterPanelCollectionUtil from "./util/writer-panel-collection-util.js";
import getModal from "./util/modal.js";

const modal = getModal();

$(() => {
	// Parse collection JSON and initiate page
	const collectionJSON = JSON.parse($("#collection-data").text());
	const util = new WriterPanelCollectionUtil(collectionJSON.Notes);

	if(collectionJSON.Tags) {
		$(".tag-selector").val(collectionJSON.Tags.map(e => {
			return e.ID;
		})).trigger("change");
	}

	$(".save-btn").click(() => {
		$.ajax({
			type:     "PUT",
			url:      "/api/v1/collection/" + collectionJSON.ID,
			data:     JSON.stringify(_.set(util.getData(), "publish", collectionJSON.Published)),
			dataType: "json",
		}).then(() => {
			modal.alert("Collection Saved");
		}).fail(errorHandler);
	});
	$(".publish-btn").click(() => {
		$.ajax({
			type:     "PUT",
			url:      "/api/v1/collection/" + collectionJSON.ID,
			data:     JSON.stringify(_.set(util.getData(), "publish", true)),
			dataType: "json",
		}).then(() => {
			window.location.reload();
		}).fail(errorHandler);
	});
});
