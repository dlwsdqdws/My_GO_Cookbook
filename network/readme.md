# Network

- [Network](#network)
  - [Network Access](#network-access)
    - [Routing](#routing)
    - [Address Resolution Protocol (ARP)](#address-resolution-protocol-arp)
    - [Internet Protocol (IP)](#internet-protocol-ip)
    - [Network Address Translation (NAT)](#network-address-translation-nat)
  - [Network Transmission](#network-transmission)

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

<p align="center"><img src="../static/img/network/access/ip.png" alt="RPC Process" width="500"/></p>

1. The IP protocol is used to unify different types of Layer 2 network protocols.

2. IPV4 / IPV6

### Network Address Translation (NAT)

<p align="center"><img src="../static/img/network/access/nat.png" alt="RPC Process" width="500"/></p>

1. NAT works by mapping private IP addresses within the local network to one or more public IP addresses. This allows multiple devices in the local network to communicate with the external internet using a single public IP address. When a packet is sent from a device within the local network to the external network, NAT replaces the source IP address with the public IP address of the NAT device. When response packets return, NAT uses its mapping table to convert the destination IP address back to the private IP address of the source device.

2. Different types:

- Static NAT: One-to-one mapping, where a private IP address is mapped to a single public IP address.
- Dynamic NAT: One-to-many mapping, where multiple private IP addresses are mapped to a pool of public IP addresses, but the mappings are temporary and may change with each use.
- PAT (Port Address Translation): One-to-many mapping that uses different port numbers to distinguish between different internal devices, in addition to mapping private IP addresses to a public IP address.

3. NAT allows multiple devices to share the same public IP address under limited availability of public IP address resources, thereby enhancing the utilization of IPv4 addresses. However, NAT can introduce certain networking issues, such as not being suitable for certain applications and protocols, and potentially adding complexity to the network.

## Network Transmission

