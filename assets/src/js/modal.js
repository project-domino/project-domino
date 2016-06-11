/** @module modal */

import $ from "jquery";

/**
 * A Modal displays information above the DOM
 */
class Modal {
	/**
	 * Constructor finds modal and alert elements in DOM
	 */
	constructor() {
		this.modalElement = $("#modal");
		this.modalHeader = $("#modal-header");
		this.modalBody = $("#modal-body");
		this.modalFooter = $("#modal-footer");
		this.alertElement = $("#alert");
		this.alertText = $("#alert-text");
	}

	/**
	 * Opens the modal.
	 * @param {DOM Element} header - Content to place in modal header
	 * @param {DOM Element} body - Content to place in modal body
	 * @param {DOM Element} footer - Content to place in modal footer
	 */
	open(header, body, footer) {
		this.modalHeader.empty().append($("<span>").attr(
			"id", "modal-close"
		).text("×").click(() => {
			this.close();
		}).css("cursor", "pointer"), header);
		this.modalBody.empty().append(body);
		this.modalFooter.empty().append(footer);
		this.modalElement.css("display", "block");
	}

	/**
	 * Closes the modal.
	 */
	close() {
		this.modalElement.css("display", "none");
	}

	/**
	 * Opens the alert modal.
	 * @param {string} message - The message to display in the alert.
	 * @param {int} time - The amount of time, in ms, for the alert to stay open
	 */
	alert(message, time) {
		this.alertText.text(message);
		this.alertElement.css("display", "block");
		setTimeout(this.refreshAlert, time);
	}

	/**
	 * Closes the alert modal.
	 */
	refreshAlert() {
		$("#alert").css("display", "none");
	}

}

export default Modal;
