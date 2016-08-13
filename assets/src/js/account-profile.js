import $ from "jquery";

import getModal from "./util/modal.js";

const modal = getModal();

$(() => {
	$(".update-btn").click(() => {
		$.ajax({
			type: "PUT",
			url:  "/account",
			data: {
				name:  $(".user-name").val().trim(),
				email: $(".user-email").val().trim(),
			},
			dataType: "text",
		}).then(() => {
			modal.alert("Profile Updated", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
