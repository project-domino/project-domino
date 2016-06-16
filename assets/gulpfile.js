require("./build_helpers/main.js")(require("./targets.json"));

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
