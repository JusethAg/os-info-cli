# os-info-cli
Minimalist CLI for getting operating system information.

## Usage
```sh
./os-info-cli [OPTIONS]

    [OPTIONS]:
        -h, --help                  Usage message
        -a, --all                   Show all info (CPU, Network and memory)
        -f, --filter=filter         (cpu | net | mem)
```



## Internals
I'm building this mini project because I started learning Go. I decided to write a simple CLI for getting the most common information that I look for in my operating system. 

Mainly the purpose of this is:

1. See how complex or easy is to get info from the OS using Go.
2. See how easy is to understand Go's documentation
3. Learn how to work with concurrency and processes in Go.

In case you want to see the goal/scope of this project, I suggest reviewing the [To-Do section](#to-do-) 
## Run tests

## To-Do
- [ ] Create the basic structure of the project (Considering best practices for structuring a Go project and adding dependencies for adding unit tests).
- [X] Create a workflow with GitHub actions
- [X] Accept flags for getting information of the OS (help, all, filter by cpu, network, or memory).
- [X] Get network info (public and private IPs).
- [ ] Get CPU info (usage % by user, by system, IDLE, top 5 processes, temperature).
- [ ] Get memory info (% wired, active, compressed, and free, top 5 processes).