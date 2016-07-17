// js for entire website

import $ from "jquery";

$(() => {
	// Wire up buttons
	$(".search-btn").click(() => {
		$.ajax({
			type: "GET",
			url:  "/api/v1/search",
			data: {
				q: $(".search-field").val(),
			},
			dataType: "json",
		}).fail(err => {
			console.log(err);
		}).then(data => {
			console.log(data);
		});
	});
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
	$(".sign-in-btn").click(() => {
		window.location.assign("/login");
	});
	$(".account-dropdown-btn").click(() => {
		$(".account-dropdown-content").toggle();
	});
	$(".new-item-dropdown-btn").click(() => {
		$(".new-item-dropdown-content").toggle();
	});
	$(".sidebar-btn").click(() => {
		$("body").toggleClass("sidebar-close");
	});
});
