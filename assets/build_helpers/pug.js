const gulp =    require("gulp");
const plumber = require("gulp-plumber");
const pug =     require("gulp-pug");
const pugLint = require("gulp-pug-lint");
const rename =  require("gulp-rename");

module.exports = (file, out, dev = false) => gulp.src(file)
	.pipe(plumber())
	.pipe(pugLint())
	.pipe(pug({locals: {dev: dev}}))
	.pipe(rename(`${out}.html`))
	.pipe(gulp.dest("dist/templates/"));
