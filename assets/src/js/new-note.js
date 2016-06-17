import $ from "jquery";
import "select2";

$(() => {
	$(".tag-selector").select2({
		placeholder: "Type to search for tags...",
		allowClear:  true,
	});
});
