# Go Workerpool

This repository aims to show how worker pool simply works in Go.

## Tech

- Go

## Intro

![image](https://github.com/harmancioglue/go-worker-pool-pattern/assets/27441734/93f76717-0518-4f09-ae01-4c8b1fcf2395)

Worker pools is used to achieve concurrency with fixed number of workers to execute multiple amount of tasks on a queue. 

This helps us to use resources (CPU, memory) efficiently. That's because only fixed number of go routines works and wait for the queue to be filled.

If there is any job in queue(channel), empty worker will process the job and then will be waiting for the next job.
