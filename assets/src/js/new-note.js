import $ from "jquery";
import "select2";

$(() => {
	$("#tag-selector").select2({
		tags:        "true",
		placeholder: "Type to search for tags...",
		allowClear:  true,
	});
});
