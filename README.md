# go-zdk

An SDK for the Network of Momentum built on top of [go-zenon](https://github.com/zenon-network/go-zenon)

See `examples` to get started.

To build the examples:

```shell
make examples
```

## Features

- Supports all JSON-RPC methods (except `ledger.publishRawTransaction`)

## Todo

- [ ] ledger.publishRawTransaction, AccountBlockTemplate, contract methods, wallet integration, decimal support
- [ ] Embedded.Accelerator Api (upon release of znnd v0.0.3)
- [ ] (?) Repackage types from go-zenon for convenient imports
- [ ] Event and filter system for better Subscriptions Api
- [ ] Evaluate pointers for non-struct return types e.g `[]PillarInfo` vs `[]*PillarInfo`
- [ ] Documentation & Versioning Scheme
- [ ] Testing framework
- [ ] (?) Provide additional IClient implementations e.g using in-memory or distributed caches
