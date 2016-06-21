import $ from "jquery";
import "select2";
import getModal from "../js/modal.js";

const modal = getModal();

$(() => {
	window.quill = new Quill("#editor", { // eslint-disable-line no-undef
		modules: {
			"toolbar": {
				container: "#editor-toolbar",
			},
			"image-tooltip": true,
			"link-tooltip":  true,
		},
		theme: "snow",
	});
	$(".tag-selector").select2({
		ajax: {
			url:      "/api/v1/search/tag",
			dataType: "json",
			delay:    250,
			cache:    true,
			width:    "100%",
			data:     function (params) {
				return {
					q: params.term,
				};
			},
			processResults: function (data) {
				if(data) {
					return {
						results: data.map(function (e) {
							return {
								id:   e.ID,
								text: e.Name + " - " + e.Description,
							};
						}),
					};
				}
				return {results: []};
			},
		},
		placeholder: "Type to search for tags...",
		allowClear:  true,
	});
	$(".save-draft-btn").click(() => {
		$.ajax({
			type: "POST",
			url:  "/api/v1/note",
			data: JSON.stringify({
				title: $(".new-note-title").val(),
				body:  window.quill.getHTML(),
				tags:  $(".tag-selector").val().map(e => {return parseFloat(e);}),
			}),
			dataType: "json",
		}).then(data => {
			console.log(data);
			modal.alert("Note Saved", 3000);
		}).fail(err => {
			console.log(err);
			modal.alert(err.responseText, 3000);
		});
	});
});
