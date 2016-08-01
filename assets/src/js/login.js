import $ from "jquery";
import getModal from "../js/modal.js";
import FormUtil from "./form-util.js";

const modal = getModal();
const util = new FormUtil();

var verifyUserName = util.verifyFilled(
	$("#username-field"),
	$(".username-notify"),
	"Username is required."
);
var verifyPassword = util.verifyFilled(
	$("#password-field"),
	$(".password-notify"),
	"Password is required."
);

$(() => {
	// Verify Handlers
	$("#username-field").keyup(verifyUserName);
	$("#password-field").keyup(verifyPassword);

	// Login Handler
	$("#login-btn").click(() => {
		verifyUserName();
		verifyPassword();
		if($(".invalid").length !== 0)
			return;
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
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});

	$("#username-field, #password-field").keyup(e => {
		if(e.keyCode === 13)
			$("#login-btn").click();
	});
});
