# Cron Jobs
Cron jobs allow you automate tasks at a specific time of the day. A user can specify any date, time or frequency to schedule a task. This is done by the crond service that runs in the background. 

To schedule a job, you need to run `crontab -e`. This will open up the crontab in VI editor. Go to the bottom of the file and add the configuration for the scheduled job. The first 5 fields are used to specify the exact schedule to run the task. The sixth field is the command to run. You can see an example below:

```
# m   h   dom   mon   dow   command
  0   21  *     *     *     uptime >> /tmp/system-report.txt
```

It is important not to use sudo with the crontab command. Doing this will result in the cron job being scheduled for the root user.

To schedule a job to run on Feb 19th at 8.10am, you would do the below;

```
m   h   dom   mon
10  8   29    2

```

If you want this to schedule only if 19th February is a Monday, you would add the below;

```
m   h   dom   mon   dow
10  8   29    2     1

```
You can set the field to * if you wish to set the job to occur in all of those occurances. For example, for every month, you would put month as *. For example, to schedule a job every minute of every day of every month, any weekday, you can do the below;

```
m   h   dom   mon   dow
*   *   *     *     *

```

If you want there to be a 2 min interval, you can do the below;

```
m    h   dom   mon   dow
*/2  *   *     *     *

```

To list cron jobs, you can run the below command;
`crontab -l`

You can also verify if the cron job has run by checking the /var/log/syslog file.