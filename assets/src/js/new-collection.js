import $ from "jquery";
import _ from "lodash";
import WriterPanelCollectionUtil from "./writer-panel-collection-util.js";

const util = new WriterPanelCollectionUtil();

$(() => {
	util.initTagSelector();
	$(".note-search-field").on("keyup", _.debounce(util.searchHandler, 250));
	$(".note-search-btn").click(util.searchHandler);
});
