# project-domino-client

This is the HTML/CSS/JS client for [Project Domino][project-domino].
It's actually written with:

 * [pug][pug] instead of HTML,
 * [scss][sass] instead of CSS,
 * and ES6 JavaScript, which is transpiled to ES5 JavaScript.

## Building

The project is automatically built into the `dist` directory when installing with `npm i`.
The project may be rebuilt without invoking `npm` by running `gulp`.

Existing build artifacts may be removed by running `gulp clean`.
The project sources may be watched, and a rebuild issued upon changes, by running `gulp watch`.

# Tests and Linting

All linters are run when building, and builds will (generally) not succeed unless code is lint-free.
Tests are run after installing with `npm i`.

All linters and tests may be run by running `npm test`.

[project-domino]: https://github.com/project-domino
[pug]: https://github.com/pugjs/pug
[sass]: https://github.com/sass/sass
