# Network

- [Network](#network)
  - [Network Access](#network-access)
    - [Routing](#routing)
    - [Address Resolution Protocol (ARP)](#address-resolution-protocol-arp)
    - [Internet Protocol (IP)](#internet-protocol-ip)

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

### Address Resolution Protocol (ARP)

<p align="center"><img src="../static/img/network/access/arp_process.jpg" alt="RPC Process" width="500"/></p>

1. ARP is to look up the mac address of the next hop.
   
2. Only devices within the same network segment can send ARP.
   
3. ARP requests are broadcasted, while ARP replies are unicast.
   
4. Gratuitous ARP: an ARP message in computer networking where a host sends an ARP request with its own IP address as the target IP address.

- Address Conflict Detection: A device can send a gratuitous ARP to check if another device is using its IP address. If it receives a response, it indicates an IP address conflict.
- Updating ARP Caches: Gratuitous ARPs can help update the ARP caches of other devices on the same network segment, ensuring they have the latest MAC address associated with the IP address.
- Failover and Redundancy: In case of a failover or a change in network configuration, a device might send a gratuitous ARP to inform other devices of its new MAC address.
- Network Reconfiguration: After a network configuration change (like subnet change), devices can send gratuitous ARPs to inform others about the new network setup.

5. ARP Proxy: a mechanism designed to solve ARP resolution challenges when communication spans different subnets or networks. It acts as a proxy for handling ARP requests and replies, enabling communication across subnet boundaries.

### Internet Protocol (IP)