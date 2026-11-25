// Package domain defines core business types and constants shared across Xinghe.
package domain

// BusinessState is the authoritative formal task stage stored in Git.
type BusinessState string

const (
	BusinessPublished          BusinessState = "published"
	BusinessSubmitted          BusinessState = "submitted"
	BusinessReviewing          BusinessState = "reviewing"
	BusinessNeedsRevision      BusinessState = "needs_revision"
	BusinessAwaitingAcceptance BusinessState = "awaiting_acceptance"
	BusinessAccepted           BusinessState = "accepted"
	BusinessCancelled          BusinessState = "cancelled"
)

// RunState is the authoritative runtime execution state stored in SQLite.
type RunState string

const (
	RunQueued   RunState = "queued"
	RunStarting RunState = "starting"
	RunRunning  RunState = "running"
	RunBlocked  RunState = "blocked"
	RunStopping RunState = "stopping"
	RunExited   RunState = "exited"
	RunFailed   RunState = "failed"
	RunUnknown  RunState = "unknown"
)

// AssistantLevel controls how much autonomy the swappable Admin Assistant has.
type AssistantLevel string

const (
	AssistantL1DraftOnly     AssistantLevel = "L1" // Default in v1: human confirms all formal actions
	AssistantL2PlannedFlow   AssistantLevel = "L2" // Auto-route pre-approved handoffs & routine wakes
	AssistantL3HighAutonomy  AssistantLevel = "L3" // Future: rule-based loops, human handles final accept
)

// DirectiveKind enumerates the 5 Admin -> Member message directives (with static vs wake inquire).
type DirectiveKind string

const (
	DirectiveAssign        DirectiveKind = "assign"
	DirectiveInquireStatic DirectiveKind = "inquire_static"
	DirectiveInquireWake   DirectiveKind = "inquire_wake"
	DirectiveSupplement    DirectiveKind = "supplement"
	DirectiveReturn        DirectiveKind = "return"
	DirectiveCancel        DirectiveKind = "cancel"
)

// ReportKind enumerates the 3 Member -> Admin report types.
type ReportKind string

const (
	ReportDeliver   ReportKind = "deliver"
	ReportProgress  ReportKind = "report"
	ReportHelpBlock ReportKind = "help_block"
)

// ErrorCode defines standard protocol errors returned by core services.
type ErrorCode string

const (
	ErrUnauthorized                  ErrorCode = "unauthorized"
	ErrForbiddenCrossMember          ErrorCode = "forbidden_cross_member"
	ErrAssistantConfirmationRequired ErrorCode = "assistant_confirmation_required"
	ErrVersionConflict               ErrorCode = "version_conflict"
	ErrLeaseLost                     ErrorCode = "lease_lost"
	ErrInvalidSubmission             ErrorCode = "invalid_submission"
	ErrFileScopeWarning              ErrorCode = "file_scope_warning"
	ErrInputChanged                  ErrorCode = "input_changed"
	ErrRunUnknown                    ErrorCode = "run_unknown"
	ErrPublicationPending            ErrorCode = "publication_pending"
)
