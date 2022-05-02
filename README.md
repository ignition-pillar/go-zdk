# go-zdk

An SDK for the Network of Momentum built on top of [go-zenon](https://github.com/zenon-network/go-zenon)

See `examples` to get started.

To build the examples:

```shell
make examples
```

## Features

- Supports all JSON-RPC methods including Subscriptions
- Supports all embedded Contract methods

## Todo

- [ ] Event and filter system for better Subscriptions Api
- [ ] ZtsAmount (decimal) support for financial calculations
- [ ] Documentation & Versioning Scheme
- [ ] Testing framework
- [ ] Support PoWCallback
- [ ] Evaluate pointers for non-struct return types e.g `[]PillarInfo` vs `[]*PillarInfo`
- [ ] (?) Add logging
- [ ] (?) Repackage types from go-zenon for convenient imports
- [ ] (?) Provide additional IClient implementations e.g using in-memory or distributed caches
