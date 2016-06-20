const gulp =   require("gulp");
const rimraf = require("rimraf");
const watch =  require("gulp-watch");

gulp.task("clean", cb => rimraf("dist/", cb));

const watchTargets = [
	"build_helpers/**",
	"src/**",
	"gulpfile.js",
];

gulp.task("watch", () => {
	watch(watchTargets, () => {
		gulp.start("default");
	});
});
