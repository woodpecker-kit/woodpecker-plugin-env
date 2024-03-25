# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.12.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.11.1...v1.12.0) (2024-03-25)

### ✨ Features

* add basic env to print CiSystemInfo for history management ([b4ade637](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/b4ade6378257e0cf6ce1f31aad0d73fe9dae1bf3))

## [1.11.1](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.11.0...v1.11.1) (2024-03-25)

### 🐛 Bug Fixes

* change load drone env or PLUGIN_ENV_FILE time ([37a7f6d4](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/37a7f6d44a8766d6804f5ce1ce63601b85d5b0d6))

## [1.11.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.10.0...v1.11.0) (2024-03-25)

### ✨ Features

* add and env file by `PLUGIN_ENV_FILE` ([176bf1a1](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/176bf1a1e7a8fb500559dc204179a45094315007))

* add kubernetes runner patch and env file by `PLUGIN_ENV_FILE` ([293597ad](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/293597ad05456d17e0f8a9c2f4b8208afcc917a7))

### 📝 Documentation

* update usage of doc ([706167fa](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/706167fabff854f585d4ebf2f8fd43618e7be419))

## [1.10.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.9.0...v1.10.0) (2024-03-20)

### BREAKING CHANGE:

* update Engineering Structure and add doc/docs.md for maintain

### ✨ Features

* refactor plugin as impl.go and update full unit test case ([7dacec4c](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/7dacec4c217c1e0219ee370e8931fb7da62e29a2))

### 📝 Documentation

* update basic doc for plugin ([a2c7ac69](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/a2c7ac6957d60b17c40e8e61a3daca7e8e0d8731))

## [1.9.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.8.0...v1.9.0) (2024-03-18)

### ✨ Features

* add OnlyArgsCheck for Plugin unit test, and let less print at global debug info ([ac624b05](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/ac624b05614383e3eadc6cc089586edcc92c35e3))

## [1.8.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.7.1...v1.8.0) (2024-03-18)

### ✨ Features

* print global debug info try print current user info ([01c7f832](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/01c7f832e9bc7a98b63ba64cd24328b1e4ee4f41))

* change default flag name and change Config to Settings let code more clear ([c5f43303](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/c5f43303000a2afa72efcb1e261a0299114a26d4))

## [1.7.1](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.7.0...v1.7.1) (2024-03-16)

### 🐛 Bug Fixes

* fix check env not use settings.not_empty_envs error ([2dab047e](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/2dab047e0f67cb7df728ec1935b6128c2a1a1898))

## [1.7.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.6.0...v1.7.0) (2024-03-16)

### ✨ Features

* add settings.not_empty_envs for support empty env check ([07356c76](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/07356c76b8d1f192938158babb93530d6350a4a0))

## [1.6.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.5.0...v1.6.0) (2024-03-16)

### ✨ Features

* github.com/woodpecker-kit/woodpecker-tools v1.18.0 and change metod ([f605a742](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/f605a742a6db35dbe85c8b4571f7216b5c6a0ee3))

## [1.5.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.4.0...v1.5.0) (2024-03-13)

### ✨ Features

* add flag settings.woodpecker_kit_steps_transfer_disable_out and base usage ([6dc24e69](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/6dc24e69a673f877a4797ffe30d28eda8d4a132d))

* github.com/woodpecker-kit/woodpecker-tools v1.17.0 and update test template ([a89d28f7](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/a89d28f7ee9ac08ba11269894deb4f093a0fe772))

### 📝 Documentation

* update temp-woodpecker-golang-plugin version ([2cd6e875](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/2cd6e8752d0f855ffc5e8fa099200d6035813024))

## [1.4.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.3.0...v1.4.0) (2024-03-08)

### ✨ Features

* github.com/woodpecker-kit/woodpecker-tools v1.15.0 ([a1f37c65](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/a1f37c655b74a59231621e934bf9fd755cf67f02))

## [1.3.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.2.0...v1.3.0) (2024-03-08)

### ✨ Features

* add TestCheckArgsPlugin template for args test ([f3196c78](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/f3196c78001e196b7dc88c4b4efe9961710534e7))

* add github.com/sinlov-go/go-common-lib v1.7.0 ([7698d6f2](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/7698d6f2133198e30e7e814a97e8aab743c805f8))

## [1.2.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.1.1...v1.2.0) (2024-03-06)

### ✨ Features

* github.com/woodpecker-kit/woodpecker-tools v1.14.0 ([58df6e4f](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/58df6e4fd721485eea734d2d3b7627df83aecd95))

## [1.1.1](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.1.0...v1.1.1) (2024-03-06)

### 👷‍ Build System

* update temp-woodpecker-golang-plugin ([c5352119](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/c5352119692836a4d7a0c8b6484ce73c5a939684))

* github.com/woodpecker-kit/woodpecker-tools v1.13.0 ([81933a7d](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/81933a7d6cbee229f33193f49ad32781635a54f1))

## [1.1.0](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.0.1...v1.1.0) (2024-03-06)

### ✨ Features

* add plugin_test template for test StepTransfer ([929f2f4c](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/929f2f4ccde2e49175adac2ea22d330cc0099b93))

* add zymosis kit for auto build ([9844deb2](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/9844deb2387c720e695e970367b1555270d1f9a5))

* add plugin test case use env PLUGIN_DEBUG ([8777c71d](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/8777c71d00b35a686001fa94cdd28966d535a870))

* add plugin_test has envMustArgsCheck() ([e6d79d12](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/e6d79d12bf7775b940d831a0b819491fd1eb8da4))

* update depends ([a2145e5f](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/a2145e5f0875761c8e9030d23f8a7da0cc63ac2a))

* update github.com/woodpecker-kit/woodpecker-tools v1.7.0 and add annotation for fast ([d4e8ecd5](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/d4e8ecd55edb9600df475512a93deba38e6e7ce7))

### 👷‍ Build System

* update temp-woodpecker-golang-plugin version for build ([6c8c7b97](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/6c8c7b977416ddbcc8577c4d2d8de292975414b2))

* action ([efa9eeee](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/efa9eeee0abc42d880ed10dd6858933f030ebbe5))

* github.com/woodpecker-kit/woodpecker-tools v1.11.0 ([54130037](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/541300374bfbae7f74812c434640ce6468055ddf))

## [1.0.1](https://github.com/woodpecker-kit/woodpecker-plugin-env/compare/1.0.0...v1.0.1) (2024-02-28)

### 🐛 Bug Fixes

* fix golang-codecov at github aciton build ([4c18f8b0](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/4c18f8b0dcca96cf29a1ef2c0996b76b00e5491a))

### 📝 Documentation

* remove docker hub badges ([b86cca60](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/b86cca60e95b6e72db5926dc5afea53678ffcdfa))

## 1.0.0 (2024-02-28)

### ✨ Features

* add temp-woodpecker-golang-plugin and update workflow ([e486bfde](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/e486bfde2115c43f30698fd9bbc11a98ce609f7a))

* add steps_transfer_demo: false and code ([fdc77386](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/fdc773866d0658df47a0e68abda7ad3335980d64))

* add Plugin loadStepsTransfer() and saveStepsTransfer() ([c95bf225](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/c95bf225d8d6f298daf8b452fc791f60455ff8da))

* template of golang cli tools, and can print basic env info ([645c73ed](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/645c73ed3ab7e21970941f99c47b622d54bd485d))

### ♻ Refactor

* let plugin code more clear ([bf56bf2a](https://github.com/woodpecker-kit/woodpecker-plugin-env/commit/bf56bf2af91cf69d4b0fd35167cc24f2830dd698))
