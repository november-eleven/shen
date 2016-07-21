# Shen

## Requirements

 * **go v1.6.2**
 * **node v6.3.0**

## Compile

```bash
make
```

## Launch

```bash
cd dist && ./shen
```

## Under the hood

 * Peer signaling is handled by a registration system inspired by bittorrent tracker system.
 * Shen is exclusively a REST API: there is no WebSocket or Server Side Event.
 * Shen is designed to be stateless and data-loss tolerant.
 * Shen use HTTP 1.1 so any client, regardless of its network topology or vendoring, can access this platform.
 * The asynchronous mecanism of registration and peers exchange enable the client to release the connection quite
   often: maintaining a network connection is no longer required.
 * Shen use an in-memory database to reduce request latency by network or filesystem read and write access.
 * Shen use a concurrent access lock-mecanism for its in-memory database: Read and Write try to be as fast as possible.
 * Scaling-up is not an issue if a node discovery mecanism is implemented: a peer will try to communicate with
   the same node, or try others as fallback if it fails...

