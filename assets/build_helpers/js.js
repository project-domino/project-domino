const path = require("path");

const _ =          require("lodash");
const gulp =       require("gulp");
const plumber =    require("gulp-plumber");
const rename =     require("gulp-rename");
const rollup =     require("gulp-rollup");
const sourcemaps = require("gulp-sourcemaps");
const uglify =     require("gulp-uglify");
const xo =         require("gulp-xo");

const rollupPlugins = _.map([
	"babel",
], name => require(`rollup-plugin-${name}`)());

module.exports = (file, out) => gulp.src(file, {read: false})
	.pipe(plumber())
	.pipe(xo())
	.pipe(rollup({
		format:     "umd",
		moduleName: path.basename(file, ".js"),
		plugins:    rollupPlugins,
		sourceMap:  true,
	}))
	.pipe(uglify())
	.pipe(rename(`${out}.js`))
	.pipe(sourcemaps.write())
	.pipe(gulp.dest("dist/assets/"));
