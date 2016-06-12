/** @module modal */

import $ from "jquery";

/**
 * A Modal displays information above the page's main content. Modal instances
 * should be treated as singletons.
 * @see module:modal~getModal
 */
class Modal {
	/**
	 * Finds modal and alert elements in the page.
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
	 * @param {HTMLElement} header - Content to place in modal header
	 * @param {HTMLElement} body - Content to place in modal body
	 * @param {HTMLElement} footer - Content to place in modal footer
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
	 * @param {number} [time] - The amount of time, in ms, for the alert to stay open.
	 */
	alert(message, time) {
		this.alertText.text(message);
		this.alertElement.css("display", "block");
		if(time > 0)
			setTimeout(this.refreshAlert, time);
	}

	/**
	 * Closes the alert modal.
	 */
	refreshAlert() {
		$("#alert").css("display", "none");
	}
}

/**
 * getModal acts as a singleton getter/constructor for the Modal class.
 * @return {module:modal~Modal} The new or existing instance.
 */
const getModal = (() => {
	let instance = null;
	return () => {
		if(instance === null)
			instance = new Modal();
		return instance;
	};
})();

export default getModal;
