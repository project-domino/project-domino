/** @module writer-panel-collection-util */

import $ from "jquery";
import "select2";

import WriterPanelUtil from "./writer-panel-util.js";

/**
 * WriterPanelNoteUtil contains utility functions for
 * writer-panel-note pages
 */
class WriterPanelNoteUtil extends WriterPanelUtil {

	// Constructs new WriterPanelNoteUtil
	// If a note is passed the contents of the page will be set to the note
	constructor(note) {
		super();

		// Initialize tag selector
		super.initTagSelector();

		// Initialize quill
		this.quill = new Quill("#editor", { // eslint-disable-line no-undef
			modules: {
				"toolbar": {
					container: "#editor-toolbar",
				},
				"image-tooltip": true,
				"link-tooltip":  true,
			},
			theme: "snow",
		});

		// If a note is passed, set contents of quill and tag-selector
		if(note) {
			if(note.Tags) {
				$(".tag-selector").val(note.Tags.map(e => {
					return e.ID;
				})).trigger("change");
			}
			this.quill.setHTML(note.Body);
		}
	}

	/**
	 * getData returns data from note page
	 */
	getData() {
		return {
			title:       $(".note-title").val(),
			description: $(".note-description").val(),
			body:        this.quill.getHTML(),
			tags:        $(".tag-selector").val().map(e => {return parseFloat(e);}),
		};
	}

}

export default WriterPanelNoteUtil;
