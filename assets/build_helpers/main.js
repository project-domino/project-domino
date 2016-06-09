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
			gulp.task(targetName, () => helper(file));
			return targetName;
		}));
		return type;
	}).flatten().value();

	gulp.task("default", targets, () => {
		gulp.src(["dist/**", "!dist/assets.zip"])
			.pipe(zip("assets.zip"))
			.pipe(gulp.dest("dist/"));
	});
};
