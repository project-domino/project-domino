const gulp =    require("gulp");
const plumber = require("gulp-plumber");
const pug =     require("gulp-pug");
const pugLint = require("gulp-pug-lint");

module.exports = file => gulp.src(file)
	.pipe(plumber())
	.pipe(pugLint())
	.pipe(pug())
	.pipe(gulp.dest("dist/templates/"));
