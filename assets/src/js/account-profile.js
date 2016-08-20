import $ from "jquery";
import {errorHandler} from "./util/error.js";
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
		}).then(() => {
			modal.alert("Profile Updated");
		}).fail(errorHandler);
	});
});
