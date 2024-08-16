# DNS

## Name Resolution
Say we have two devices on the same network, device A (192.168.1.10) and device B (192.168.1.11). You are able to ping the other device by running the ping command followed by the other devices IP address. For example;
`ping 192.168.1.11`
You can also use the hostname instead of the IP address like below;
`ping db`
You would need to first provide the IP address and the hostname in the /etc/hosts file before you can do this. This does not need to be the hostname device B knows itself as. You can check the hostname of the device by running the `hostname` command.

The /etc/hosts file can have as many names for as many servers as you want. The issue with this method is if there are too many servers, it can become difficult to keep the /etc/hosts file up to date on each device. This is where the DNS server comes in.

## DNS Server
Companies with large number of servers tend to use a separate DNS server to hold the IP addresses and hostnames rather than updating the IP addresses and hostnames on each server individually. The DNS servers IP address is added to the /etc/resolv.conf on each server and named as "nameserver".

You can still add IP addresses and hostnames to the /etc/hosts file if you want and Linux would prioritise this over looking at the DNS server. This behaviour can be changed by adding switching the order of the "hosts" entry in the /etc/nsswitch.conf file.

If you try ping a host which is not in either file or DNS server, the command will fail. To avoid this, you can add a public name server to the DNS server file such as 8.8.8.8.

To access an IP by just a subdomain, you make an entry into your /etc/resolv.conf file called search followed by the domain name which you wish to append like below;

```

nameserver  192.168.1.100
search      mycompany.com

```
Now if you want to ping web.mycompany.com, you can just use `ping web`. You can still use `ping web.mycompany.com` and the system will know not to append "mycompany.com" at the end.


## Domain Names
Domain names are names which can be translated to IP addresses on the public internet. Domain names are mae of several components;
- `.com`, `.net`, `.org`: Top level Domain
- `.`: Root
- `www`, `map`: Sub-domain, are in front of the domain name, separated by a '.'

## Other DNS Commands
- `nslookup www.google.com`: Queries a hostname from a DNS server. Does not consider entries in the /etc/hosts file, only the DNS server.
- `dig www.google.com`: Tests DNS name resolution. Provides more details than nslookup. Does not consider entries in the /etc/hosts file, only the DNS server.