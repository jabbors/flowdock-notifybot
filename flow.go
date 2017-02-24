package main

import (
	"github.com/gnyman/flowdock"
)

// Flows is a map of know flows
type Flows map[string]flowdock.Flow

// NewFlows returns a empty flow map
func NewFlows() Flows {
	return make(map[string]flowdock.Flow)
}

// Exists returns true the flow exits
func (f Flows) Exists(ID string) bool {
	if _, ok := f[ID]; ok {
		return true
	}
	return false
}

// Add adds a flow to the map
func (f Flows) Add(ID string, flow flowdock.Flow) {
	f[ID] = flow
}

// Clear clears all flows in the map
func (f Flows) Clear() {
	f = make(map[string]flowdock.Flow)
}
