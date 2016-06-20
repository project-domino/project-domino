const path = require("path");
const gulp =       require("gulp");
const plumber =    require("gulp-plumber");
const rename =     require("gulp-rename");
const rollup =     require("gulp-rollup");
const sourcemaps = require("gulp-sourcemaps");
const uglify =     require("gulp-uglify");
const xo =         require("gulp-xo");

module.exports = (file, out) => gulp.src(file, {read: false})
	.pipe(plumber())
	.pipe(xo())
	.pipe(rollup({
		format:     "umd",
		moduleName: path.basename(file, ".js"),
		plugins:    [
			require("rollup-plugin-babel")({
				exclude: "node_modules/**",
			}),
			require("rollup-plugin-commonjs")({
				include: "node_modules/**",
			}),
			require("rollup-plugin-node-resolve")({
				browser: true,
				jsnext:  true,
			}),
		],
		sourceMap: true,
	}))
	.pipe(uglify())
	.pipe(rename(`${out}.js`))
	.pipe(sourcemaps.write())
	.pipe(gulp.dest("dist/assets/"));
