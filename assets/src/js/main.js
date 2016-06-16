// js for entire website

import $ from "jquery";

$(() => {
	// Hide sidebar if screen is smaller than 900px
	if($(window).width() <= 900)
		$("body").addClass("sidebar-close");

	// Show/Hide sidebar when window is resized
	$(window).resize(() => {
		if(!$("body").hasClass("sidebar-default-close")) {
			if($(window).width() <= 900)
				$("body").addClass("sidebar-close");
			else
				$("body").removeClass("sidebar-close");
		}
	});

	// Wire up buttons
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
