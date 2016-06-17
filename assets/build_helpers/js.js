const path = require("path");

const gulp =       require("gulp");
const addsrc =     require("gulp-add-src");
const concat =     require("gulp-concat");
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
			"select2",
			"stacktrace-js",
			"zxcvbn",
		];
		rollupConfig.globals = {
			"jquery":        "jQuery",
			"stacktrace-js": "StackTrace",
			"zxcvbn":        "zxcvbn",
		};
	}

	const inner = [rollup(rollupConfig)];
	if(dev) {
		inner.push(addsrc.prepend([
			"node_modules/jquery/dist/jquery.min.js",
			"node_modules/select2/dist/js/select2.js",
			"node_modules/stacktrace-js/stacktrace.js",
			"node_modules/zxcvbn/dist/zxcvbn.js",
		]));
		inner.push(concat("rename-me.js"));
	} else {
		inner.push(uglify());
	}

	return [
		plumber(),
		xo(),
	].concat(inner).concat([
		rename(`${out}.js`),
		sourcemaps.write("."),
	]).reduce((pipeline, stage) => {
		return pipeline.pipe(stage);
	}, gulp.src(file, {read: false}));
};
