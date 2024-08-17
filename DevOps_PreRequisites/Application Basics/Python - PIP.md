# Python
Python is a free and open source, cross platform compatible programming language.

## Installation
Some operating systems come installed with Python2 or Python3. To check, you can run the `python2` or `python3` commands to verify this. To install Python, use the below commands;
`yum install python2 -y` or `yum install python36 -y`
To enter invoke the interpreter, you would run one of the below commands;
`python2` or `python3`
To exit, you type `exit()`

## Python Commands
- To check the version, type
`python2 -V` or `python3 -V`
- To run a Python application (main.py) run
`python2 main.py` or `python3 main.py`

# Python Package Manager (PIP)
When Python is installed, PIP is also installed alongside. As with Python, you can have two separate versions of PIP installed, one for version 2 and version 3.

## Commands
- `pip2 -V` or `pip3 -V`: Displays the version of pip.
- `pip -V`: Displays the version of PIP as well as the version of Python.
- `pip install flask`: Installs the specified package, in this case, Flask. The package would be installed under `/usr/lib/python{version-number}/site-packages` for 32-bit packages and `/usr/lib64/python{version-number}/site-packages` for 64 bit packages.
- `pip show flash`: Provides details on the package including the location of the package.
- `python2 -c "import sys; print(sys.path)"`: Imports package into code.

## Requirements
Large applications require multiple packages. Rather than running a command for each dependancy, you can run them all at once like below;
`pip install flask jinja2 markupsafe`
But a better approach would be to use a requirements.txt file which contains the names of each package. You would then run the below command;
`pip install -r requirements.txt`

You can also include the version of the packages in this file. An example requirements.txt file can be seen below
```

Flask==0.10.1
Jinja2==2.7.3
MarkupSafe==0.23
Werkzeug==0.9.6
requests==2.3.0
gunicorn==18.0

```

## Upgrade/Uninstall Package
- `pip install flask --upgrade`: Upgrades package to the latest version.
- `pip uninstall flask`: Uninstalls package.