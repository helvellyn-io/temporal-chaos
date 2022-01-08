# chaos
**Owner**: dylan@helvellyn.io

Temporal Chaos Experiments 

* Project using the Temporal (https://www.temporal.io/) framework to enable chaos experiments on a component at any level of the OSI stack. 

  - Targeted/random deleteion of pods (SAFE)
  - Corden/Drain node(s) (SAFE)
  - Targeting/ranom deletion of deployments (UNSAFE UNLESS USING OPERATORS)
  - Targeted/random deletion of nodes (UNSAFE WITHOUT CLUSTER AUTOSCALE)
  - Inject latency for gRPC/HTTPS response (SAFE)
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
