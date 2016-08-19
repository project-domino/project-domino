import $ from "jquery";
import getModal from "./util/modal.js";
import {verifyPassword, verifyRetypePassword} from "./util/password-form-verify.js";

const modal = getModal();

var usernameTest = new RegExp("^[a-zA-Z0-9_-]*$");

var verifyUserName = () => {
	var field = $("#username-field");
	var notify = $(".username-notify");
	var invalid = false;
	var notificationText = "";

	if(field.val() === "") {
		invalid = true;
		notificationText = "Username is required";
	} else if(!usernameTest.test(field.val())) {
		invalid = true;
		notificationText =
			"Only _, -, letters and numbers are allowed in a username.";
	}

	if(invalid) {
		field.addClass("invalid");
		notify
			.addClass("invalid")
			.removeClass("hidden")
			.empty()
			.text(notificationText);
	} else {
		field.removeClass("invalid");
		notify
			.removeClass("invalid")
			.addClass("hidden")
			.empty();
	}
};

var verifyPasswordField = verifyPassword($("#password-field"), $(".password-notify"));
var verifyRetypePasswordField = verifyRetypePassword(
	$("#password-field"),
	$("#retype-password-field"),
	$(".retype-password-notify")
);

$(() => {
	// Verify Handlers
	$("#password-field").keyup(verifyPasswordField);
	$("#retype-password-field").keyup(verifyRetypePasswordField);
	$("#username-field").keyup(verifyUserName);

	// Register Handler
	$("#register-btn").click(() => {
		verifyUserName();
		verifyPasswordField();
		verifyRetypePasswordField();
		if($(".invalid").length !== 0)
			return;
		$.ajax({
			type: "POST",
			url:  "/register",
			data: {
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
