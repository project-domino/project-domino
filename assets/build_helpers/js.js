const path = require("path");

const gulp =       require("gulp");
const plumber =    require("gulp-plumber");
const rename =     require("gulp-rename");
const rollup =     require("gulp-rollup");
const sourcemaps = require("gulp-sourcemaps");
const uglify =     require("gulp-uglify");
const xo =         require("gulp-xo");

module.exports = (file, out, dev = false) => {
	const rollupPlugins = [
		require("rollup-plugin-babel")({
			exclude: "node_modules/**",
		}),
	];
	if(!dev) {
		rollupPlugins.push(require("rollup-plugin-commonjs")({
			include: "node_modules/**",
		}));
		rollupPlugins.push(require("rollup-plugin-node-resolve")({
			browser: true,
			jsnext:  true,
		}));
	}

	const rollupConfig = {
		format:     "umd",
		moduleName: path.basename(file, ".js"),
		plugins:    rollupPlugins,
		sourceMap:  true,
	};
	if(dev) {
		rollupConfig.external = [
			"jquery",
			"stacktrace-js",
		];
		rollupConfig.globals = {
			"jquery":        "jQuery",
			"stacktrace-js": "StackTrace",
		};
	}

	const inner = [
		rollup(rollupConfig),
	];
	if(!dev) inner.push(uglify());

	return [
		plumber(),
		xo(),
	].concat(inner).concat([
		rename(`${out}.js`),
		sourcemaps.write("."),
		gulp.dest("dist/assets"),
	]).reduce((pipeline, stage) => {
		return pipeline.pipe(stage);
	}, gulp.src(file, {read: false}));
};
