/** @module writer-panel-collection-util */

import $ from "jquery";
import "select2";

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
		if(selectedNotes)
			this.selectedNotes = selectedNotes;
		this.resultNotes = [];
	}

	/**
	 * toggleSelectedNoteVisibility hides the selectedNote div if there are
	 * no selected notes
	 */
	toggleSelectedNoteVisibility() {
		if(this.selectedNotes.length > 0) {
			$(".selected-notes").removeClass("hidden");
			$(".no-notes-text").addClass("hidden");
		} else {
			$(".selected-notes").addClass("hidden");
			$(".no-notes-text").removeClass("hidden");
		}
	}
	/**
	 * numberSelections numbers selected notes in the DOM
	 */
	numberSelections() {
		$(".select-note-number").each(function (i) {
			$(this).text(i + 1);
		});
	}
	/**
	 * refreshSelectArray updates select array from DOM
	 */
	refreshSelectArray() {
		this.selectedNotes = $.makeArray($(".selected-notes").children()).map(e => {
			return JSON.parse($(e).data("note"));
		});
	}

	/**
	 * renderSelectNotes renders all selected notes in selectNotes
	 */
	renderSelectNotes() {
		this.toggleSelectedNotes();
		$(".selected-notes").empty().append(this.selectedNotes.map((note, i) => {
			return $("<div>").append(
				$("<div>").append(
					$("<div>").append(
						$("<span>").addClass("fa fa-trash").click(function () {
							$(this).closest(".list-item").remove();
							this.numberSelections();
							this.refreshSelectArray();
							this.toggleSelectedNotes();
						})
					).addClass("select-options-left"),
					$("<div>").append(
						$("<span>").addClass("fa fa-chevron-up").click(function () {
							var current = $(this).closest(".list-item");
							current.prev(".list-item").before(current);
							this.numberSelections();
							this.refreshSelectArray();
						}),
						$("<span>").addClass("fa fa-chevron-down").click(function () {
							var current = $(this).closest(".list-item");
							current.next(".list-item").after(current);
							this.numberSelections();
							this.refreshSelectArray();
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
			).addClass("list-item").data("note", JSON.stringify(note));
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
						$("<button>").click(function () {
							var note = JSON.parse($(this).parent().parent().data("note"));
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
		}).then(function (data) {
			this.resultNotes = data;
			this.renderResultNotes();
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	}

}

export default WriterPanelCollectionUtil;
