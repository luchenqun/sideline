package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EmployerList:  []Employer{},
		DeveloperList: []Developer{},
		TaskList:      []Task{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in employer
	employerIndexMap := make(map[string]struct{})

	for _, elem := range gs.EmployerList {
		index := string(EmployerKey(elem.Index))
		if _, ok := employerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for employer")
		}
		employerIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in developer
	developerIndexMap := make(map[string]struct{})

	for _, elem := range gs.DeveloperList {
		index := string(DeveloperKey(elem.Index))
		if _, ok := developerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for developer")
		}
		developerIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in task
	taskIdMap := make(map[uint64]bool)
	taskCount := gs.GetTaskCount()
	for _, elem := range gs.TaskList {
		if _, ok := taskIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for task")
		}
		if elem.Id >= taskCount {
			return fmt.Errorf("task id should be lower or equal than the last id")
		}
		taskIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
