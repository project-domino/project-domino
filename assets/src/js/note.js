import $ from "jquery";

import getModal from "./util/modal.js";
import {changeVote} from "./ranking.js";

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

	$(".questions-input-placeholder").focus(function (e) {
		e.preventDefault();
		$(this).addClass("hidden");
		$(".questions-input-container").removeClass("hidden");
		$(".question-input-area").focus();
	});
	$(".suggestions-input-placeholder").focus(function (e) {
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

// Globals to store current comment page numbers
var suggestionPageNumber = 0;
var questionPageNumber = 0;

// postComment posts a comment
var postComment = (body, type, parent, fn) => {
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
	}).then(fn).fail(err => {
		console.log(err);
		modal.alert(err.responseText, 3000);
	});
};

// renderCommentItem renders an individual comment
var renderCommentItem = (user, comment, showReply) => {
	var votingStatus = "";

	if(user.UpvoteComments) {
		if(user.UpvoteComments.map(e => e.ID).includes(comment.ID))
			votingStatus = "upvoted";
	}
	if(user.DownvoteComments) {
		if(user.DownvoteComments.map(e => e.ID).includes(comment.ID))
			votingStatus = "downvoted";
	}

	var itemRight = $("<div>").addClass("item-right").append(
		$("<div>").addClass("item-title").append(comment.User.UserName),
		$("<div>").addClass("item-date").append(comment.CreatedAt),
		$("<div>").addClass("item-description").append(comment.Body)
	);

	if(showReply) {
		itemRight.append(
			$("<div>").addClass("item-links").append(
				$("<a>").addClass("comment-reply-link").attr("href", "#")
					.attr("data-comment-id", comment.ID)
					.text("Reply").click(
						function (e) {
							e.preventDefault();
							$(".comment-input-container[data-comment-id='" + comment.ID + "']")
								.removeClass("hidden");
							$(this).addClass("hidden");
						}
				)
			)
		);
	}

	return $("<div>").addClass("list-item").append(
		$("<div>").addClass("item-left").append(
			$("<div>").addClass("item-ranking-container")
				.addClass(votingStatus)
				.attr({
					"data-type": "comment",
					"data-id":   comment.ID,
				})
				.append(
					$("<span>").addClass("fa fa-caret-up item-upvote").click(
						function () {
							changeVote($(this).parent(), "1");
						}
					),
					$("<span>").addClass("item-ranking").text(comment.Ranking),
					$("<span>").addClass("fa fa-caret-down item-downvote").click(
						function () {
							changeVote($(this).parent(), "-1");
						}
					)
				)
		),
		itemRight
	);
};

// Returns a jquery object from a given comment
var renderComment = (user, comment) => {
	var subComments = $("<div>").addClass("subcomments").attr("data-comment-id", comment.ID);
	if(comment.Children)
		subComments.append(comment.Children.map(e => renderCommentItem(user, e, false)));

	return $("<div>").addClass("comment").append(
		renderCommentItem(user, comment, true),
		$("<div>").addClass("comment-input-container hidden")
			.attr("data-comment-id", comment.ID)
			.append(
				$("<textarea>").addClass("comment-input-area")
					.attr("data-comment-id", comment.ID),
				$("<div>").addClass("comment-button-container").append(
					$("<button>").addClass("btn btn-primary").text("Reply").click(() => {
						postComment(
							$(".comment-input-area[data-comment-id='" + comment.ID + "']").val(),
							comment.Type,
							comment.ID,
							data => {
								$(".subcomments[data-comment-id='" + comment.ID + "']")
									.append(renderCommentItem(user, data, false));
								$(".comment-input-container[data-comment-id='" + comment.ID + "']")
									.addClass("hidden");
								$(".comment-reply-link[data-comment-id='" + comment.ID + "']")
									.removeClass("hidden");
							}
						);
					}),
					$("<button>").addClass("btn").text("Cancel").click(function () {
						$(this).parent().parent().addClass("hidden");
						$(".comment-reply-link[data-comment-id='" + comment.ID + "']")
							.removeClass("hidden");
					})
				)
			),
			subComments
		);
};

// Load next comment page
var loadComments = type => {
	var items = 25;

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
			page:  page,
			items: items,
		},
		dataType: "json",
	}).then(data => {
		if(data.length > 0) {
			if(type === "question")
				questionPageNumber++;
			else if(type === "suggestion")
				suggestionPageNumber++;

			$("." + type + "s-comment-list").append(
				data.map(e => {
					return renderComment(user, e);
				})
			);
		}

		if(data.length < items || data.length === 0) {
			$(".load-" + type + "s-container").empty().append(
				$("<span>").text("There are no more " + type + "s.")
			);
		}
	}).fail(err => {
		console.log(err);
		modal.alert(err.responseText, 3000);
	});
};

$(() => {
	initToggles();

	loadComments("question");
	loadComments("suggestion");

	$(".question-button").click(() => {
		postComment($(".question-input-area").val(), "question", "", () => {
			$(".question-input-area").val("");
			$(".questions-input-container").addClass("hidden");
			$(".questions-input-placeholder").removeClass("hidden");

			questionPageNumber = 0;

			$(".questions-comment-list").empty();
			loadComments("question");
		});
	});
	$(".suggestion-button").click(() => {
		postComment($(".suggestion-input-area").val(), "suggestion", "", () => {
			$(".suggestion-input-area").val("");
			$(".suggestions-input-container").addClass("hidden");
			$(".suggestions-input-placeholder").removeClass("hidden");

			suggestionPageNumber = 0;

			$(".suggestions-comment-list").empty();
			loadComments("suggestion");
		});
	});

	$(".load-questions-btn").click(() => loadComments("question"));
	$(".load-suggestions-btn").click(() => loadComments("suggestion"));
});
