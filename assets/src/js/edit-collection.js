import $ from "jquery";
import _ from "lodash";
import WriterPanelCollectionUtil from "./writer-panel-collection-util.js";
import getModal from "./modal.js";

const modal = getModal();

$(() => {
	// Parse collection JSON and initiate page
	const collectionJSON = JSON.parse($("#collection-data").text());
	// TODO get notes to actually show up
	const util = new WriterPanelCollectionUtil(collectionJSON.Notes);

	if(collectionJSON.Tags) {
		$(".tag-selector").val(collectionJSON.Tags.map(e => {
			return e.ID;
		})).trigger("change");
	}

	// Wire up button handlers
	var editCollectionHandler = (publish, callback) => {
		return () => {
			$.ajax({
				type:     "PUT",
				url:      "/api/v1/collection/" + collectionJSON.ID,
				data:     JSON.stringify(_.set(util.getData(), "publish", publish)),
				dataType: "json",
			}).then(callback).fail(err => {
				console.log(err);
				modal.alert(err.responseText, 3000);
			});
		};
	};
	$(".save-btn").click(editCollectionHandler(collectionJSON.Published), () => {
		modal.alert("Collection Saved", 3000);
	});
	$(".publish-btn").click(editCollectionHandler(true), () => {
		window.location.reload();
	});
});
