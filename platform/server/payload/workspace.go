package payload

import (
	"mime/multipart"
	"time"
)

type WorkspacePath struct {
	WorkspaceId int `params:"workspaceId" validate:"required" json:"-"`
}

type WorkspaceParticipantPath struct {
	WorkspacePath
	UserId string `params:"userId" validate:"required" json:"-"`
}

type AssignmentPath struct {
	WorkspacePath
	AssignmentId int `params:"assignmentId" validate:"required" json:"-"`
}

type CreateWorkspacePayload struct {
	Name    string         `json:"name" validate:"required"`
	Profile multipart.File `file:"profile"`
}

type UpdateWorkspacePayload struct {
	WorkspacePath
	Name     *string        `json:"name"`
	Favorite *bool          `json:"favorite"`
	Archive  *bool          `json:"archive"`
	Profile  multipart.File `file:"profile"`
}

type CreateInvitationPayload struct {
	WorkspacePath
	ValidAt    time.Time `json:"validAt" validate:"required"`
	ValidUntil time.Time `json:"validUntil" validate:"required"`
}

type ListSubmissionPayload struct {
	AssignmentPath
	All bool `query:"all"`
}

// TODO: Fix panic caused by sending empty body payload to be validate by struct with only pointer
type UpdateParticipantPayload struct {
	WorkspaceParticipantPath
	Role string `json:"role" validate:"required"`
}
