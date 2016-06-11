require("./build_helpers/main.js")({
	js: {
		editor:   "src/editor/main.js",
		main:     "src/js/main.js",
		login:    "src/js/login.js",
		register: "src/js/register.js",
	},
	pug: {
		editor:   "src/pug/editor.pug",
		home:     "src/pug/home.pug",
		search:   "src/pug/search.pug",
		login:    "src/pug/login.pug",
		register: "src/pug/register.pug",
	},
	sass: {
		editor: "src/sass/editor.scss",
		home:   "src/sass/home.scss",
		header: "src/sass/header.scss",
		main:   "src/sass/main.scss",
		login:  "src/sass/login.scss",
	},
});

const gulp = require("gulp");
gulp.task("jsdoc", cb => {
	const jsdoc = require("gulp-jsdoc3");
	gulp.src(["src/editor/**.js", "src/editor/README.md"], {read: false})
		.pipe(jsdoc({
			opts: {
				destination: "dist/docs/",
			},
		}, cb));
});
