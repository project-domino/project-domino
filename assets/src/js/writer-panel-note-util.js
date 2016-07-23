/** @module writer-panel-collection-util */

import $ from "jquery";
import "select2";

import WriterPanelUtil from "./writer-panel-util.js";

const maxDescriptionChars = 250;

/**
 * WriterPanelNoteUtil contains utility functions for
 * writer-panel-note pages
 */
class WriterPanelNoteUtil extends WriterPanelUtil {

	// Constructs new WriterPanelNoteUtil
	// If a note is passed the contents of the page will be set to the note
	constructor(note) {
		super();

		// Get elements
		this.title = $(".note-title");
		this.description = $(".note-description");
		this.tagSelector = $(".tag-selector");
		this.remainChars = $(".char-remaining");

		// Initialize tag selector
		super.initTagSelector();

		// Initialize ckeditor
		CKEDITOR.replace("editor"); // eslint-disable-line no-undef
		this.editor = CKEDITOR.instances.editor; // eslint-disable-line no-undef

		// If a note is passed, set page contents
		if(note) {
			if(note.Tags) {
				$(".tag-selector").val(note.Tags.map(e => {
					return e.ID;
				})).trigger("change");
			}
			this.editor.setData(note.Body);
		}

		// Set chars remaining
		this.setRemainingChars();

		// Set listener on description field
		this.description.on("keyup", $.proxy(this.setRemainingChars, this));
	}

	/**
	 * setRemainingChars sets the charRemaining notifier for the description
	 */
	setRemainingChars() {
		var charRemaining = maxDescriptionChars - this.description.val().length;
		this.remainChars.text(charRemaining + " characters remaining...");
	}

	/**
	 * getData returns data from note page
	 */
	getData() {
		return {
			title:       this.title.val(),
			description: this.description.val(),
			body:        this.editor.getData(),
			tags:        this.tagSelector.val().map(e => {return parseFloat(e);}),
		};
	}

}

export default WriterPanelNoteUtil;
