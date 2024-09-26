# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## [0.9.0] - 2024-09-27

### Removed

- **[BC]** The test marshaler in the `fixtures` package no longer supports the
  messages that were defined in `dogma/fixtures` package, which was removed in
  Dogma v0.14.3.
- **[BC]** Removed all "canned" `Packet` and "portable name" values from the
  `fixtures` package.

## [0.8.0] - 2024-09-27

### Removed

- **[BC]** Removed support for CBOR encoding. Application developers should
  prefer Protocol Buffers where possible. JSON support remains for encoding of
  "arbitrary" types.

### Deprecated

- The `fixtures` package is now deprecated and will be removed in a future
  release.

## [0.7.6] - 2024-09-26

### Changed

- Bumped minimum Go version to v1.23.
- Added support for `enginekit` stubs to the test marshaler.

## [0.7.5] - 2024-08-17

### Added

- Added `MustMarshalCommandIntoEnvelope()`, `MustMarshalEventIntoEnvelope()` and `MustMarshalTimeoutIntoEnvelope()`.
- Added `UnmarshalCommandFromEnvelope()`, `UnmarshalEventFromEnvelope()` and `UnmarshalTimeoutFromEnvelope()`.

### Deprecated

- Deprecated `MustMarshalMessageIntoEnvelope()` and `UnmarshalMessageFromEnvelope()`.

## [0.7.4] - 2024-06-21

### Added

- Added `stateless.DefaultCodec`, a codec for marshaling/unmarshaling
  `dogma.StatelessProcessRoot` values without assuming that any other codec is
  available.

### Changed

- `codec.NewMarshaler()` no longer returns an error when two (or more) types
  have a conflicting portable name so long as at least one codec has a unique
  portable name for that type. Any single codec that produces the same portable
  name for multiple types is not used for those types.

## [0.7.3] - 2023-03-27

### Fixed

- Don't assume `dogma.Message` will always be equivalent to `any`

## [0.7.2] - 2021-06-09

### Fixed

- Fix mispelling of `json.DefaultCodec` and `cbor.DefaultCodec` (missing `t`)

### Changed

- Change `json.Codec` and `cbor.Codec` to use non-pointer receivers

## [0.7.1] - 2021-06-09

### Added

- Add `json.DefaultCodec` and `cbor.DefaultCodec`

## [0.7.0] - 2021-05-05

This release upgrades the protocol buffers implementation to use
`google.golang.org/protobuf` instead of `github.com/golang/protobuf`.

There have been several breaking changes to marshalkit's `protobuf` package,
although the functionality remains the same.

In summary, the various codec types in this package have been replaced by a
single `Codec` type which is configured for different encoding formats by using
different marshalers/unmarshalers provided by `google.golang.org/protobuf`.

### Added

- Add `protobuf.Codec`, `Marshaler` and `Unmarshaler`
- Add `protobuf.NativeBasicMediaType`, `DefaultNativeMarshaler` and `DefaultNativeMarshaler`
- Add `protobuf.JSONBasicMediaType`, `DefaultJSONMarshaler` and `DefaultJSONMarshaler`
- Add `protobuf.TextBasicMediaType`, `DefaultTextMarshaler` and `DefaultTextMarshaler`

### Removed

- Remove `protobuf.NativeCodec`, `JSONCodec` and `TextCodec`

## [0.6.0] - 2021-02-03

### Added

- Add `codec/cbor` package

### Changed

- **[BC]** `ValueMarshaler.MarshalAs()` now accepts multiple media-types in order of preference
- **[BC]** `ValueMarshaler.MarshalAs()` now returns a boolean to indicate unsupported media-types

### Fixed

- Fix `MarshalAs()` issue that prevented encoding when the media-type's portable name differed to that of the default codec

## [0.5.0] - 2021-01-20

### Added

- **[BC]** Added `MarshalAs()` method to `ValueMarshaler`
- **[BC]** Added `MediaTypesFor()` method to `ValueMarshaler`
- Added `String()` method to `PanicSentinel`
- Added `MustMarshalMessageIntoEnvelope()` and `UnmarshalMessageFromEnvelope()`
- Added `MustMarshalEnvelopeIdentity()` and `UnmarshalEnvelopeIdentity()`
- Added `MustMarshalEnvelopeTime()` and `UnmarshalEnvelopeTime()`

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

[unreleased]: https://github.com/dogmatiq/marshalkit
[0.1.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.1.0
[0.2.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.2.1
[0.3.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.3.0
[0.4.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.4.0
[0.5.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.5.0
[0.6.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.6.0
[0.7.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.7.0
[0.7.1]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.7.1
[0.7.3]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.7.3
[0.7.5]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.7.5
[0.7.6]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.7.6
[0.8.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.8.0
[0.9.0]: https://github.com/dogmatiq/marshalkit/releases/tag/v0.9.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
