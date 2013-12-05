# Sesame

[![Build Status](https://travis-ci.org/plants/sesame.png?branch=master)](https://travis-ci.org/plants/sesame)

A Go library for implementing the distributed password storage technique
described in [A Better Way to Store Password Hashes][1] ([Part 2][2])

Left to do:

 - salt2/hash2 from second article above
 - storage backends
 - a server/client (probably using [go-discover][3] and some sort of JSON RPC)

And if you know anything about cryptography, please do have a look. I would
really appreciate more eyes on this code.

[1]: http://www.opine.me/a-better-way-to-store-password-hashes/ "A Better Way to Store Password Hashes"
[2]: http://www.opine.me/all-your-hashes-arent-belong-to-us/ "Concluding: A Better Way to Store Password Hashes"
[3]: https://github.com/flynn/go-discover "flynn/go-discover"


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/plants/sesame/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

