const _ =    require("lodash");
const gulp = require("gulp");
const zip =  require("gulp-zip");

const helpers = _([
	"copy",
	"js",
	"pug",
	"sass",
	"util",
]).map(name => [
	name,
	require(`./${name}.js`),
]).fromPairs().value();

module.exports = files => {
	gulp.task("ckeditor-copy", () => {
		gulp.src("../ckeditor/**")
			.pipe(gulp.dest("dist/ckeditor/"));
	});
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
	}).flatten().value().concat("ckeditor-copy");

	gulp.task("default", targets, () => {
		gulp.src(["dist/**", "!dist/assets.zip", "!dist/doc"])
			.pipe(zip("assets.zip"))
			.pipe(gulp.dest("dist/"));
	});
};
