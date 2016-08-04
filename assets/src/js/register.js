import $ from "jquery";
import zxcvbn from "zxcvbn";
import getModal from "./util/modal.js";
import FormUtil from "./util/form-util.js";

const util = new FormUtil();
const modal = getModal();

var scoreToString = score => {
	switch(score) {
	case 0:
		return "very weak";
	case 1:
		return "weak";
	case 2:
		return "ok";
	case 3:
		return "strong";
	case 4:
		return "very strong";
	default:
		return "?";
	}
};

// Verifiers
var verifyPassword = () => {
	var verification = zxcvbn($("#password-field").val());
	var notificationContent = ["Password Strength: " +
		scoreToString(verification.score)];

	if(verification.score < 2) {
		if(verification.feedback.suggestions) {
			notificationContent.push(
				$("<br>"),
				verification.feedback.suggestions[0]
			);
		}

		$("#password-field").addClass("invalid");
		$(".password-notify")
			.addClass("invalid")
			.removeClass("hidden")
			.empty()
			.append(notificationContent);
	} else {
		$("#password-field").removeClass("invalid");
		$(".password-notify")
			.removeClass("invalid")
			.removeClass("hidden")
			.empty()
			.append(notificationContent);
	}
};
var verifyRetypePassword = () => {
	var valid = $("#password-field").val() === $("#retype-password-field").val();
	if(valid) {
		$("#retype-password-field").removeClass("invalid");
		$(".retype-password-notify")
			.removeClass("invalid")
			.addClass("hidden")
			.text("");
	} else {
		$("#retype-password-field").addClass("invalid");
		$(".retype-password-notify")
			.addClass("invalid")
			.removeClass("hidden")
			.text("Passwords do not match.");
	}
};
var verifyUserName = util.verifyFilled(
	$("#username-field"),
	$(".username-notify"),
	"Username is required."
);

$(() => {
	// Verify Handlers
	$("#password-field").keyup(verifyPassword);
	$("#retype-password-field").keyup(verifyRetypePassword);
	$("#username-field").keyup(verifyUserName);

	// Register Handler
	$("#register-btn").click(() => {
		verifyUserName();
		verifyPassword();
		verifyRetypePassword();
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

	$("#password-field, #retype-password-field, #username-field").keyup(e => {
		if(e.keyCode === 13)
			$("#register-btn").click();
	});
});
