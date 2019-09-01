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
The specifics of this HTTP connection are not important (expect maybe the use of TLS),
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

### Codecs

Depending on the protocol of a Multiaddr component, different algorithms are used to
convert their values from/to binary representation. The name of the codec to use
for each protocol is noted in [protocols.csv](protocols.csv).

In general empty values in the string representation are always disallowed unless
explicitely noted otherwise. In case of conversion errors implementation must
refuse to process the given string/binary value and report the error to the caller
instead.

Depending on the codec type codecs may either be encoded using the standard variable
length encoding style, or into a specific static-length binary value without the
extra length information if this is noted in the respective codec's description.

All code examples are written in Python-based pseudo code and are optimized for
legibility rather than speed. In general you should always use existing libraries
and functions for performing the below conversions rather than rolling your own.

#### `fspath`

Encodes the given Unicode string using the system's local file system encoding.
On Windows this encoding likely being UTF-16, while being UTF-8 on most other
systems. It is up to the library to figure out the best encoding value for
these kinds of strings.

 * String → Binary: `str.encode(SYSTEM_FILESYSTEM_ENCODING)`
 * Binary → String: `bytes.decode(SYSTEM_FILESYSTEM_ENCODING)`

Protocols using the `fspath` encoding must not be shared between different hosts.

#### `idna`

Encodes the given Unicode representation according to IDNA-2008 ([RFC 5890](https://tools.ietf.org/html/rfc5890)) conventions using the [UTS-46](https://tools.ietf.org/html/rfc5890) input normalization and processing rules.

 * String → Binary:
    1. Normalize and validate the given input string according to [UTS-46 Section 4 (Processing)](https://www.unicode.org/reports/tr46/#Processing) and [UTS-46 Section 4.1 (Validity Criteria)](https://www.unicode.org/reports/tr46/#Validity_Criteria) with the following parameters:
       * UseSTD3ASCIIRules = true
       * CheckHyphens = true
       * CheckBidi = true
       * CheckJoiners = true
       * Transitional_Processing = false
    2. Convert the Unicode string to ASCII using the [UTS-46 Section 4.2 (ToASCII)](https://www.unicode.org/reports/tr46/#ToASCII) algorithm steps 2–6 with parameter *VerifyDnsLength* set to *true* and return the result.
 * Binary → String:  
   Convert the ASCII text string to Unicode according to the rules of [UTS-46 Section 4.3 (ToUnicode)](https://www.unicode.org/reports/tr46/#ToUnicode) using the same parameters as in step 1 of the *String → Binary* algorithm.

Examples of libraries for performing the above steps include the [Python idna](https://pypi.org/project/idna/) library.

#### `ip4`

Encodes an IPv4 address according to the conventional [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) first specificed in [RFC 3986 section 3.2.2 page 20 § 2](https://tools.ietf.org/html/rfc3986#page-20).

Protocols using this codec must encode it as binary value of exactly 4 bytes without
an extra length value.

 * String → Binary:
    1. Split the input string into parts at each dot (U+002E FULL STOP):  
       `sparts = str.split(".")`
    2. Assert that exactly 4 string parts were created by the split operation:  
       `assert len(parts) == 4`
    3. Convert each part from its ASCII base-10 number representation to an integer type, aborting if the conversion fails for any of the decimal string parts:  
       `octets = [int(p) for p in parts]`
    4. Validate that each part of the resulting integer list is in rage 0 – 255:  
       `assert all(i in range(0, 256) for i in octets)`
    4. Copy each of the resulting integers into a binary string of length 4 in network byte-order:  
       `return b"%c%c%c%c" % (octets[0], octets[1], octets[2], octets[3])`
 * Binary → String:
    1. Take the four bytes of the binary input and convert each to its equivalent base-10 ASCII representation without any leading zeros:  
       `octets = [str(binary[idx]) for idx in range(4)]`
    2. Concatinate resulting list of stringified octets using dots (U+002E FULL STOP):  
       `return ".".join(octets)`

Converting from string to binary addresses may be done using the POSIX
[`inet_addr`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/inet_addr.html)
function or the similar common Unix [`inet_aton`](https://man.cx/inet_aton(3))
function and its equivalent bindings in many other languages. Similarily the POSIX
[`inet_ntoa`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/inet_ntoa.html)
function available in many languages implements the previously mentioned binary
to string address transformation.

#### `ip6`

Encodes an IPv6 address according to the rules of [RFC 4291 section 2.2](https://tools.ietf.org/html/rfc4291#section-2.2) and [RFC 5962 section 4](https://tools.ietf.org/html/rfc5952#section-4).

Protocols using this codec must encode it as binary value of exactly 16 bytes without
an extra length value.

 * String → Binary:  
   Parse the given input address string according to the rules of [RFC 4291 section 2.2](https://tools.ietf.org/html/rfc4291#section-2.2) creating a 16-byte binary string. All textual variations (upper-/lower-casing, IPv4-mapped addresses, zero-compression, stripping of leading zeros) must be supported by the parser. Note that [scoped IPv6 addressed containing a zone identifier](https://tools.ietf.org/html/draft-ietf-ipngwg-scopedaddr-format-02) may not appear in the input string; external mechanisms may be used to encode the zone identifier separately through.
 * Binary → String:  
   Generate a canonical textual representation of the given binary input address according to rules of [RFC 5962 section 4](https://tools.ietf.org/html/rfc5952#section-4). Implementations must not produce any of the variations allowed by RFC 4291 mentioned above to ensure that all implementation produce a character by character identical string representation.

Converting between string to binary addresses should be done using the equivalent
of the POSIX [`inet_pton`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/inet_pton.html)
and [`inet_ntop`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/inet_ntop.html)
functions. Alternatively, using the BSD
[`getaddrinfo`/`freeaddrinfo`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/getaddrinfo.html)
and [`getnameinfo` with `NI_NUMERICHOST`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/getnameinfo.html)
may be a viable alternative for some environments.

### `onion`

Encodes a [TOR rendezvous version 2 service pointer](https://gitweb.torproject.org/torspec.git/tree/rend-spec-v2.txt?id=471af27b55ff3894551109b45848f2ce1002441b#n525) (aka .onion-address) and exposed service port on that system.

Protocols using this codec must encode it as binary value of exactly 12 bytes without
an extra length value.

 * String → Binary:
    1. Split the input string into 2 parts at the colon character (U+003A COLON):  
       `(service_str, port_str) = str.split(":")`
    2. Decode the *service* part before the colon using base32 into binary:  
       `service_bin = b32decode(service_str)`
    3. Convert the *port* part to a binary string as specified by the [`uint16be`](#uint16be) codec.
    4. Concatenate the service and port parts to obtain the final binary encoding:  
       `return service_bin + port_bin`
 * Binary → String:
    1. Split the binary value at the last two bytes into an service name and a port
       number:  
       `(service_bin, port_bin) = binary.split_at(-2)`
    2. Convert the service part into a base32 string:  
       `service_str = b32encode(service_bin)`
    3. Convert the *port* part to text as specified by the [`uint16be`](#uint16be) codec.
    4. Concatenate the result strings using a colon:  
       `return service_str + ":" + port_str`

### `p2p`

Encodes a libp2p node address.

TBD: Is this really always a base58btc encoded string of at least 5 characters in length!?


### `uint16be`

Encodes an unsigned 16-bit integer value (such as a port number) in network byte
order (big endian).

Protocols using this codec must encode it as binary value of exactly 2 bytes without
an extra length value.

 * String → Binary:
    1. Parse the input string as base-10 integer:  
       `integer = int(str, 10)`
    2. Verify that the integer is in a valid range for a positive 16-bit integer:  
       `assert integer in range(65536)`
    3. Convert the integer to a 2-byte long big endian binary string:  
       `return b"%c%c" % ((integer >> 8) & 0xFF, integer & 0xFF)`
 * Binary → String:
    1. Convert the two input bytes to a native integer:  
       `integer = port_bin[0] << 8 | port_bin[1]`
    2. Generate a base-10 string representation from this integer:  
       `return str(integer, 10)`

POSIX/BSD provides [`strtoul`](https://en.cppreference.com/w/c/string/byte/strtoul)
and [`htons`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/htons.html)
for the string to binary conversion and
[`ntohs`](https://pubs.opengroup.org/onlinepubs/9699919799/functions/ntohs.html)
and [`snprintf`](https://en.cppreference.com/w/c/io/snprintf) for the performing
the inverse operation.

## Protocols

See [protocols.csv](protocols.csv) for a list of protocol codes and names,
and [protocols/](protocols/) for specifications of the currently supported protocols.

TODO: most of these are way underspecified

- /ip4, /ip6
- /dns4, /dns6
- /dnsaddr
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
  - [go-multiaddr-net](https://github.com/multiformats/go-multiaddr-net)
- [java-multiaddr](https://github.com/multiformats/java-multiaddr) - stable
- [haskell-multiaddr](https://github.com/MatrixAI/haskell-multiaddr) - stable
- [py-multiaddr](https://github.com/multiformats/py-multiaddr) - stable
- [rust-multiaddr](https://github.com/multiformats/rust-multiaddr) - beta
- [cs-multiaddress](https://github.com/multiformats/cs-multiaddress) - alpha
- [net-ipfs-core](https://github.com/richardschneider/net-ipfs-core) - stable
- [swift-multiaddr](https://github.com/lukereichold/swift-multiaddr) - stable

TODO: reconsider these alpha/beta/stable labels


## Contribute

Contributions welcome. Please check out [the issues](https://github.com/multiformats/multiaddr/issues).

Check out our [contributing document](https://github.com/multiformats/multiformats/blob/master/contributing.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to multiformats are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.


## License

This repository is only for documents. All of these are licensed under the [CC-BY-SA 4.0](https://ipfs.io/ipfs/QmVreNvKsQmQZ83T86cWSjPu2vR3yZHGPm5jnxFuunEB9u) license, © 2016 Protocol Labs Inc, © 2019 Alexander Schlarb. Any code is under a [MIT](LICENSE) © 2016 Protocol Labs Inc, © 2019 Alexander Schlarb.
