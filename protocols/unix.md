# `unix`

This protocol encodes a Unix domain socket path to a resource. In the string
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

The absence of a `/` character at the start of the decoded address indicates a
relative path, otherwise the path is absolute.
