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
| /file.socket                | `/unix/file.socket`                     |
| /dir/file.socket            | `/unix/dir%2Ffile.socket`               |
| /dir/file.socket/p2p/12D... | `/unix/dir%2Ffile.socket/p2p/12D...`    |
| /tmp/foo/../bar             | `/unix/tmp%2Ffoo%2F..%2Fbar`            |
| /%2F                        | `/unix/%252F`                           |
| /a%20space                  | `/unix/a%2520space`                     |
| /a%2Fslash                  | `/unix/a%252Fslash`                     |

## Usage

`/unix` would typically be found at the start of a multiaddr, however it may
appear anywhere, for example in the case where we route through some sort of
proxy server or SSH tunnel.

The leading `/` character of the path can be omitted, unless it is the only
character in the path, in which case it must be escaped as normal.
