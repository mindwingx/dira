package dira

type RemoteAccessEntity struct {
	url            string
	container      string
	command        string
	timeout        int
	removeMatching bool
}

// SetOpts insert params to execute the command
func SetOpts(hostIpOrUrl string, containerIdOrName string, cmd string) ExecutionAbstraction {
	return &RemoteAccessEntity{
		url:       hostIpOrUrl,
		container: containerIdOrName,
		command:   cmd,
	}
}

// SetUrl set the base url/ip & port of docker container host
func (e *RemoteAccessEntity) SetUrl(hostIpOrUrl string) ExecutionAbstraction {
	e.url = hostIpOrUrl
	return e
}
func (e *RemoteAccessEntity) GetUrl() string {
	return e.url
}

// SetContainer set the container id or name
func (e *RemoteAccessEntity) SetContainer(containerName string) ExecutionAbstraction {
	e.container = containerName
	return e
}

func (e *RemoteAccessEntity) GetContainer() string {
	return e.container
}

// SetCommand set desired command to be run in the container
func (e *RemoteAccessEntity) SetCommand(cmd string) ExecutionAbstraction {
	e.command = cmd
	return e
}

func (e *RemoteAccessEntity) GetCommand() string {
	return e.command
}

// SetTimeout set timeout in seconds, if it will be asserted too much time for execution
func (e *RemoteAccessEntity) SetTimeout(timeOutInSec int) ExecutionAbstraction {
	e.timeout = timeOutInSec
	return e
}

func (e *RemoteAccessEntity) GetTimeout() int {
	return e.timeout
}

// RemoveMatching remove unwanted characters of "\n, SOH, NUL, EOT, ACK" of STDOUT
func (e *RemoteAccessEntity) RemoveMatching() ExecutionAbstraction {
	e.removeMatching = true
	return e
}

func (e *RemoteAccessEntity) GetRemoveMatching() bool {
	return e.removeMatching
}

// Exec inquire the container name and execute the command
func (e *RemoteAccessEntity) Exec() (res string, err error) {
	cmd := Command()
	cmd.SetRemoteAccess(e)
	cmd.InquireContainer()

	if err = cmd.GetError(); err != nil {
		return "", err
	}

	return cmd.ExecuteContainerCommand()
}
