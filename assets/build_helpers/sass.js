const gulp =       require("gulp");
const cleanCSS =   require("gulp-clean-css");
const plumber =    require("gulp-plumber");
const rename =     require("gulp-rename");
const sass =       require("gulp-sass");
// const sassLint =   require("gulp-sass-lint");
const sourcemaps = require("gulp-sourcemaps");

module.exports = (file, out) => gulp.src(file)
	.pipe(plumber())
	// .pipe(sassLint())
	// .pipe(sassLint.format())
	// .pipe(sassLint.failOnError())
	.pipe(sourcemaps.init())
	.pipe(sass())
	.pipe(cleanCSS())
	.pipe(rename(`${out}.css`))
	.pipe(sourcemaps.write());
