# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [Unreleased]

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

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
