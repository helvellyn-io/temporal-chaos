# chaos
Temporal Chaos Experiments 
 
* Project using the Temporal framework to enable chaos experiments on a component at any level of the OSI stack.

  - Targeted/random deleteion of pods
  - Targeting/ranom deletion of deployments
  - Targeted/random deletion of nodes
  - Inject latency for gRPC/HTTPS response
  - TCP/IP fault injection
  - DNS query latency injection.
  - Override VPA to reduce|increase replica count
  - Region blackholing
  - AZ blackholing
  - Vampiric CPU
  - more ... 
```
BRANCH DETAILS: 

Master -

This project randomly deletes pods of a specified deployment. 
It uses the Temporal platform to run a singletask to delete a pod. 

Cron -

This project randomly deletes pods of a specified deployment. 
It uses the Temporal platform to run the task.
The schedule is determined via the CRON specifications.

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
