// js for entire website

import $ from "jquery";

$(() => {
	// Wire up buttons
	$(".logout-btn").click(() => {
		$.ajax({
			type:     "POST",
			url:      "/logout",
			dataType: "text",
		}).fail(err => {
			console.log(err);
		}).then(() => {
			window.location.reload();
		});
	});
	$(".account-dropdown-btn").click(() => {
		$(".account-dropdown-content").toggle();
	});
	$(".notification-dropdown-btn").click(() => {
		$(".notification-dropdown-content").toggle();
	});
	$(".sidebar-btn").click(() => {
		$("body").toggleClass("sidebar-close");
	});
});
