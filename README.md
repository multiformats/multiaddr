# multiaddr

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-multiformats-blue.svg?style=flat-square)](http://github.com/multiformats/multiformats)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)

> The network addresses multiformat

## Table of Contents

- [Table of Contents](#table-of-contents)
- [What is multiaddr?](#what-is-multiaddr)
  - [Binary format:](#binary-format)
  - [String format:](#string-format)
- [Implementations](#implementations)
- [Maintainer](#maintainer)
- [Contribute](#contribute)
- [License](#license)

## What is multiaddr?

Multiaddr is a standard way to represent addresses that: 
- Support any standard network protocols.
- Self-describe (include protocols).
- Have a binary packed format.
- Have a nice string representation.
- Encapsulate well.

## Specification

Normally, addresses have been represented using string addresses, like:

```
tcp4://127.0.0.1:1234
udp4://10.20.30.40:5060
ws://1.2.3.4:5678
tcp6://[1fff:0:a88:85a3::ac1f]:8001
```

This isn't optimal. Instead, addresses should be formatted so:

### Binary format:

```
(varint proto><n byte addr>)+
<1 byte ipv4 code><4 byte ipv4 addr><1 byte udp code><2 byte udp port>
<1 byte ipv6 code><16 byte ipv6 addr><1 byte tcp code><2 byte tcp port>
```

### String format:

```
(/<addr str code>/<addr str rep>)+
/ip4/<ipv4 str addr>/udp/<udp int port>
/ip6/<ipv6 str addr>/tcp/<tcp int port>
```

### Protocols

See [protocols.csv](protocols.csv).

Originally from here:
https://github.com/jbenet/random-ideas/issues/11

## Implementations

- [js-multiaddr](https://github.com/multiformats/js-multiaddr) - stable
- [go-multiaddr](https://github.com/multiformats/go-multiaddr) - stable
- [java-multiaddr](https://github.com/multiformats/java-multiaddr) - stable
- [hs-multiaddr](https://github.com/basile-henry/hs-multiaddr) - draft
- [py-multiaddr](https://github.com/sbuss/py-multiaddr) - alpha
- [rust-multiaddr](https://github.com/multiformats/rust-multiaddr) - draft
- [cs-multiaddress](https://github.com/tabrath/cs-multiaddress) - alpha

## Maintainer

Captain: [@jbenet](https://github.com/jbenet).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/multiformats/multiaddr/issues).

Check out our [contributing document](https://github.com/multiformats/multiformats/blob/master/contributing.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to multiformats are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

## License

[MIT](LICENSE) © Protocol Labs Inc.
