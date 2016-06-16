const _ =    require("lodash");
const gulp = require("gulp");
const zip =  require("gulp-zip");

const helpers = _([
	"js",
	"pug",
	"sass",
	"util",
]).map(name => [
	name,
	require(`./${name}.js`),
]).fromPairs().value();

module.exports = files => {
	const targets = _(files).map((targets, type) => {
		const helper = helpers[type];
		gulp.task(type, _.map(targets, (file, name) => {
			const targetName = `${type}:${name}`;
			gulp.task(targetName, () => {
				return helper(file, name)
					.pipe(gulp.dest("dist"));
			});
			return targetName;
		}));
		return type;
	}).flatten().value();

	gulp.task("default", targets, () => {
		gulp.src(["dist/**", "!dist/assets.zip", "!dist/doc"])
			.pipe(zip("assets.zip"))
			.pipe(gulp.dest("dist/"));
	});

	gulp.task("js-dev", _(files.js).map((file, name) => {
		const targetName = `js-dev:${name}`;
		gulp.task(targetName, () => helpers.js(file, name, true));
		return targetName;
	}).value());
	gulp.task("default-dev", targets.map(name => {
		if(name === "js")
			return "js-dev";
		return name;
	}));
};
