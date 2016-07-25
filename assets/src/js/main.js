// js for entire website

import $ from "jquery";

$(() => {
	// Wire up buttons
	var searchFunction = () => {
		window.location.assign("/search/all?q=" +
			encodeURIComponent($(".search-field").val()));
	};
	$(".search-btn").click(searchFunction);
	$(".search-field").keyup(e => {
		if(e.keyCode === 13)
			searchFunction();
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
