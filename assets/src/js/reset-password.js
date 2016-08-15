import $ from "jquery";
import getModal from "./util/modal.js";
import FormUtil from "./util/form-util.js";
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
		$.ajax({
			type: "POST",
			url:  "/reset-password",
			data: {
				userName: $("#username-field").val(),
			},
			dataType: "text",
		}).then(() => {
			modal.alert("A verification was sent to your email.", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
	$("#reset-passwd-btn").click(() => {
		verifyUserName();
		verifyCode();
		verifyPasswordField();
		verifyRetypePasswordField();
		$.ajax({
			type: "PUT",
			url:  "/reset-password",
			data: {
				userName:       $("#username-field").val(),
				resetCode:      $("#verification-code-field").val(),
				password:       $("#password-field").val(),
				retypePassword: $("#retype-password-field").val(),
			},
			dataType: "text",
		}).then(() => {
			modal.alert("Your password has been reset.", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
