# `dns4`

`dns4` is a protocol that defines which domain name should be used, but only by resolving it to a IPv4 address.

## Representation Format

### Human-readable

The human-readable format of the `dns4` protocol uses the well-known textual representation:

	example.com

TODO: Consider also supporting with trailing dot.
	
### Binary

TODO: Consider binary format. Is it a string prefixed by an unsigned varint or should the FQDN encoding be used, where every label is prefixed and the last label has zero length?

## Binary Size

Unknown