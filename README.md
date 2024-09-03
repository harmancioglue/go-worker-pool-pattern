# Go Workerpool

This repository aims to show how worker pool simply works in Go.

## Tech

- Go

## Intro

![image](![image](https://github.com/user-attachments/assets/6fde1377-6c07-4854-82c0-e56db0a5a5d4))

Worker pools is used to achieve concurrency with fixed number of workers to execute multiple amount of tasks on a queue. 

This helps us to use resources (CPU, memory) efficiently. That's because only fixed number of go routines works and wait for the queue to be filled.

If there is any job in queue(channel), empty worker will process the job and then will be waiting for the next job.
