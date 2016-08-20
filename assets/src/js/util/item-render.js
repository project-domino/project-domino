/** @module item-render */

import $ from "jquery";

// tagList returns a list of tags
const tagList = tags => {
	return $("<div>")
		.addClass("tag-list")
		.append(
			tags.map(
				e => {
					return $("<a>")
						.addClass("tag")
						.attr({
							href:   encodeURI("/search/all?q=tag:\"" + e.Name + "\""),
							target: "_blank",
						})
						.text(e.Name);
				}
			)
		);
};

// note returns a note-item jquery object
const noteRender = (note, noteLefts) => {
	return $("<div>")
		.addClass("list-item")
		.append(noteLefts.map(n => n.addClass("item-left")))
		.append(
			$("<div>").addClass("item-right").append(
				$("<div>").addClass("item-title").append(
					$("<a>").attr({
						href:   "/note/" + note.ID,
						target: "_blank",
					}).text(note.Title)
				),
				$("<div>").addClass("item-date").text("Last Updated - " + note.UpdatedAt),
				$("<div>").addClass("item-description").text(note.Description)
			).append(tagList(note.Tags))
		)
		.data("note-id", note.ID);
};

export {noteRender, tagList};
