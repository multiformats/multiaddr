# /http and /ws multiaddrs

- /http
  - /https vs. /tls/http
  - encapsulating another multiaddr
    - with delimiter: /httpe/$ipfs.io/chat$/ws/ipfs/Qmfoobar
    - or with escaping? /http/ipfs.io\/chat/ws/ipfs/Qmfoobar
    - both need change to how multiaddrs are being parsed
      - the escaping variant keeps split-on-slash relatively intact,
        and is maybe easier to parse than the delimiter variant
    - the binary form doesn't have to deal with this, it's simply length-delimited
  - http proxies
    - would likely be a /socks5 multiaddr that encapsulates /http
  - http auth
    - just simply: /http/user:pass@example.net
  - page fragment
    - just simply: /http/example.net#foo
