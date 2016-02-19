# multiaddr - network addresses format


## Implementations

- [js-multiaddr](https://github.com/jbenet/js-multiaddr) - stable
- [go-multiaddr](https://github.com/jbenet/go-multiaddr) - stable

## What is multiaddr?

A standard way to represent addresses that

- support any standard network protocols
- self-describe (include protocols)
- have a binary packed format
- have a nice string representation
- encapsulate well

Unsure at this point of the existance of such a spec. The closest I've seen is the string representations like:

```
tcp4://127.0.0.1:1234
udp4://10.20.30.40:5060
ws://1.2.3.4:5678
tcp6://[1fff:0:a88:85a3::ac1f]:8001
```

Instead, I want something like:

### Binary format:

```
(<1 byte proto><n byte addr>)+
<1 byte ipv4 code><4 byte ipv4 addr><1 byte udp code><2 byte udp port>
<1 byte ipv6 code><16 byte ipv6 addr><1 byte tcp code><2 byte tcp port>
```

### String format:

```
(/<addr str code>/<addr str rep>)+
/ip4/<ipv4 str addr>/udp/<udp int port>
/ip6/<ipv6 str addr>/tcp/<tcp int port>
```

Originally from here:
https://github.com/jbenet/random-ideas/issues/11

