import $ from "jquery";
import getModal from "./util/modal.js";
import FormUtil from "./util/form-util.js";
import {verifyPassword, verifyRetypePassword} from "./util/password-form-verify.js";
import {errorHandler} from "./util/error.js";

const formUtil = new FormUtil();
const modal = getModal();

var verifyOldPassword = formUtil.verifyFilled(
	$("#old-password-field"),
	$(".old-password-field"),
	"Old password is required."
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
	$("#old-password-field").keyup(verifyOldPassword);

	$(".change-password-btn").click(() => {
		verifyPasswordField();
		verifyRetypePasswordField();
		verifyOldPassword();
		if($(".invalid").length !== 0)
			return;
		$.ajax({
			type: "PUT",
			url:  "/account/password",
			data: {
				oldPassword:       $("#old-password-field").val(),
				newPassword:       $("#password-field").val(),
				newRetypePassword: $("#retype-password-field").val(),
			},
		}).then(() => {
			modal.alert("Your password has been changed.");
		}).fail(errorHandler);
	});
});
