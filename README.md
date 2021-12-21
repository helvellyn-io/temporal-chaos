# chaos
**Owner**: Dylan@helvellyn.io
Temporal Chaos Experiments 
 
* Project using the Temporal (https://www.temporal.io/) framework to enable chaos experiments on a component at any level of the OSI stack. 

  - Targeted/random deleteion of pods (GA) 
  - Targeting/ranom deletion of deployments (GA) 
  - Targeted/random deletion of nodes
  - Inject latency for gRPC/HTTPS response
```
BRANCH DETAILS: 

Master -

This project randomly deletes pods of a specified deployment. 
It uses the Temporal platform to run a singletask to delete a pod. 

Cron -

This project randomly deletes pods of a specified deployment. 
It uses the Temporal platform to run the task.
The schedule is determined via the CRON specifications.

Chaos Executions DB

Running Redis locally for the moment to store execution history.
Will look at using Temporal for this but not 100% its the right place 
atm. 

```
```
RUNNING: 
go run ./worker/main.go
go run ./start/main.go

```
```
TODO:
write tests.
redis integration to log chaos experiments.

```
