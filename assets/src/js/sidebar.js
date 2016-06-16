import $ from "jquery";

$(() => {
	// Hide sidebar if screen is smaller than 900px
	if($(window).width() <= 900)
		$("body").addClass("sidebar-close");

	// Show/Hide sidebar when window is resized
	$(window).resize(() => {
		if($(window).width() <= 900)
			$("body").addClass("sidebar-close");
		else
			$("body").removeClass("sidebar-close");
	});
});
