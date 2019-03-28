Feature: Multiaddr

  Scenario: Simple string form
    Given the multiaddr /ip4/192.0.2.42/tcp/443
    Then the packed form is 0x04c000022a0601bb
    And the packed size is 8 bytes
    And the components are:
      | string          | stringSize | packed       | packedSize | value      | valueSize | protocol | codec | uvarint | lengthPrefix | rawValue   |
      | /ip4/192.0.2.42 | 15         | 0x04c000022a | 5          | 192.0.2.42 | 4         | ip4      | 4     | 0x04    |              | 0xc000022a |
      | /tcp/443        | 8          | 0x0601bb     | 3          | 443        | 2         | tcp      | 6     | 0x06    |              | 0x01bb     |

  Scenario: Simple packed form
    Given the multiaddr 0x04c000022a0601bb
    Then the string form is /ip4/192.0.2.42/tcp/443
    And the string size is 23 bytes
    And the components are:
      | string          | stringSize | packed       | packedSize | value      | valueSize | protocol | codec | uvarint | lengthPrefix | rawValue   |
      | /ip4/192.0.2.42 | 15         | 0x04c000022a |       5    | 192.0.2.42 | 4         | ip4      | 4     | 0x04    |              | 0xc000022a |
      | /tcp/443        | 8          | 0x0601bb     |       3    | 443        | 2         | tcp      | 6     | 0x06    |              | 0x01bb     |
