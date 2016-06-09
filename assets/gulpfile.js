require("./build_helpers/main.js")({
	js: {
		editor: "src/editor/main.js",
		main:   "src/js/main.js",
	},
	pug: {
		editor: "src/pug/editor.pug",
		index:  "src/pug/index.pug",
		search: "src/pug/search.pug",
	},
	sass: {
		editor: "src/sass/editor.scss",
	},
});
