# chaos
Temporal Chaos Experiments 


Define and schedule chaos experients. 

```
Branch: Master
This project randomly deletes pods of a specified deployment. 
It uses the Temporal platform to run a singletask to delete a pod. 

Branch: Cron

This project randomly deletes pods of a specified deployment. 
It uses the Temporal platform to run the task.
The schedule is determined via the CRON specifications.

Chaos Executions DB

Running Redis locally for the moment to store execution history.
Will look at using Temporal for this but not 100% its the right place 
atm. 

```

```
TODO

write tests

```