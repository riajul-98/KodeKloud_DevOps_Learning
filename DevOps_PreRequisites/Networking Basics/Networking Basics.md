# Networking Basics
Networking is crucial when working with Linux in an IT role due to the need for different servers to communicate for various reasons. This section covers the basics of switching, routing, default gatweay and DNS configuration on Linux

## Switches
For two devices to communicate, we need to connect these devices to a switch. The switch then creates a network between these devices. To connect these devices to a switch, a physical or virtual interface is needed on each host. To view the interfaces for the host, we use the following command;
`ip link`

The devices would need to be assigned IP addresses on the same network as the switch. For example, if the switch IP address was 192.168.1.0, we could use 192.168.1.10 for device A and 192.168.1.11 for device B. We would assign these to network interface eth0 by using the below commands;
`ip addr add 192.168.1.10/24 dev eth0` and `ip addr add 192.168.1.11/24 dev eth0`

Now they are on the same network and can communicate with each other. You can test this by using the `ping` command followed by the other devices IP address. For example;
`ping 192.168.1.11`

The switch can only allow communication in its own network, not cross-network. For cross-network communication, we need a router.

## Router
Say we have two networks, 192.168.1.0 and 192.168.2.0. How would we get these networks to communicate with each other? We would use a router. A router helps to connect multiple networks together. In our example, the router gets assigned two IPs, one from each network. For example, the router could be assigned 192.168.1.1 and 192.168.2.1. 

The router is now connected to the two networks and now device B wishes to send a packet to system C (192.168.2.10) on network 2, how will device B know where the router is on the network? The router is just another device on the network. This is where we configure a gateway, also known as a route.

## Gateway
The route / gateway tells the devices on a network about other devices on the same network. To display the existing routing configurations, run the `route` command. This will display the Kernel routing table.

To configure the gateway on device B to reach devices on network 2, run the following command;
`ip route add 192.168.2.0/24 via 192.168.1.1`
The IP range contains the IPs for devices in network 2 and the IP address is the IP address for the router known by the devices in network 1.

In order to connect to the internet, you would need to add a new route to the route table to route all traffic to the IP address of the website. Rather than adding all the IP addresses on the internet, you can use the router as the default gateway using the below command;
`ip route add default via 192.168.2.1` or `ip route add 0.0.0.0 via 192.168.2.1`
where 192.168.2.1 is the IP address of the router from network 2. This will route the request to the router and the router will route the traffic to the given website.

## Setting Up Linux Host as a Router
Say we have device A (192.168.1.5) on network 192.168.1.0 and device C (192.168.2.5) on network 192.168.2.0 which are both connected to device B. Device B has IP addresses 192.168.1.6 on interface eth0 and 192.168.2.6 on interface eth1. To get device A to communicate with device C, we need to tell the devices to use device B as a gateway to network 2 by adding a routing entry like the below;
`ip route add 192.168.2.0/24 via 192.168.1.6`
The same command with the correct IP addresses would need to be run on device C if you want device C to communicate with device A. By default, this will still not work as packets are not forwarded from one interface to the other for security reasons. To change this, you would need to configure the /proc/sys/net/ipv4/ip_forward. You would need to set this to 1. To make this change persistent after reboots, you would need to also change the net.ipv4.ip_forward to equal 1 in the /etc/sysctl.conf file.
