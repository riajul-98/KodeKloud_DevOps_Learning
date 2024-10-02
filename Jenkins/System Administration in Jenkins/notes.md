# Backup and restoring Jenkins
Jenkins has multiple types of backups such as full backups using plugins and snapshots as well as writing a shell script which backs up the Jenkins Instance. It is most crucial to backup the JENKINS_HOME directory as this contains the configuration files and Jenkin jobs.

According to the video, the main backup plugin used is ThinBackup. After installing the plugin, go into Manage Jenkins > ThinBackup > backup now.
In the ThinBackup settings, you can choose the directory to backup to, the schedule and some more different settings. For the backup directory, the Jenkins user will need read, write access.

To restore the backup, you would go into Manage Jenkins > ThinBackup > Restore. You can choose the timestamp where you can restore from.

# Monitoring
There are many ways to monitor Jenkins such as through DataDog, JavaMelody and Prometheus. These are done through plugins. These can be useful to scrape data from the Jenkins server such as successful build counts.