package dira

type CommandEntity struct {
	rae         ExecutionAbstraction
	containerId string
	err         error
}

func Command() CommandAbstraction {
	return &CommandEntity{}
}

func (ce *CommandEntity) SetRemoteAccess(rae ExecutionAbstraction) CommandAbstraction {
	ce.rae = rae
	return ce
}

func (ce *CommandEntity) GetRemoteAccess() ExecutionAbstraction {
	return ce.rae
}

func (ce *CommandEntity) SetContainerId(containerId string) CommandAbstraction {
	ce.containerId = containerId
	return ce
}

func (ce *CommandEntity) GetContainerId() string {
	return ce.containerId
}

func (ce *CommandEntity) SetError(error error) CommandAbstraction {
	ce.err = error
	return ce
}

func (ce *CommandEntity) GetError() error {
	return ce.err
}
