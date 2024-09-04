# RPM and YUM

## RPM (Red Hat Package Manager)
RPM only installs the specified packages but does not install any dependancies.
- `rpm -i telnet.rpm`: Installs the specified package.
- `rpm -e telnet`: Uninstalls specified package.
- `rpm -Uvh telnet`: Upgrades specified package.
- `rpm -q telnet`: Queries specified package.
- `rpm -Vf <path to file>`: Verifies a package.

## YUM
Yum installs the specified package as well as any dependancies needed by that package. Yum knows what is needed by checking the /etc/yum.repos.d directory.
- `yum install ansible`: Installs specified package as well as any dependancies.
- `yum repolist`: Displays the available repositories on the system.
- `yum provides <command>`: Searches which package is needed for a specific command to work.
- `yum list`: Displays installed packages. You can follow this by the package name if there is a specific package you are looking for.
- `yum remove ansible`: Uninstalls the specified package.
- `yum --showduplicates list ansible`: Displays all available versions of the specified package.
- `yum install ansible-2.4.2.0`: Adding a hyphen followed by a version number allows you to install a specific version of a package.
- `yum update ansible`: Updates the specified package.
- `yum update`: Checks which installed packages need to be updated and updates them.