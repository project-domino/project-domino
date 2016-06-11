import $ from "jquery";

$(() => {
	$("#login-btn").click(() => {
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
		}).fail((err) => {
			console.log(err);
		});
	});
});
