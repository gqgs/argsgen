# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [1.1.1](https://github.com/gqgs/argsgen/compare/v1.1.0...v1.1.1) (2023-10-07)


### Bug Fixes

* don't output positional args if not used ([cd27063](https://github.com/gqgs/argsgen/commit/cd27063c04761e6bbed5581639c57b6ad33a7157)), closes [#3](https://github.com/gqgs/argsgen/issues/3)

## [1.1.0](https://github.com/gqgs/argsgen/compare/v1.0.1...v1.1.0) (2021-05-26)


### Features

* allow required flag for time.Duration ([d122828](https://github.com/gqgs/argsgen/commit/d122828ade0601fd4360380d5acb4512b379a93a))
* support time.Duration ([396a99d](https://github.com/gqgs/argsgen/commit/396a99dd0e9d187b3cba15cb619673d41e12ba46)), closes [#1](https://github.com/gqgs/argsgen/issues/1)

### [1.0.1](https://github.com/gqgs/argsgen/compare/v1.0.0...v1.0.1) (2021-05-22)


### Bug Fixes

* alias not being defined ([d57cacb](https://github.com/gqgs/argsgen/commit/d57cacbed9fcbe20a088b4f0c8dea2f9bfb7b669))

## 1.0.0 (2021-05-22)


### Features

* generate comments for exported functions ([ef44f76](https://github.com/gqgs/argsgen/commit/ef44f7698ab5b1a3bc0a26b21fbab8eeced9311d))
* generate from template ([1a2ac47](https://github.com/gqgs/argsgen/commit/1a2ac479d54cb24b23c2488ed97e9b14f61ed2b3))
* generate MustParse ([6de26f0](https://github.com/gqgs/argsgen/commit/6de26f0fd4391c28f0edee2d0231b32d043d6583))
* graceful shutdown on MustParse error ([2204bb1](https://github.com/gqgs/argsgen/commit/2204bb18d4886e0a7cacd516e45a2c2f2f251906))
* required fields ([7b24f74](https://github.com/gqgs/argsgen/commit/7b24f74518fd26150d27ac67c4bfe18ed79b1bd4)), closes [#2](https://github.com/gqgs/argsgen/issues/2)
* support remaining built-in types ([1d78211](https://github.com/gqgs/argsgen/commit/1d78211c276ba3de0d8a7c672d87d01cd31ebc49))


### Bug Fixes

* check range to support partially defined positional ([c06e48c](https://github.com/gqgs/argsgen/commit/c06e48cf7c3b69ddca9732938463a9eff14ad8b7))
* embed templates ([1a8f2ff](https://github.com/gqgs/argsgen/commit/1a8f2ff85c70c10dc5de9d1e21a8c56cccfe1c74))
* flag alias ([6b21f46](https://github.com/gqgs/argsgen/commit/6b21f46e970ca149b0a00088ede8a8492d00824e))
* flagSet PrintDefaults ([e746e35](https://github.com/gqgs/argsgen/commit/e746e35383a095c8eb3b3ea1550188d504bbfbe8))
* improve unit test output ([1bf335f](https://github.com/gqgs/argsgen/commit/1bf335f8edbdca32a2c2aa1065c394702bdb784b))
* only import error if necessary ([fd17c87](https://github.com/gqgs/argsgen/commit/fd17c8704e71bb561fffe31ac715c38c6865b6b9))
* remove import extraneous space ([f5b0c44](https://github.com/gqgs/argsgen/commit/f5b0c443ce8512cbf2451e90d1b6e4c28e52d621))
* required check if not positional args ([5f91015](https://github.com/gqgs/argsgen/commit/5f9101509b2a61fd06e61a560e2a6ed46f60391f))
* s/ContinueOnError/ExitOnError/ ([645606a](https://github.com/gqgs/argsgen/commit/645606a039624f2a4143f69f2ed2578d077767ef))
* s/field/argument/g ([61f7342](https://github.com/gqgs/argsgen/commit/61f7342d484a4399cdc0a917d83c991b56e4d325))
* template path ([c399064](https://github.com/gqgs/argsgen/commit/c3990648e58e8bd8a5157fabcce0defc3c641e58))
