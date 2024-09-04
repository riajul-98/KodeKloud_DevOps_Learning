# DPKG and APT

## DPKG
DPKG only installs the specified packages but does not install any dependancies.
- `dpkg -i telnet.deb`: Installs the specified package.
- `dpkg -r telnet`: Uninstalls specified package.
- `dpkg -l telnet`: Lists installed packages with specified name.
- `dpkg -s telnet`: Checks status of the package
- `dpkg -p <path to file>`: Verifies a package.

## APT / APT-GET
APT installs the specified package as well as any dependancies needed by that package. Yum knows what is needed by checking the /etc/apt/sources.list file.
- `apt install ansible`: Installs specified package as well as any dependancies.
- `apt list`: Displays the available packages.
- `apt remove ansible`: Uninstalls the specified package.
- `apt install ansible-2.4.2.0`: Adding a hyphen followed by a version number allows you to install a specific version of a package.
- `apt upgrade`: Used to install upgrades of all software packages.
- `apt update`: Refreshes repository.
- `apt edit-sources`: Updates repositories.

APT is more user-friendly than APT-GET. You can also search packages.