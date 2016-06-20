import $ from "jquery";
import _ from "lodash";
import getModal from "../js/modal.js";

const modal = getModal();

var tagResultHandler = function (data) {
	console.log(data);
};

$(() => {
	$(".new-tag-name-field").on("keyup", _.debounce(() => {
		$.ajax({
			type: "GET",
			url:  "/search/tag",
			data: {
				q: $(".new-tag-name-field").val(),
			},
			dataType: "text",
		}).then(tagResultHandler).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	}, 250));
});
