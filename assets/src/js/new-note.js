import $ from "jquery";
import "select2";

$(() => {
	var quill = new Quill("#editor", {
		modules: {
			"toolbar": {
				container: "#editor-toolbar",
			},
			"image-tooltip": true,
			"link-tooltip":  true,
		},
		theme: "snow",
	});
	window.quill = quill;
	$(".tag-selector").select2({
		placeholder: "Type to search for tags...",
		allowClear:  true,
	});
});
