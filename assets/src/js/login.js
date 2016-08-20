import $ from "jquery";
import FormUtil from "./util/form-util.js";
import {errorHandler} from "./util/error.js";

const formUtil = new FormUtil();

var verifyUserName = formUtil.verifyFilled(
	$("#username-field"),
	$(".username-notify"),
	"Username is required."
);
var verifyPassword = formUtil.verifyFilled(
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
		}).then(() => {
			window.location.assign("/");
		}).fail(errorHandler);
	});

	$("#username-field, #password-field").keyup(e => {
		if(e.keyCode === 13)
			$("#login-btn").click();
	});
});
