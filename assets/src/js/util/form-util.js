/** @module form-util */

/**
 * FormUtil contains utility functions for forms
 */
class FormUtil {

	// Checks if a field in a form is filled
	verifyFilled(field, notify, text) {
		return () => {
			if(field.val()) {
				field.removeClass("invalid");
				notify
					.removeClass("invalid")
					.addClass("hidden")
					.text("");
			} else {
				field.addClass("invalid");
				notify
					.addClass("invalid")
					.removeClass("hidden")
					.text(text);
			}
		};
	}

}

export default FormUtil;
