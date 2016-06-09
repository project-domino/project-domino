const gulp =   require("gulp");
const rimraf = require("rimraf");
const watch =  require("gulp-watch");

gulp.task("clean", cb => rimraf("dist/", cb));
gulp.task("watch", () => {
	watch([
		"build_helpers/**",
		"src/**",
		"gulpfile.js",
	], () => {
		gulp.start("default");
	});
});
