# Definition files
When you have created a definition file and wish to create the pod/cluster, run the `kubectl create -f <file_name>` command.

When you have created any modifications to the definition files, use the `kubectl apply -f <file_name>` command or `kubectl edit pod <file_name>`.

To view pods created, use the `kubectl get pods` command.
To view replication controllers created, use the `kubectl get replicationcontroller` command.
To view replica sets created, use the `kubectl get replicaset` command. 
To view deployments created, use the `kubectl get deployments` command. 
To view all created objects, use the `kubectl get all` command.

To scale a replica set, update the definition file with the number of replicas you desire and run the `kubectl replace -f <file_name>` command or use the `kubectl scale --replicas=<number_of_replicas> <file_name>` command. 

## Updates
- `kubectl rollout status deployment<deployment_name>`: Shows status of update rollout.
- `kubectl rollout history deployment<deployment_name>`: Displays rollout history.

## Types of Deployment strategies
- Recreate: Destroy all running instances and bringing up the updated instances (not recommended)
- Rolling Update: Take down the older version and bring up the newer version, one by one.

## How to update
- Make changes to the definition file and then run the `kubectl apply -f <file_name>` command.
- Or you can use the `kubectl set image deployment/<deployment_name> <change_name>=<change>`. For example, 
`kubectl set image deployment/myapp-deployment nginx-container=nginx:1.9.1`

## Rollback
`kubectl rollout undo deployment/<deployment_name>`

## Cause of change
To instruct kubernetes to record the cause of change when creating a deployment, use the following command;
`kubectl create -f <file_name> --record`