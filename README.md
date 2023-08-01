# multiaddr

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](https://protocol.ai)
[![](https://img.shields.io/badge/project-multiformats-blue.svg?style=flat-square)](https://github.com/multiformats/multiformats)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](https://webchat.freenode.net/?channels=%23ipfs)
[![](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> Composable and future-proof network addresses

- [Introduction](#introduction)
- [Use cases](#use-cases)
  - [Encapsulation based on context](#encapsulation-based-on-context)
- [Specification](#specification)
  - [Encoding](#encoding)
  - [Decoding](#decoding)
- [Protocols](#protocols)
- [Implementations](#implementations)
- [Contribute](#contribute)
- [License](#license)

Captain: [@lgierth](https://github.com/lgierth)


## Introduction

Multiaddr aims to make network addresses future-proof, composable, and efficient.

Current addressing schemes have a number of problems.

1. They hinder protocol migrations and interoperability between protocols.
2. They don't compose well. There are plenty of X-over-Y constructions,
   but only few of them can be addressed in a classic URI/URL or host:port scheme.
3. They don't multiplex: they address ports, not processes.
4. They're implicit, in that they presume out-of-band values and context.
5. They don't have efficient machine-readable representations.

Multiaddr solves these problems by modelling network addresses as arbitrary encapsulations of protocols.

- Multiaddrs support addresses for any network protocol.
- Multiaddrs are self-describing.
- Multiaddrs conform to a simple syntax, making them trivial to parse and construct.
- Multiaddrs have human-readable and efficient machine-readable representations.
- Multiaddrs encapsulate well, allowing trivial wrapping and unwrapping of encapsulation layers.

Multiaddr was originally [thought up by @jbenet](https://github.com/jbenet/random-ideas/issues/11).

## Interpreting multiaddrs

Multiaddrs are parsed from left to right, but they should be interpreted right
to left. Each component of a multiaddr wraps all the left components in its
context. For example, the multiaddr `/dns4/example.com/tcp/1234/tls/ws/tls`
(ignore the double encryption for now) is interpreted by taking the first `tls`
component from the right and interpreting it as the libp2p security protocol to
use for the connection, then passing the rest of the multiaddr to the websocket
transport to create the websocket connection. The websocket transport sees
`/dns4/example.com/tcp/1234/tls/ws/` and interprets the `tls` in this context to
mean that this is going to be a secure websocket connection. The websocket
transport also gets the host to dial along with the tcp port from the rest of
the multiaddr.

Components to the right can also provide parameters to components to the left,
since they are in charge of the rest of the multiaddr's interpretation. For
example, in `/ip4/1.2.3.4/tcp/1234/tls/p2p/QmFoo` the `p2p` component has the
value of the peer id and it passes it to the next component, in this case the
`tls` security protocol, as the expected peer id for this connection. Another
example is `/ip4/.../p2p/QmR/p2p-circuit/p2p/QmA`, here `p2p/QmA` is passed to
`p2p-circuit` and then the `p2p-circuit` component knows it needs to use the
rest of the multiaddr as the information to connect to the relay node.

This enables nesting and arbitrary parameters. A component can parse
arbitrary data with some encoding and pass it as a parameter to the next
component of the multiaddr. For example, we could reference a specific HTTP path
by composing `path` and `urlencode` components along with an `http` component.
This would look like
`/dns4/example.com/http/GET/path/percentencode/somepath%2ftosomething`. The
`percentencode` parses the data and passes it as a parameter to `path`, which
passes it as a named parameter (`path=somepath/tosomething`) to a `GET` request. A user may not
like percentencode for their use case and may prefer to use `lenprefixencode` to
have the multiaddr instead look like
`/dns4/example.com/http/GET/path/lenprefixencode/20_somepath/tosomething`. This
would work the same and require no changes to the `path` or `GET` component.
It's important to note that the binary representation of the data in
`percentencode` and `lenprefixencode` would be the same. The only difference is
how it appears in the human-readable representation.

## Use cases

- TODO: unpack the shortcomings of URLs
  - example: hostnames in https://
    - can't sidestep DNS
    - can't use different SNI vs. Host headers
    - can't do http-over-utp
    - TODO check out how http/1.1 vs. http/2 is distinguished
  - rift between filesystem, web, and databases

- TODO: case study: domain fronting
- TODO: case study: tunnelling
- TODO: case study: http proxying
- TODO: case study: multi-hop circuit relay
- TODO: case study: protocol migrations (e.g. ip4/ip6, 4in6, 6in4)


### Encapsulation based on context

Although multiaddrs are self-describing, it's possible to further encapsulate them based on context.
For example in a web browser, it's obvious that, given a hostname, HTTP should be spoken.
The specifics of this HTTP connection are not important (except maybe the use of TLS),
and will be derived from the browser's capabilities and configuration.

1. example.com/index.html
2. /http/example.com/index.html
3. /tls/sni/example.com/http/example.com/index.html
4. /dns4/example.com/tcp/443/tls/sni/example.com/http/example.com/index.html
5. /ip4/1.2.3.4/tcp/443/tls/sni/example.com/http/example.com/index.html

The resulting layers of encapsulation reflect exactly
how the bidirectional stream between client and server is constructed.

Now you can imagine how based on the browser's configuration, the multiaddr might look different.
For example you could use HTTP proxying or SOCKS proxying, or use domain fronting to evade censorship.
This kind of proxying is of course possible without multiaddr,
but only with multiaddr do we have a way of consistently addressing these networking constructions.


## Specification

- Human-readable multiaddr: `(/<protoName string>/<value string>)+`
  - Example: `/ip4/127.0.0.1/udp/1234`
- Machine-readable multiaddr: `(<protoCode uvarint><value []byte>)+`
  - Same example: `0x4 0x7f 0x0 0x0 0x1 0x91 0x2 0x4 0xd2`
  - Values are usually length-prefixed with a uvarint

Multiaddr and all other multiformats use unsigned varints (uvarint).
Read more about it in [multiformats/unsigned-varint](https://github.com/multiformats/unsigned-varint).


### Encoding

TODO: specify the encoding (byte-array to string) procedure

### Decoding

TODO: specify the decoding (string to byte-array) procedure


## Protocols

See [protocols.csv](protocols.csv) for a list of protocol codes and names,
and [protocols/](protocols/) for specifications of the currently supported protocols.

TODO: most of these are way underspecified

- /ip4, /ip6
- /ipcidr
- /dns4, /dns6
- [/dnsaddr](protocols/DNSADDR.md)
- /tcp
- /udp
- /utp
- /tls
- /ws, /wss
- /ipfs
- /p2p-circuit
- /p2p-webrtc-star, /p2p-webrtc-direct
- /p2p-websocket-star
- /onion


## Implementations

- [js-multiaddr](https://github.com/multiformats/js-multiaddr) - stable
- [go-multiaddr](https://github.com/multiformats/go-multiaddr) - stable
  - [go-multiaddr-dns](https://github.com/multiformats/go-multiaddr-dns)
- [java-multiaddr](https://github.com/multiformats/java-multiaddr) - stable
- [haskell-multiaddr](https://github.com/MatrixAI/haskell-multiaddr) - stable
- [py-multiaddr](https://github.com/multiformats/py-multiaddr) - stable
- [rust-multiaddr](https://github.com/multiformats/rust-multiaddr) - stable
- [cs-multiaddress](https://github.com/multiformats/cs-multiaddress) - alpha
- [net-ipfs-core](https://github.com/richardschneider/net-ipfs-core) - stable
- [swift-multiaddr](https://github.com/lukereichold/swift-multiaddr) - stable
- [elixir-multiaddr](https://github.com/aratz-lasa/ex_multiaddr) - alpha
- `multiaddr` sub-module of Python module [multiformats](https://github.com/hashberg-io/multiformats) - alpha
- [dart-multiaddr](https://github.com/YogiLiu/dart_multiaddr) - alpha
- Kotlin
  - [kotlin-multiaddr](https://github.com/changjiashuai/kotlin-multiaddr) - stable
  - `multiaddr` part of Kotlin project [multiformat](https://github.com/erwin-kok/multiformat) - alpha

TODO: reconsider these alpha/beta/stable labels


## Contribute

Contributions welcome. Please check out [the issues](https://github.com/multiformats/multiaddr/issues).

Check out our [contributing document](https://github.com/multiformats/multiformats/blob/master/contributing.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to multiformats are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.


## License

This repository is only for documents. All of these are licensed under the [CC-BY-SA 3.0](https://ipfs.io/ipfs/QmVreNvKsQmQZ83T86cWSjPu2vR3yZHGPm5jnxFuunEB9u) license, © 2016 Protocol Labs Inc. Any code is under a [MIT](LICENSE) © 2016 Protocol Labs Inc.
