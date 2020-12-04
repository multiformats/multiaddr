# `dnsaddr`

`dnsaddr` is a protocol that instructs the resolver to lookup multiaddr(s) in DNS TXT records for the domain name in it's value section.

To resolve a `dnsaddr` multiaddr, the domain name in the value section must first be prefixed with `_dnsaddr.`. Then a DNS query to lookup TXT records for the domain must be made. There may be multiple DNS TXT records for the domain. Valid `dnsaddr` TXT records begin with `dnsaddr=`, followed by a single multiaddr. Recursive lookups are allowed.

## Example

`/dnsaddr/bootstrap.libp2p.io` would result in a DNS TXT record query for `_dnsaddr.bootstrap.libp2p.io`.

```sh
# TXT records for `_dnsaddr.bootstrap.libp2p.io`:
dnsaddr=/dnsaddr/sjc-1.bootstrap.libp2p.io/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
dnsaddr=/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
# ...
```

...which further resolve to:

```sh
# TXT records for `_dnsaddr.sjc-1.bootstrap.libp2p.io`:
dnsaddr=/ip6/2604:1380:1000:6000::1/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
dnsaddr=/ip4/147.75.69.143/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN

# TXT records for `_dnsaddr.ams-2.bootstrap.libp2p.io`:
dnsaddr=/ip4/147.75.83.83/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
dnsaddr=/ip6/2604:1380:2000:7a00::1/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
```

So, `/dnsaddr/bootstrap.libp2p.io` resolves to (at least) four multiaddrs:

```
/ip6/2604:1380:1000:6000::1/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
/ip4/147.75.69.143/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN

/ip4/147.75.83.83/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
/ip6/2604:1380:2000:7a00::1/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
```

## Suffix matching

A `/dnsaddr` may specify additional nested protocols that must match when resolving the TXT record. When the TXT records are resolved, items whose suffix doesn't match the provided one are dropped.

The [default IPFS bootstrap list](https://github.com/ipfs/go-ipfs-config/blob/v0.0.11/bootstrap_peers.go#L22-L25) contains 4 dnsaddrs for the domain `bootstrap.libp2p.io`, and each one specifies has a different Peer ID. The first one is:

```go
"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
```

To resolve that multiaddr first look up the txt records at `_dnsaddr.bootstrap.libp2p.io`

```console
$ dig +short txt _dnsaddr.bootstrap.libp2p.io
"dnsaddr=/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/4001/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb"
"dnsaddr=/dnsaddr/sjc-1.bootstrap.libp2p.io/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN"
"dnsaddr=/dnsaddr/nrt-1.bootstrap.libp2p.io/tcp/4001/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt"
"dnsaddr=/dnsaddr/ewr-1.bootstrap.libp2p.io/tcp/4001/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa"
```

Check each record to find ones where the suffix matches the one specified on the multiaddr: `/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN` which leaves us with just:

```
/dnsaddr/sjc-1.bootstrap.libp2p.io/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
```

In this case, it's another `/dnsaddr`, so take the same steps again, this time with the `_dnsaddr.sjc-1.bootstrap.libp2p.io` domain:

```console
$ dig +short txt _dnsaddr.sjc-1.bootstrap.libp2p.io
"dnsaddr=/ip4/147.75.69.143/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN"
"dnsaddr=/ip6/2604:1380:1000:6000::1/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN"
```

Now both records match the provided suffix of `/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN`, so our multiaddr finally resolves to:

```
/ip4/147.75.69.143/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
/ip6/2604:1380:1000:6000::1/tcp/4001/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
```

## Caveats

Some implementations may fail to resolve `/dnsaddr` addresses if the total size of all your published TXT records for a given domain exceed 512 bytes, as the initial dns response will be truncated to 512 bytes.

See: https://serverfault.com/questions/840241/do-dns-queries-always-travel-over-udp

For the default IPFS bootstrap list, we use recursive `/dnsaddr` resolution, as described above, so we only publish 4 txt records on the primary domain `bootstrap.libp2p.io`, which in turn resolve to a pair of `ip6` and `ip4` multiaddrs per bootstrap node. In that way the dns response at both steps fits within 512 bytes.
