const gulp =    require("gulp");
const plumber = require("gulp-plumber");
const rename =  require("gulp-rename");

module.exports = (file, out) => gulp.src(file)
	.pipe(plumber())
	.pipe(rename(`${out}`));
