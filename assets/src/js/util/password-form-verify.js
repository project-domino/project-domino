import $ from "jquery";
import zxcvbn from "zxcvbn";

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
const verifyPassword = (field, notify) => {
	return () => {
		var verification = zxcvbn(field.val());
		var notificationContent = ["Password Strength: " +
			scoreToString(verification.score)];

		if(verification.score < 2) {
			if(verification.feedback.suggestions) {
				notificationContent.push(
					$("<br>"),
					verification.feedback.suggestions[0]
				);
			}

			field.addClass("invalid");
			notify
				.addClass("invalid")
				.removeClass("hidden")
				.empty()
				.append(notificationContent);
		} else {
			field.removeClass("invalid");
			notify
				.removeClass("invalid")
				.removeClass("hidden")
				.empty()
				.append(notificationContent);
		}
	};
};
const verifyRetypePassword = (passwordField, retypePasswordField, notify) => {
	return () => {
		var valid = passwordField.val() === retypePasswordField.val();
		if(valid) {
			retypePasswordField.removeClass("invalid");
			notify
				.removeClass("invalid")
				.addClass("hidden")
				.text("");
		} else {
			retypePasswordField.addClass("invalid");
			notify
				.addClass("invalid")
				.removeClass("hidden")
				.text("Passwords do not match.");
		}
	};
};

export {verifyPassword, verifyRetypePassword};
