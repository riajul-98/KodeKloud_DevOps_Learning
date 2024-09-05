# File Compression and Archival
 
- `du -sk <file_name>`: Displays the size of a file or directory in kilobytes.
- `du -sh <file_name>`: Displays the size of a file or directory in a human readable format.
- `ls -lh <file_name>`: Displays the size of a file or directory as well other key information.

## Archiving
Tar is used to group multiple files and directories into a single archive file. The below are tar commands:

- `tar -cf test.tar file1 file2 file3`: Archives file1, file2 and file 3 into a tarball called test.tar.
- `tar -tf test.tar`: Displays the contents of the tarball.
- `tar -xf test.tar`: Extracts the contents of the tarball.
- `tar -zcf test.tar file1 file2 file3`: Creates and compresses the tarball to reduce its size.

## Compressing
Compression is used to reduce the size consumed by a file or data set. The below are methods of compressing:

- `bzip2 test.img`
- `gzip test1.img`
- `xz test2.img`

The files can then be uncompressed using the below commands:

- `bunzip2 test.img.bz2`
- `gunzip test1.img.gz`
- `unxz test2.img.xz`

Compressed files do not always need to be uncompressed. There are tools which can read compressed files without uncompressing them such as:

- zcat
- bzcat
- xzcat