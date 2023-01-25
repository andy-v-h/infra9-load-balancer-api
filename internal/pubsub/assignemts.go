package pubsub

import (
	"go.infratographer.com/x/pubsubx"
)

func NewAssignmentURN(assignmentID string) string {
	return "urn:infratographer:infratographer.com:assignment:" + assignmentID + "/"
}

func NewAssignmentMessage(actorURN string, tenantURN string, assignmentURN string, additionalSubjectURNs ...string) (*pubsubx.Message, error) {
	if !validURN([]byte(actorURN)) {
		return nil, ErrInvalidActorURN
	}
	if !validURN([]byte(tenantURN)) {
		return nil, ErrInvalidTenantURN
	}
	if !validURN([]byte(assignmentURN)) {
		return nil, ErrInvalidAssignmentURN
	}

	for _, u := range additionalSubjectURNs {
		if !validURN([]byte(u)) {
			return nil, ErrInvalidURN
		}
	}
	return newMessage(actorURN, assignmentURN, additionalSubjectURNs),nil
}
