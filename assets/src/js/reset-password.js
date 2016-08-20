import $ from "jquery";

import getModal from "./util/modal.js";
import FormUtil from "./util/form-util.js";
import {errorHandler} from "./util/error.js";
import {verifyPassword, verifyRetypePassword} from "./util/password-form-verify.js";

const formUtil = new FormUtil();
const modal = getModal();

var verifyUserName = formUtil.verifyFilled(
	$("#username-field"),
	$(".username-notify"),
	"Username is required."
);
var verifyCode = formUtil.verifyFilled(
	$("#verification-code-field"),
	$(".verification-code-notify"),
	"Verification Code is required."
);
var verifyPasswordField = verifyPassword(
	$("#password-field"),
	$(".password-notify")
);
var verifyRetypePasswordField = verifyRetypePassword(
	$("#password-field"),
	$("#retype-password-field"),
	$(".retype-password-notify")
);

$(() => {
	$("#username-field").keyup(verifyUserName);
	$("#verification-code-field").keyup(verifyCode);
	$("#password-field").keyup(verifyPasswordField);
	$("#retype-password-field").keyup(verifyRetypePasswordField);

	$("#passwd-verification-send-btn").click(() => {
		verifyUserName();

		if($(".invalid").length !== 0)
			return;

		$.ajax({
			type: "POST",
			url:  "/reset-password",
			data: {
				userName: $("#username-field").val(),
			},
		}).then(() => {
			modal.alert("A verification was sent to your email.");
		}).fail(errorHandler);
	});
	$("#reset-passwd-btn").click(() => {
		verifyUserName();
		verifyCode();
		verifyPasswordField();
		verifyRetypePasswordField();

		if($(".invalid").length !== 0)
			return;

		$.ajax({
			type: "PUT",
			url:  "/reset-password",
			data: {
				userName:       $("#username-field").val(),
				resetCode:      $("#verification-code-field").val(),
				password:       $("#password-field").val(),
				retypePassword: $("#retype-password-field").val(),
			},
		}).then(() => {
			modal.alert("Your password has been reset.");
		}).fail(errorHandler);
	});
});
