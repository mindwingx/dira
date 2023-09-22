# DIRA: Docker Remote Command Execution in Golang

DIRA is a Golang package that simplifies the execution of commands in Docker containers
via Docker Remote Access.

### Installation

Install DIRA using the following command:

``` go
go get github.com/mindwingx/dira
```
### Prerequisite
Please ensure that you have enabled <a href="https://docs.docker.com/config/daemon/remote-access" target="_blank">Docker Remote Access</a> on your machine.

Note: Enabling it with `0.0.0.0:<port>` is not recommended.

### Snapshot

Executing commands in Docker containers is fast and straightforward with DIRA. Here's an example of how to use it:

```go
const (
hostURLOrIP = "http://127.0.0.1:2375"
containerName = "ubuntu"
cmd           = "echo Hello World"
)

// Create a DIRA instance with the desired configuration.
dockerCmd := dira.SetOpts(hostURLOrIP, containerName, cmd)

// Execute the command in the Docker container.
result, err := dockerCmd.Exec()
if err != nil {
fmt.Printf("Error: %v\n", err)
return
}

fmt.Printf("Command Output: %s\n", result)
```

### Usecase
Imagine needing to perform operations like file encryption, media compression, or OS layer executions within a Microservice process, but not directly executable. These tasks might run on a VPS's localhost or a remote server.

Handling container interactions and complex commands can complicate the process. DIRA simplifies this by enabling communication between containers through Docker's REST API. While other solutions exist, DIRA offers a developer-friendly approach, especially in remote access scenarios.

### Example

- If you find that the related process is taking longer, you can increase the timeout by specifying the duration in seconds.
```go
dockerCmd := dira.SetOpts(hostURLOrIP, containerName, cmd)

// Execute the command in the Docker container.
result, err := dockerCmd.
            SetTimeout(20).
            Exec()
```

- If the execution contains responses with special characters like `\n`, `EOF`, and others, utilize the `RemoveMatching()` method to eliminate them from `Stdout`.

```go
dockerCmd := dira.SetOpts(hostURLOrIP, containerName, cmd)

result, err := dockerCmd.
            RemoveMatching().
            Exec()
```

- If you have multiple commands to execute on a specific container, create an instance object and execute them sequentially.

```go
dockerCmd := dira.SetOpts(hostURLOrIP, containerName, "")

result1, err1 := dockerCmd.
                SetContainer("new-container-name"). // modify the container of the same network to run the command
                SetCommand("echo Command 1").
                Exec()

result2, err2 := dockerCmd.
                SetCommand("echo Command 2").
                SetTimeout(20).
                Exec()

result3, err3 := dockerCmd.
                SetCommand("echo Command 3").
                RemoveMatching().
                Exec()

```

### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please submit an issue or a pull
request on the GitHub repository.

### License

The DIRA package is open-source software licensed under the MIT license.

### Credits

The DIRA package is developed and maintained by Milad Roudgarian.
