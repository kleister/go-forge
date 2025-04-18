# Changelog

## [2.1.0](https://github.com/kleister/go-forge/compare/v2.0.0...v2.1.0) (2025-04-14)


### Features

* upgrade revive and golangci-lint and make lint happy ([a3d697b](https://github.com/kleister/go-forge/commit/a3d697b6a0893c9b7551e1ea2fc49cc0c47df2b1))

## [2.0.0](https://github.com/kleister/go-forge/compare/v1.0.1...v2.0.0) (2024-06-05)


### Features

* apply new repo structure and integrate automated releases ([ff0488e](https://github.com/kleister/go-forge/commit/ff0488e7c1245a5ebed2bebd295347ef4e24bb70))
* **major:** update actions/setup-go action to v4 ([#26](https://github.com/kleister/go-forge/issues/26)) ([c79cbbe](https://github.com/kleister/go-forge/commit/c79cbbe047d50d38192fccbcdf46912c15fa5f7f))


## [1.0.1]() (2022-10-27)

The following sections list the changes for 1.0.1.

### Summary

 * Fix #21: Fixed filtering by Minecraft version

### Details

 * Bugfix #21: Fixed filtering by Minecraft version

   The filter for Minecraft versions haven't been working as expected because of a typo within the
   filter, now filtering by Mionecraft versions works as expected.

   https://github.com/kleister/go-forge/issues/21


## [1.0.0]() (2022-06-23)

The following sections list the changes for 1.0.0.

### Summary

 * Chg #23: Dropped support for latest and recommended
 * Chg #23: Replaced Forge version source URL

### Details

 * Change #23: Dropped support for latest and recommended

   Since we had to change the URL for fetching te forge version information we also had to drop the
   attributes to show the latest and the recommended version of Forge.

   https://github.com/kleister/go-forge/issues/23

 * Change #23: Replaced Forge version source URL

   Since the old used URL for fetching all available Forge versions doesn't work anymore we had to
   replace it with a new URL which sadly offers a different set of attribuites.

   https://github.com/kleister/go-forge/issues/23


## [0.2.1]() (2018-11-05)

The following sections list the changes for 0.2.1.

### Summary

 * Chg #2: Switch to go modules instead of gopkg
 * Chg #2: Use standard JSON parser instead of jsoniter

### Details

 * Change #2: Switch to go modules instead of gopkg

   As go modules have been the new standard we switched to it in favor of using gopkg.

   https://github.com/kleister/go-forge/pull/2

 * Change #2: Use standard JSON parser instead of jsoniter

   Intitially we implemented jsoniter as a JSON parser library, but in the end we decided that it's
   better to stick with the standard library JSON parser.

   https://github.com/kleister/go-forge/pull/2


## [0.2.0]() (2018-08-15)

The following sections list the changes for 0.2.0.

### Summary

 * Chg #1: Integrate the sortable interface
 * Chg #1: Integrate a version filter
 * Chg #23: Limited list of versions

### Details

 * Change #1: Integrate the sortable interface

   To get a properly sorted list of Forge versions we implemented the sortable interface for a
   slice of versions.

   https://github.com/kleister/go-forge/pull/1

 * Change #1: Integrate a version filter

   To be able to filter out specific versions we added a filtering method to a slice of Forge
   versions.

   https://github.com/kleister/go-forge/pull/1

 * Change #23: Limited list of versions

   To reduce the complexity of the list of available versions we added a default limit to list only
   versions for Minecraft newer than v1.6.4.

   https://github.com/kleister/go-forge/issues/23


## [0.1.0]() (2018-08-15)

The following sections list the changes for 0.1.0.

### Summary

 * Chg #23: Initial release of basic version

### Details

 * Change #23: Initial release of basic version

   Just prepared an initial basic version which could be released to the public.

   https://github.com/kleister/go-forge/issues/23
