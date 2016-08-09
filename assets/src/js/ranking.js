import $ from "jquery";

import getModal from "./util/modal.js";

const modal = getModal();

var changeVote = (rankingContainer, dir) => {
	if($("#logged-in-val").text() === "false") {
		modal.alert("You must sign in to rank items.", 3000);
		return;
	}

	var type = rankingContainer.data("type");
	var id = rankingContainer.data("id");
	var itemRankingElement = rankingContainer.children(".item-ranking");
	var ranking = itemRankingElement.text();

	var change = 0;

	if(dir === "1") {
		if(rankingContainer.hasClass("upvoted")) {
			dir = "0";
			change = -1;
			rankingContainer.removeClass("upvoted");
		} else if(rankingContainer.hasClass("downvoted")) {
			change = 2;
			rankingContainer.addClass("upvoted");
		} else {
			change = 1;
			rankingContainer.addClass("upvoted");
		}
		rankingContainer.removeClass("downvoted");
	} else if(dir === "-1") {
		if(rankingContainer.hasClass("downvoted")) {
			dir = "0";
			change = 1;
			rankingContainer.removeClass("downvoted");
		} else if(rankingContainer.hasClass("upvoted")) {
			change = -2;
			rankingContainer.addClass("downvoted");
		} else {
			change = -1;
			rankingContainer.addClass("downvoted");
		}
		rankingContainer.removeClass("upvoted");
	}

	itemRankingElement.text(parseFloat(ranking) + change);

	$.ajax({
		type: "PUT",
		url:  "/api/v1/" + type + "/" + id + "/vote",
		data: {
			dir: dir,
		},
		dataType: "text",
	});
};

$(() => {
	$(".item-upvote").click(function () {
		changeVote($(this).parent(), "1");
	});
	$(".item-downvote").click(function () {
		changeVote($(this).parent(), "-1");
	});
});

export {changeVote};
