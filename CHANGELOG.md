# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [Unreleased]

### Added

- Added `String()` method to `PanicSentinel`

### Changed

- **[BC]** Rename `codec.Codec.MediaType()` to `BasicMediaType()`

### Removed

- **[BC]** Remove `MarshalMessage()`, `UnmarshalMessage()`, `MustMarshalMessage()`, and `MustUnmarshalMessage()`

## [0.4.0] - 2020-11-07

### Changed

- Updated Dogma to v0.9.0

## [0.3.0] - 2020-11-03

### Changed

- Updated Dogma to v0.8.0

## [0.2.2] - 2020-03-26

### Added

- Add test fixtures for expected portable type names

## [0.2.1] - 2020-03-26

### Added

- Add `NewPacket()`
- Add `Packet.ParseMediaType()`

## [0.2.0] - 2020-01-23

### Changed

- **[BC]** Renamed `Marshaler` to `ValueMarshaler`
- **[BC]** The `Marshaler` interface is now union of `TypeMarshaler` and `ValueMarshaler`

## [0.1.0] - 2019-12-09

- Initial release

<!-- references -->
[Unreleased]: https://github.com/dogmatiq/marshalkit
[0.1.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.1.0
[0.2.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.2.1
[0.3.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.3.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
