import $ from "jquery";
import _ from "lodash";
import {errorHandler} from "./util/error.js";
import getModal from "./util/modal.js";

const modal = getModal();

// Changes resultNotification and resultTable based on results from tag search query
var tagResultHandler = json => {
	var resultNotification = $(".result-notification");
	var resultTable = $(".tag-search-table");
	if($(".new-tag-name-field").val() === "") {
		resultTable.hide();
		resultNotification.show();
		resultNotification.text("Please enter your tag name...");
	} else if(json.length === 0) {
		resultTable.hide();
		resultNotification.show();
		resultNotification.text("No similar tags found...");
	} else {
		resultTable.show();
		resultNotification.hide();
		$(".result-table-element").remove();
		resultTable.append(json.map(e => {
			return $("<tr>").append(
				$("<td>").text(e.Name),
				$("<td>").text(e.Description)
			).addClass("result-table-element").data("id", e.ID);
		}));
	}
};

// Toggles the description field in the new tag panel
var toggleCreateHandler = () => {
	var btnIcon = $(".toggle-create-btn-icon");
	var panel = $(".new-note-container");
	if(panel.hasClass("create-closed")) {
		panel.removeClass("create-closed").addClass("create-open");
		btnIcon.removeClass("fa-plus").addClass("fa-minus");
		$(".toggle-content").removeClass("hidden");
	} else {
		panel.removeClass("create-open").addClass("create-closed");
		btnIcon.removeClass("fa-minus").addClass("fa-plus");
		$(".toggle-content").addClass("hidden");
	}
};

$(() => {
	$(".toggle-create-btn").click(toggleCreateHandler);

	$(".new-tag-name-field").on("keyup", _.debounce(() => {
		$.ajax({
			type: "GET",
			url:  "/api/v1/search/tag",
			data: {
				q: $(".new-tag-name-field").val(),
			},
			dataType: "json",
		}).then(tagResultHandler).fail(errorHandler);
	}, 100));

	$(".new-tag-create-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/api/v1/tag",
			data: {
				name:        $(".new-tag-name-field").val(),
				description: $(".new-tag-description-field").val(),
			},
			dataType: "json",
		}).then(() => {
			$(".new-tag-name-field").val("");
			$(".new-tag-description-field").val("");
			modal.alert("Tag Created");
		}).fail(errorHandler);
	});
});
