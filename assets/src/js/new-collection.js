import $ from "jquery";
import WriterPanelUtil from "./writer-panel-util.js";

const util = new WriterPanelUtil();

$(() => {
	util.initTagSelector();
});
