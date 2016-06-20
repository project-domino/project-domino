import $ from "jquery";
import "select2";

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
			url:         "/api/v1/search/tag",
			dataType:    "json",
			quietMillis: 250,
			cache:       true,
			width:       "100%",
			data:        function (params) {
				return {
					q: params.term,
				};
			},
			results: function (data) {
				return data.map(function (e) {
					return {
						id:  e.id,
						tag: e.Name + " - " + e.Description,
					};
				});
			},
		},
		placeholder: "Type to search for tags...",
		allowClear:  true,
	});
});
