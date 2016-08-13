import $ from "jquery";

import getModal from "./util/modal.js";

const modal = getModal();

$(() => {
	$(".verify-btn").click(() => {
		$.ajax({
			type:     "POST",
			url:      "/email/verify",
			dataType: "text",
		}).then(() => {
			modal.alert("Sent Verification Email", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
