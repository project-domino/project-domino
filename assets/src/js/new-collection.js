import $ from "jquery";
import _ from "lodash";
import getModal from "../js/modal.js";
import WriterPanelUtil from "./writer-panel-util.js";

const modal = getModal();

const util = new WriterPanelUtil();

var searchHandler = () => {
	$.ajax({
		type: "GET",
		url:  "/api/v1/search/note",
		data: {
			q: $(".note-search-field").val(),
		},
		dataType: "json",
	}).then(data => {
		$(".search-result-container").append(
			data.map(e => {
				return $("<div>").append(
					$("<h5>").text(e.Title)
				).addClass("list-item");
			})
		);
	}).fail(err => {
		console.log(err);
		modal.alert(err.responseText, 3000);
	});
};

$(() => {
	util.initTagSelector();
	$(".note-search-field").on("keyup", _.debounce(searchHandler, 250));
	$(".note-search-btn").click(searchHandler);
});
