package dira

type CommandAbstraction interface {
	SetRemoteAccess(rae ExecutionAbstraction) CommandAbstraction
	GetRemoteAccess() ExecutionAbstraction
	SetContainerId(containerId string) CommandAbstraction
	GetContainerId() string
	SetError(error error) CommandAbstraction
	GetError() error
	InquireContainer() CommandAbstraction
	ExecuteContainerCommand() (res string, err error)
}
