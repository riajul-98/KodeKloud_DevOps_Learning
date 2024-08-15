# Package Management
Package managers allow you to install various softwares on the Linux system.

## RPM (Red Hat Package Manager)
RPM only installs the specified packages but does not install any dependancies.
- `rpm -i telnet.rpm`: Installs the specified package.
- `rpm -e telnet`: Uninstalls specified package.
- `rpm -q telnet`: Queries specified package.

## YUM
Yum installs the specified package as well as any dependancies needed by that package. Yum knows what is needed by checking the /etc/yum.repos.d directory.
- `yum install ansible`: Installs specified package as well as any dependancies.
- `yum repolist`: Displays the available repositories on the system.
- `yum list`: Displays installed packages. You can follow this by the package name if there is a specific package you are looking for.
- `yum remove ansible`: Uninstalls the specified package.
- `yum --showduplicates list ansible`: Displays all available versions of the specified package.
- `yum install ansible-2.4.2.0`: Adding a hyphen followed by a version number allows you to install a specific version of a package.