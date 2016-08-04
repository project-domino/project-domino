import $ from "jquery";
import _ from "lodash";
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
			modal.alert("Collection Saved", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
	$(".publish-btn").click(() => {
		$.ajax({
			type:     "PUT",
			url:      "/api/v1/collection/" + collectionJSON.ID,
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
