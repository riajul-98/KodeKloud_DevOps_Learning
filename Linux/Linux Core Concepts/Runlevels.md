# Runlevels
Linux can run in multiple modes such as the graphical mode and the command line mode. To see the operation mode set in the system, run the `runlevel` command.

In systemd, runlevels are called targets, you can see the different types of runlevels below;

| Runlevel | systemd target    | Information                                                      |
|----------|-------------------|------------------------------------------------------------------|
| 0        | poweroff.target   | Shuts system down                                                |
| 1        | rescue.target     | To start single-user mode (good for troubleshooting/admin tasks) |
| 2        | multi-user.target | System starts with multi-user mode but without networking        |
| 3        | multi-user.target | With Multi-user mode with networking                             |
| 4        | multi-user.target | Reserved                                                         |
| 5        | graphical.target  | With GUI                                                         |
| 6        | reboot.target     | Reboots the system                                               |

## Viewing and Changing Systemd Target
- `systemctl get-default`: Displays the set systemd target.
- `systemctl set-default <desired target>`: Allows you to change the runlevel. e.g. `systemctl set-default multi-user.target`