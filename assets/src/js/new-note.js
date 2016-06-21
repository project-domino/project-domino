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
});
