# `httppath`

This protocol encodes an HTTP Path to a resource. In the string representation,
the path is encoded in a way consistent with a single URI Path segment per [RFC 3986 Section
3.3](https://datatracker.ietf.org/doc/html/rfc3986#autoid-23). Specifically
following the grammar of a single `segment-nz`. In the binary
representation, no encoding is needed as the value is length prefixed.

When comparing multiaddrs, implementations should compare their binary
representation to avoid ambiguities over which characters were escaped.

## Examples

The following is a table of examples converting some common HTTP paths to their Multiaddr string form.

| HTTP Path       | Multiaddr string form            |
| --------------- | -------------------------------- |
| /               | n/a. This is implied.            |
| /user           | `/httppath/user`                 |
| /api/v0/login   | `/httppath/api%2Fv0%2Flogin`     |
| /tmp/foo/../bar | `/httppath/tmp%2Ffoo%2F..%2Fbar` |
| a%20space       | `/httppath/a%2520space`          |
| a%2Fslash       | `/httppath/a%252Fslash`          |

## Usage

`/httppath` should be appended to the end of an existing multiaddr, including after the peer id component (p2p). As an example, here's a multiaddr referencing the `.well-known/libp2p` HTTP resource along with a way to reach that peer:

```
/ip4/1.2.3.4/tcp/443/tls/http/p2p/12D.../httppath/.well-known%2Flibp2p
```

The `/httppath` component can also be appended to just the `/p2p/...` component, and rely on a separate peer discovery mechanism to actually identify the peer's address:

```
/p2p/12D.../httppath/.well-known%2Flibp2p
```
