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

1. Initiate an empty string
2. Check if there is input data left. If there is, read a unsigned varint to get the protocol code, else go to step 9.
3. Check if protocol exists in the protocol list, else stop encoding and throw error.
4. Get the name of the protocol from the protocol list and append it to the string of step 1, prefixed by a `/`.
5. Get the size of the protocol value from the protocol list. If it is zero, go to step 2.
6. If the size is fixed, read that amount of bytes, encode it using the procedures defined by the protocol and append it to the string of step 1, prefixed by a `/`. Then go to step 2.
7. If the size is variable, read the size of the data. Usually, this is another unsigned varint, but protocols can decide otherwise.
8. Read that amount of bytes from the remaining part of the data, encode it using the procedures defined by the protocol and append it to the string of step 1, prefixed by a `/`. Then go to step 2.
9. Return the string from step 1.

### Decoding

1. Initiate an empty byte-array
2. Get all segments that are prefixed by a `/` character.
3. Start with the first segment.
4. Check if the current segment contains a protocol name from the protocol list, else stop decoding an throw error.
5. Get the code of the protocol, decode it to a unsigned varint and append it to the byte-array of step 1.
6. If the size of that protocol is zero, go to the next segment. If that segment exist, go to step 4, else go to step 10.
7. If the protocol size is variable and/or non-zero, then go to the next segment.
8. Decode this segment using the procedures defined by the protocol and append it to the byte-array of step 1.
9. Go to the next segment. If that segment exists, go to step 4, else go to step 10.
10. Return the byte-array from step 1.


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
- [kotlin-multiaddr](https://github.com/changjiashuai/kotlin-multiaddr) - stable
- [go-multiaddr](https://github.com/multiformats/go-multiaddr) - stable
  - [go-multiaddr-dns](https://github.com/multiformats/go-multiaddr-dns)
- [java-multiaddr](https://github.com/multiformats/java-multiaddr) - stable
- [haskell-multiaddr](https://github.com/MatrixAI/haskell-multiaddr) - stable
- [py-multiaddr](https://github.com/multiformats/py-multiaddr) - stable
- [rust-multiaddr](https://github.com/multiformats/rust-multiaddr) - beta
- [cs-multiaddress](https://github.com/multiformats/cs-multiaddress) - alpha
- [net-ipfs-core](https://github.com/richardschneider/net-ipfs-core) - stable
- [swift-multiaddr](https://github.com/lukereichold/swift-multiaddr) - stable
- [elixir-multiaddr](https://github.com/aratz-lasa/ex_multiaddr) - alpha

TODO: reconsider these alpha/beta/stable labels


## Contribute

Contributions welcome. Please check out [the issues](https://github.com/multiformats/multiaddr/issues).

Check out our [contributing document](https://github.com/multiformats/multiformats/blob/master/contributing.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to multiformats are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.


## License

This repository is only for documents. All of these are licensed under the [CC-BY-SA 3.0](https://ipfs.io/ipfs/QmVreNvKsQmQZ83T86cWSjPu2vR3yZHGPm5jnxFuunEB9u) license, © 2016 Protocol Labs Inc. Any code is under a [MIT](LICENSE) © 2016 Protocol Labs Inc.
