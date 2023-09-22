package dira

type ExecutionAbstraction interface {
	SetUrl(hostIpOrUrl string) ExecutionAbstraction
	GetUrl() string
	SetContainer(containerName string) ExecutionAbstraction
	GetContainer() string
	SetCommand(cmd string) ExecutionAbstraction
	GetCommand() string
	SetTimeout(timeOutInSec int) ExecutionAbstraction
	GetTimeout() int
	RemoveMatching() ExecutionAbstraction
	GetRemoveMatching() bool
	Exec() (res string, err error)
}
