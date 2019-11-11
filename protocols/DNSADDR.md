# `dnsaddr`

`dnsaddr` is a protocol that instructs the resolver to lookup multiaddr(s) in DNS TXT records for the domain name in it's value section.

To resolve a `dnsaddr` multiaddr, the domain name in the value section must first be prefixed with `_dnsaddr.`. Then a DNS query to lookup TXT records for the domain must be made. There may be multiple DNS TXT records for the domain. Valid `dnsaddr` TXT records begin with `dnsaddr=`, followed by a single multiaddr. Recursive lookups are allowed.

Example:

`/dnsaddr/bootstrap.libp2p.io` would result in a DNS TXT record query for `_dnsaddr.bootstrap.libp2p.io`.

```console
# TXT records for `_dnsaddr.bootstrap.libp2p.io`:
dnsaddr=/dnsaddr/sjc-1.bootstrap.libp2p.io/tcp/4001/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
dnsaddr=/dnsaddr/ams-2.bootstrap.libp2p.io/tcp/4001/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
# ...
```

...which further resolve to:

```console
# TXT records for `_dnsaddr.sjc-1.bootstrap.libp2p.io`:
dnsaddr=/ip6/2604:1380:1000:6000::1/tcp/4001/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
dnsaddr=/ip4/147.75.69.143/tcp/4001/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN

# TXT records for `_dnsaddr.ams-2.bootstrap.libp2p.io`:
dnsaddr=/ip4/147.75.83.83/tcp/4001/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
dnsaddr=/ip6/2604:1380:2000:7a00::1/tcp/4001/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
```

So, `/dnsaddr/bootstrap.libp2p.io` resolves to (at least) four multiaddrs:

```console
/ip6/2604:1380:1000:6000::1/tcp/4001/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN
/ip4/147.75.69.143/tcp/4001/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN

/ip4/147.75.83.83/tcp/4001/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
/ip6/2604:1380:2000:7a00::1/tcp/4001/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb
```
