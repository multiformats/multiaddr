# `httppath`

This protocol encodes an HTTP Path to a resource. In the string representation,
characters in the reserved set are percent-encoded ( "/" becomes "%2F").
Percent-encoding itself is defined by [RFC 3986 Section
2.1](https://datatracker.ietf.org/doc/html/rfc3986#section-2.1). In the binary representation, no escaping is needed as the value is length prefixed.

To ease implementation and benefit from reusing existing percent-encoding logic
present in many environments (Go's
[url.PathEscape](https://pkg.go.dev/net/url@go1.21.0#PathEscape); JS's
[encodeURIComponent](http://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent);
and Rust's [many
crates](https://crates.io/search?q=percent%20encode&sort=downloads)), it is
acceptable to encode more characters than the reserved set. While it's not
necessary to encode a space " " as %20, it's is acceptable to do so.

When comparing multiaddrs, implementations should compare their binary
representation to avoid ambiguities over which characters were escaped.

## Reserved Characters

| Character | Reason                                                                                               |
| --------- | ---------------------------------------------------------------------------------------------------- |
| `/`       | Multiaddr component separator                                                                        |
| `%`       | Percent encoding indicator                                                                           |
| `?`       | Marks the end of an HTTP path                                                                        |
| `#`       | Reserved gen-delim character by [rfc3986](https://datatracker.ietf.org/doc/html/rfc3986#section-2.2) |
| `[`       | Reserved gen-delim character by [rfc3986](https://datatracker.ietf.org/doc/html/rfc3986#section-2.2) |
| `]`       | Reserved gen-delim character by [rfc3986](https://datatracker.ietf.org/doc/html/rfc3986#section-2.2) |

## Usage

`/httppath` should be appended to the end of an existing multiaddr, including after the peer id component (p2p). As an example, here's a multiaddr referencing the `.well-known/libp2p` HTTP resource along with a way to reach that peer:

```
/ip4/1.2.3.4/tcp/443/tls/http/p2p/12D.../httppath/.well-known%2Flibp2p
```

The `/httppath` component can also be appended to just the `/p2p/...` component, and rely on a separate peer discovery mechanism to actually identify the peer's address:

```
/p2p/12D.../httppath/.well-known%2Flibp2p
```
