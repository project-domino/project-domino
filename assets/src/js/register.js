import $ from "jquery";
import getModal from "./util/modal.js";
import FormUtil from "./util/form-util.js";
import {verifyPassword, verifyRetypePassword} from "./util/password-form-verify.js";

const util = new FormUtil();
const modal = getModal();

var verifyUserName = util.verifyFilled(
	$("#username-field"),
	$(".username-notify"),
	"Username is required."
);
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
