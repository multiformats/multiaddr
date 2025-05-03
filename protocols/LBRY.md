# `lbry`

`lbry` is a protocol that identifies specific content on the LBRY network. This can be a channel, a stream or a playlist.

## Representation Format

### Human-readable

The human-readable format of the `lbry` protocol is used in the same way as the URLs at [https://lbry.tech/spec](https://spec.lbry.com/#urls), but in this case the scheme is `/lbry/` instead of `lbry://`, so:

#### Stream Claim Name

	/lbry/meet-lbry

#### Channel Claim Name

	/lbry/@lbry

#### Channel Claim Name and Stream Claim Name

	/lbry/@lbry/meet-lbry

#### Claim ID

	/lbry/meet-lbry:7a0aa95c5023c21c098
	/lbry/meet-lbry:7a
	/lbry/@lbry:3f/meet-lbry

#### Sequence

	/lbry/meet-lbry*1
	/lbry/@lbry*1/meet-lbry

#### Amount Order

	/lbry/meet-lbry$2
	/lbry/meet-lbry$3
	/lbry/@lbry$2/meet-lbry
	
### Binary

TODO: Define a binary format for the `lbry` protocol.

## Binary Size

Variable (Explicit)