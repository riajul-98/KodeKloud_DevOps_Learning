# The below is given as a challenge to find the error causing the script to fail and fix the issue.

# This script creates a backup of a given file by creating a copy as bkp
# For example some-file is backed up as some-file_bkp

# file_name="create-and-launch-rocket"

# cp $file_name $file_name_bkp

file_name="create-and-launch-rocket"

cp $file_name ${file_name}_bkp