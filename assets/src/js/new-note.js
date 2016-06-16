import $ from "jquery";
import "selectivity";

$(() => {
	$("#tag-selector").selectivity({
		items:       ["example", "example2", "example3"],
		multiple:    true,
		placeholder: "Type to search tags...",
	});
});
