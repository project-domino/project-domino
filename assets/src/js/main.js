// js for entire website

import $ from "jquery";
import {errorHandler} from "./util/error.js";

$(() => {
	$(".logout-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/logout",
		}).then(() => {
			window.location.reload();
		}).fail(errorHandler);
	});

	// Setup dropdowns
	["account", "notification", "mobile"].forEach(e => {
		$(`.${e}-dropdown-btn`).click(() => {
			$(".dropdown-content").not(`.${e}-dropdown-content`).hide();
			$(`.${e}-dropdown-content`).toggle();
		});
	});
	$(window).resize(() => {
		$(".dropdown-content").hide();
	});

	// Notifications
	if(JSON.parse($("#logged-in-val").text())) {
		$.ajax({
			type:     "GET",
			url:      "/api/v1/notifications",
			dataType: "json",
		}).then(data => {
			$(".unread-notifications-list").empty();
			if(data.length <= 0) {
				$(".unread-notifications-list").append(
					$("<span>").addClass("dropdown-item no-items")
					.text("You have no notifications.")
				);
				return;
			}
			$(".unread-notifications-list").append(
				data.map(e => {
					return $("<div>")
						.addClass(e.Read ? "inactive" : "active")
						.addClass("dropdown-item")
						.append(
							$("<a>")
								.addClass("dropdown-item-title")
								.addClass("notification-title")
								.attr({
									"href":                 e.Link,
									"data-notification-id": e.ID,
								})
								.text(e.Title)
								.click(function (e) {
									e.preventDefault();
									var notificationPage = () => {
										window.location.assign($(this).attr("href"));
									};
									$.ajax({
										type: "PUT",
										url:  "/api/v1/notification/" +
											$(this).data("notification-id") +
											"/read",
									}).then(notificationPage);
								}),
							$("<span>").addClass("dropdown-item-date")
								.text(moment(e.CreatedAt).from(moment())), // eslint-disable-line no-undef
							$("<p>").addClass("dropdown-item-description").text(e.Message)
						);
				})
			);
		});
	}
});
