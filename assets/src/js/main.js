// js for entire website

import $ from "jquery";

$(() => {
	$(".logout-btn").click(() => {
		$.ajax({
			type:     "POST",
			url:      "/logout",
			dataType: "text",
		}).fail((err) => {
			console.log(err);
		}).then(() => {
			window.location.reload();
		});
	});
});
