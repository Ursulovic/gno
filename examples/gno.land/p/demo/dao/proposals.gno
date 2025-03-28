package dao

import "std"

// ProposalStatus is the currently active proposal status,
// changed based on DAO functionality.
// Status transitions:
//
// ACTIVE -> ACCEPTED -> EXECUTION(SUCCEEDED/FAILED)
//
// ACTIVE -> NOT ACCEPTED
type ProposalStatus string

var (
	Active              ProposalStatus = "active"               // proposal is still active
	Accepted            ProposalStatus = "accepted"             // proposal gathered quorum
	NotAccepted         ProposalStatus = "not accepted"         // proposal failed to gather quorum
	ExecutionSuccessful ProposalStatus = "execution successful" // proposal is executed successfully
	ExecutionFailed     ProposalStatus = "execution failed"     // proposal has failed during execution
)

func (s ProposalStatus) String() string {
	return string(s)
}

// PropStore defines the proposal storage abstraction
type PropStore interface {
	// Proposals returns the given paginated proposals
	Proposals(offset, count uint64) []Proposal

	// ProposalByID returns the proposal associated with
	// the given ID, if any
	ProposalByID(id uint64) (Proposal, error)

	// Size returns the number of proposals in
	// the proposal store
	Size() int
}

// Proposal is the single proposal abstraction
type Proposal interface {
	// Author returns the author of the proposal
	Author() std.Address

	// Title returns the title of the proposal
	Title() string

	// Description returns the description of the proposal
	Description() string

	// Status returns the status of the proposal
	Status() ProposalStatus

	// Executor returns the proposal executor
	Executor() Executor

	// Stats returns the voting stats of the proposal
	Stats() Stats

	// IsExpired returns a flag indicating if the proposal expired
	IsExpired() bool

	// Render renders the proposal in a readable format
	Render() string
}
