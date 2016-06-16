import $ from "jquery";
import getModal from "../js/modal.js";

const modal = getModal();

$(() => {
	$("body").addClass("sidebar-close sidebar-default-close");

	$("#register-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/register",
			data: {
				email:          $("#email-field").val(),
				userName:       $("#username-field").val(),
				password:       $("#password-field").val(),
				retypePassword: $("#retype-password-field").val(),
			},
			dataType: "text",
		}).then(() => {
			window.location.assign("/");
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
