/** @module writer-panel-collection-util */

import $ from "jquery";
import "select2";
import _ from "lodash";

import getModal from "../js/modal.js";
import WriterPanelUtil from "./writer-panel-util.js";

const modal = getModal();

/**
 * WriterPanelCollectionUtil contains utility functions for
 * writer-panel-collection pages
 */
class WriterPanelCollectionUtil extends WriterPanelUtil {

	// Constructs new WriterPanelCollectionUtil
	constructor(selectedNotes) {
		super();
		if(selectedNotes) {
			this.selectedNotes = selectedNotes;
			this.renderSelectNotes();
		} else {
			this.selectedNotes = [];
		}
		this.resultNotes = [];

		super.initTagSelector();
		var handler = $.proxy(this.searchHandler, this);
		$(".note-search-field").on("keyup", _.debounce(handler, 250));
		$(".note-search-btn").click(handler);
	}

	/**
	 * renderSelectNotes renders all selected notes in selectNotes
	 */
	renderSelectNotes() {
		if(this.selectedNotes.length > 0) {
			$(".selected-notes").removeClass("hidden");
			$(".no-notes-text").addClass("hidden");
		} else {
			$(".selected-notes").addClass("hidden");
			$(".no-notes-text").removeClass("hidden");
		}
		$(".selected-notes").empty().append(this.selectedNotes.map((note, i) => {
			return $("<div>").append(
				$("<div>").append(
					$("<div>").append(
						$("<span>").addClass("fa fa-trash").click(() => {
							_.remove(this.selectedNotes, e => {
								return e.ID === note.ID;
							});
							this.renderSelectNotes();
						})
					).addClass("select-options-left"),
					$("<div>").append(
						$("<span>").addClass("fa fa-chevron-up").click(() => {
							var i = _.findIndex(this.selectedNotes, e => {
								return e.ID === note.ID;
							});
							if(i > 0) {
								var temp = this.selectedNotes[i - 1];
								this.selectedNotes[i - 1] = this.selectedNotes[i];
								this.selectedNotes[i] = temp;
							}
							this.renderSelectNotes();
						}),
						$("<span>").addClass("fa fa-chevron-down").click(() => {
							var i = _.findIndex(this.selectedNotes, e => {
								return e.ID === note.ID;
							});
							if(i < this.selectedNotes.length) {
								var temp = this.selectedNotes[i + 1];
								this.selectedNotes[i + 1] = this.selectedNotes[i];
								this.selectedNotes[i] = temp;
							}
							this.renderSelectNotes();
						})
					).addClass("select-options-right")
				).addClass("item-left select-options-container"),
				$("<div>").append(
					$("<span>").text(i + 1).addClass("select-note-number")
				).addClass("item-left"),
				$("<div>").append(
					$("<div>").append(
						$("<a>").attr({
							href:   "/note/" + note.ID,
							target: "_blank",
						}).text(note.Title)
					).addClass("item-title"),
					$("<div>").append(
						$("<p>").text("description")
							.addClass("item-description")
					).addClass("item-description")
				).addClass("item-right")
			).addClass("list-item").data("note-id", note.ID);
		}));
	}

	/**
	 * renderResultNotes renders all selected notes in resultNotes
	 */
	renderResultNotes() {
		$(".search-result-container").empty();
		if(this.resultNotes.length) {
			$(".search-result-container").removeClass("hidden");
			$(".result-notification").addClass("hidden");
			$(".search-result-container").append(this.resultNotes.map(note => {
				return $("<div>").append(
					$("<div>").append(
						$("<button>").click(() => {
							if(this.selectedNotes.map(e => e.ID).includes(note.ID)) {
								modal.alert("This note has already been selected", 3000);
								return;
							}
							this.selectedNotes.push(note);
							this.renderSelectNotes();
						}).addClass("add-item-btn btn btn-primary fa fa-plus")
							.data("note-id", note.ID)
					).addClass("item-left"),
					$("<div>").append(
						$("<div>").append(
							$("<a>").attr({
								href:   "/note/" + note.ID,
								target: "_blank",
							}).text(note.Title)
						).addClass("item-title"),
						$("<div>").append(
							$("<p>").text("description")
								.addClass("item-description")
						).addClass("item-description")
					).addClass("item-right")
				).addClass("list-item").data("note", JSON.stringify(note));
			}));
		} else {
			$(".search-result-container").addClass("hidden");
			$(".result-notification").removeClass("hidden").text(
				"No notes found..."
			);
		}
	}

	/**
	 * searchHandler handles a user search
	 */
	searchHandler() {
		var query = $(".note-search-field").val();
		if(query === "") {
			$(".search-result-container").addClass("hidden");
			$(".result-notification").removeClass("hidden").text(
				"Please enter a search term..."
			);
			return;
		}
		$.ajax({
			type: "GET",
			url:  "/api/v1/search/note",
			data: {
				q: query,
			},
			dataType: "json",
		}).then(data => {
			this.resultNotes = data;
			this.renderResultNotes();
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	}

	/**
	 * getData returns data from collection page
	 */
	getData() {
		return {
			title:       $(".new-collection-title").val(),
			description: $(".new-collection-description").val(),
			notes:       this.selectedNotes.map(note => {return note.ID;}),
			tags:        $(".tag-selector").val().map(e => {return parseFloat(e);}),
		};
	}

}

export default WriterPanelCollectionUtil;
