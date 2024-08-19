# IP Addresses
Devices use different kinds of interfaces or adapters such as ethernet interfaces and wireless interfaces. When the device connects to a network, the interface is assigned an IP address. To view the IP addresses on CentOS systems, run the below command

`ip addr show`

## Loopback address
The loopback address is the IP address which the device refers to itself as. It is usually 127.0.0.1. You can use this to access applications located on the device or use http://localhost and specify the port number.

# Ports
Ports are communication endpoints which allow programmes to listen on these ports for requests. If you have multiple IP addresses, you cannot communicate with these programmes on any IP address as long as you enter the port number. This behaviour can be changed by adding the IP address you wish the port to be listened on in the configuration file of the application under "host". If you want to use all the IP addresses to access this application, you can set the host to 0.0.0.0 in the configuration file. If you wish for no one else to access the application apart from the device where the application is being hosted, you can use the loopback address.
