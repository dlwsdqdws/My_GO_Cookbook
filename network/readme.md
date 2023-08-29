# Network

- [Network](#network)
  - [Network Access](#network-access)
    - [Routing](#routing)
    - [ARP Protocol](#arp-protocol)

## Network Access

### Routing

<p align="center"><img src="../static/img/network/access/routing.png" alt="RPC Process" width="500"/></p>

1. Routing does not have to be symmetrical.

2. Routing works at Layer 3 (Network Layer).

3. Routing is to change the mac address to find the packet sending port.

```go
func send_one_pkt(){
    rt = find_rt(dst)  // including host egress network port and next_hop
    ...
    l2->dst_mac = rt->next_hop->mac // change mac
    p = append(p, l2)
    ...
    send(p, rt->port) // Specify when sending
}
```

4. Dynamic Routing: BGP / OSPF

### ARP Protocol

<p align="center"><img src="../static/img/network/access/arp_process.jpg" alt="RPC Process" width="500"/></p>