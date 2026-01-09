# `unix`

This protocol encodes a Unix-domain socket path to a resource. In the string
representation, the path is encoded in a way consistent with a single URI Path
segment per [RFC 3986 Section 3.3](https://datatracker.ietf.org/doc/html/rfc3986#autoid-23).

Specifically following the grammar of a single `segment-nz`. In the binary
representation, no encoding is needed as the value is length prefixed.

When comparing multiaddrs, implementations should compare their binary
representation to avoid ambiguities over which characters were escaped.

## Examples

The following is a table of examples converting some common Unix paths to their
Multiaddr string form.

| Unix Path                   | Multiaddr string form                   |
| --------------------------- | --------------------------------------- |
| /                           | `/unix/%2F`                             |
| /file.socket                | `/unix/%2Ffile.socket`                  |
| /dir/file.socket            | `/unix/%2Fdir%2Ffile.socket`            |
| /dir/file.socket/p2p/12D... | `/unix/%2Fdir%2Ffile.socket/p2p/12D...` |
| /tmp/foo/../bar             | `/unix/%2Ftmp%2Ffoo%2F..%2Fbar`         |
| /%2F                        | `/unix/%252F`                           |
| /a%20space                  | `/unix/%2Fa%2520space`                  |
| /a%2Fslash                  | `/unix/%2Fa%252Fslash`                  |
| socket                      | `/unix/socket`                          |

## Usage

`/unix` would typically be found at the start of a multiaddr, however it may
appear anywhere, for example in the case where we route through some sort of
proxy server or SSH tunnel.

# `unix-abstract`

This protocol encodes a Linux Unix-domain abstract socket address,
which are distinguished by their first byte being 0.
It is encoded the same way as `unix`;
the marker byte is not part of the path.

## Examples

In the following table, the address column follows the userspace convention
of 0 bytes in the address being rendered as an `@` for abstract addresses
for display only.

| Rendered Address                             | multiaddr string form                                              |
| -------------------------------------------- | ------------------------------------------------------------------ |
| @f87a1c847a4ecaf3/bus/systemd/bus-api-system | `/unix-abstract/f87a1c847a4ecaf3%2Fbus%2Fsystemd%2Fbus-api-system` |
| @/run/fsid.sock@@@@@@@@@@@@@@@@@@@@@@@@@@... | `/unix-abstract/%2Frun%2Ffsid.sock%00%00%00%00%00%00%00%00%00...`  |

# `stream`, `seqpacket`, `dgram`

These correspond to the *type* of Unix-domain socket:
`SOCK_STREAM`, `SOCK_SEQPACKET`, `SOCK_DGRAM`.

Previous versions of this specification did not contain these types;
for compatibility, their absence should not indicate an error,
if a default makes sense when decoding.
