# `percentencode`

This encoding escapes `/` and `%` in string representations of a multiaddr.
Binary representations are unaffected since they are encoded as a length prefix
byte array. `/` is encoded as `%2f` and `%` is encoded as `%25`. These are their
ascii values as hex. Implementations should accept both upper and lower case hex
values.

This encoding allows a user to represent a path in a multiaddr. e.g.
`http://example.com/something/nested/here` could be represented as
`/dns4/example.com/tcp/80/http/GET/path/percentencode/something%2Fnested%2Fhere`
(note `path` is not a multiaddr component yet).
