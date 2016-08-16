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

	// Notifications
	$.ajax({
		type:     "GET",
		url:      "/api/v1/notifications",
		dataType: "json",
	}).then(data => {
		$(".unread-notifications-list").empty();
		if(data.length > 0) {
			$(".unread-notifications-list").append(
				data.map(e => {
					return $("<div>").addClass("dropdown-item").append(
						$("<a>").addClass("dropdown-item-title").attr({
							"href":                 e.Link,
							"data-notification-id": e.ID,
						}).text(e.Title),
						$("<span>").addClass("dropdown-item-date").text(e.CreatedAt),
						$("<p>").addClass("dropdown-item-description").text(e.Message)
					);
				})
			);
		} else {
			$(".unread-notifications-list").append(
				$("<span>").addClass("dropdown-item no-items")
					.text("You have no notifications.")
			);
		}
	});
});
