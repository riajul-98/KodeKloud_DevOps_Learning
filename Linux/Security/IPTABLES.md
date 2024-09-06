# IP Tables Introduction
To run SSH and SCP commands, the target machines port 22 should be open from the client to the server. Network security is important to allow authorised users to have access and stop unauthorised users gaining access to the servers. This can be achieved through measures such as external firewalls or appliances such as CiscoASA. These appliances allow you to define rules to allow or block any traffic coming into your server. You can also use tools such as IPtables and firewalld to do this at an individual server level. 

In this example, the client server has IP 172.16.238.187, the application server (DEVAPP01) with IP 172.16.238.10 and database server (DEVDB01) with IP 172.16.238.11. Without a firewall configured, every server can connect to every other server in the network. To improve security, we want to restrict network access. The client server should only be able to access the application server on port 80. The application server should not be accessible from any other servers such as DEVDB01. The database server runs the DB on port 5432 which should only be accessible by the application server and not the client server or any other servers. The application server should also have HTTP (port 80) access to the software repo server (172.16.238.15). We also want to block outgoing internet access from the application server. 

We can do this by using the below;
- `sudo apt install iptables -y`: Installs IP Tables.
- `sudo iptables -L`: Lists the default rules. There are three types of chain rules; 
    - Input: Network coming into the system.
    - Forward: Data is forwarded to other devices in the network. Usually this rule is for routers.
    - Output: Network going out of the system. initiated by this server
The default policy is set to accept, meaning all traffic is allowed to go in and out of the system. These are called chain rules as there may be multiple rules before the network is accepted into the system.

To add an incoming input rule to the application server to allow an SSH connection from the client laptop, use the below command;
- `iptables -A INPUT -p tcp -s 172.16.238.187 --dport 22 -j ACCEPT`

| Option   | Description                      |
|----------|----------------------------------|
| -A       | Add Rule                         |
| -p       | Protocol                         |
| -s       | Source                           |
| -d       | Destination                      |
| --dport  | Destination port                 |
| -j       | Action to take                   |
| -I       | Add rule to top of the chain     |
| -D       | Deletes a rule                   |

We would also need to add a reject rule for all other servers. We can do this with the below command;
- `iptables -A INPUT -p tcp --dport 22 -j DROP`

The order of the rules added is important as IP table rules are implemented from top to bottom. If multiple rules match a source or destination and ports, the first rule will be applied and others will be ignored.

To add an output rule, we use the below command;
- `iptables -A OUTPUT -p tcp -d 172.16.238.15 --dport 80 -j ACCEPT`

To add a rule which precedes a previous rule, we can use the following command;
- `iptables -I OUTPUT -p tcp -d 172.16.238.100 --dport 443 -j ACCEPT`

To delete a rule, we can use the following command;
- `iptables -D OUTPUT 5` or `iptables -D INPUT 5` where the number corresponds to the rule number.