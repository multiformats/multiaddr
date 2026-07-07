# ech

This protocol encodes a server's ECHConfigList as defined in [RFC
9849](https://www.rfc-editor.org/info/rfc9849/).

Its binary representation is the length prefixed bytes of the ECHConfigList.
Its string representation is the multibase encoding bytes of the ECHConfigList.

## Usage

`/ech` should be appended directly after `/tls` or `/quic-v1`.
