# `ip6`

`ip6` is a protocol that defines which IPv6 address should be used.

## Representation Format

### Human-readable

The human-readable format of the `ip6` protocol uses the colon-seperated format:

	2604:1380:4602:5c00::3

TODO: Consider also supporting decimal format and/or format that encloses the IP address with `[` and `]`.
	
### Binary

The binary format of the `ip6` protocol uses the well-known binary format of 128 bits:

	0x26 0x04 0x13 0x80 0x46 0x02 0x5c 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x03

## Binary Size

128 bits (Implicit)