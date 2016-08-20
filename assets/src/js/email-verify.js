import $ from "jquery";
import {errorHandler} from "./util/error.js";
import getModal from "./util/modal.js";

const modal = getModal();

$(() => {
	$(".verify-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/email/verify",
		}).then(() => {
			modal.alert("Sent Verification Email");
		}).fail(errorHandler);
	});
});
