import $ from "jquery";
import getModal from "../js/modal.js";

const modal = getModal();

$(() => {
	$("#login-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/login",
			data: {
				userName: $("#username-field").val(),
				password: $("#password-field").val(),
			},
			dataType: "text",
		}).then(() => {
			window.location.assign("/");
		}).fail((err) => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
