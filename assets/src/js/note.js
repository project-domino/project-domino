import $ from "jquery";
import _ from "lodash";

import getModal from "./util/modal.js";

const modal = getModal();

// Initiates the input toggles
var initToggles = () => {
	$(".questions-link").click(function (e) {
		e.preventDefault();
		$(this).addClass("active");
		$(".suggestions-link").removeClass("active");
		$(".questions-container").removeClass("hidden");
		$(".suggestions-container").addClass("hidden");
	});
	$(".suggestions-link").click(function (e) {
		e.preventDefault();
		$(this).addClass("active");
		$(".questions-link").removeClass("active");
		$(".suggestions-container").removeClass("hidden");
		$(".questions-container").addClass("hidden");
	});

	$(".questions-input-placeholder").click(function (e) {
		e.preventDefault();
		$(this).addClass("hidden");
		$(".questions-input-container").removeClass("hidden");
		$(".question-input-area").focus();
	});
	$(".suggestions-input-placeholder").click(function (e) {
		e.preventDefault();
		$(this).addClass("hidden");
		$(".suggestions-input-container").removeClass("hidden");
		$(".suggestion-input-area").focus();
	});
	$(".question-cancel-button").click(() => {
		$(".questions-input-container").addClass("hidden");
		$(".questions-input-placeholder").removeClass("hidden");
	});
	$(".suggestion-cancel-button").click(() => {
		$(".suggestions-input-container").addClass("hidden");
		$(".suggestions-input-placeholder").removeClass("hidden");
	});
};

// Returns a jquery object from a given comment
var renderComment = (user, comment) => {
	var votingStatus = "";
	if(_(user.UpvoteComments).map(e => e.ID).value().contains(comment.ID))
		votingStatus = "upvoted";
	if(_(user.DownvoteComments).map(e => e.ID).value().contains(comment.ID))
		votingStatus = "downvoted";

	return $("<div>").addClass("list-item").append(
		$("<div>").addClass("item-left").append(
			$("<div>").addClass("item-ranking-container")
				.addClass(votingStatus)
				.attr({
					"data-type": "comment",
					"data-id":   comment.ID,
				})
				.append(
					$("<span>").addClass("fa fa-caret-up item-upvote"),
					$("<span>").addClass("item-ranking").text(comment.Ranking),
					$("<span>").addClass("fa fa-caret-up item-downvote")
				)
		),
		$("<div>").addClass("item-right").append(
			$("<div>").addClass("item-title").append(comment.User.UserName),
			$("<div>").addClass("item-date").append(comment.CreatedAt),
			$("<div>").addClass("item-description").append(comment.Body)
		)
	);
};

// Globals to store current comment page numbers
var suggestionPageNumber = 0;
var questionPageNumber = 0;

// Load next comment page
var loadComments = type => {
	// Get objects from page
	var note = JSON.parse($("#note-data").text());
	var user = JSON.parse($("#user-data").text());

	var page = 0;
	if(type === "question")
		page = questionPageNumber + 1;
	else if(type === "suggestion")
		page = suggestionPageNumber + 1;
	else
		return;

	// Get next page comments
	$.ajax({
		type: "GET",
		url:  "/api/v1/note/" + note.ID + "/comments/" + type,
		data: {
			page: page,
		},
		dataType: "json",
	}).then(data => {
		if(type === "question")
			questionPageNumber++;
		else if(type === "suggestion")
			suggestionPageNumber++;

		$("." + type + "s-comment-list > .other-comments").append(
			data.map(e => {
				return renderComment(user, e);
			})
		);
	}).fail(err => {
		console.log(err);
		modal.alert(err.responseText, 3000);
	});
};

window.loadComments = loadComments;

// postComment posts a comment
var postComment = (body, type, parent) => {
	var note = JSON.parse($("#note-data").text());
	var user = JSON.parse($("#user-data").text());

	if(user.ID === 0) {
		modal.alert("You must sign in to ask a question.", 3000);
		return;
	}
	if(body === "")
		return;

	$.ajax({
		type: "POST",
		url:  "/api/v1/note/" + note.ID + "/comments/" + type,
		data: {
			body:     body,
			parentID: parent,
		},
		dataType: "json",
	}).then(data => {
		$("." + type + "-input-area").val("");
		$(".questions-input-container, .suggestions-input-container").addClass("hidden");
		$(".questions-input-placeholder, .suggestions-input-placeholder").removeClass("hidden");

		$("." + type + "s-comment-list > .user-comments").append(
			renderComment(user, data)
		);
	}).fail(err => {
		console.log(err);
		modal.alert(err.responseText, 3000);
	});
};

$(() => {
	initToggles();

	$(".question-button").click(() => {
		postComment($(".question-input-area").val(), "question", "");
	});
	$(".suggestion-button").click(() => {
		postComment($(".suggestion-input-area").val(), "suggestion", "");
	});
});
