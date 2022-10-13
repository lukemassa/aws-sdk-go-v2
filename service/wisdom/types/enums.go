// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AssistantStatus string

// Enum values for AssistantStatus
const (
	AssistantStatusCreateInProgress AssistantStatus = "CREATE_IN_PROGRESS"
	AssistantStatusCreateFailed     AssistantStatus = "CREATE_FAILED"
	AssistantStatusActive           AssistantStatus = "ACTIVE"
	AssistantStatusDeleteInProgress AssistantStatus = "DELETE_IN_PROGRESS"
	AssistantStatusDeleteFailed     AssistantStatus = "DELETE_FAILED"
	AssistantStatusDeleted          AssistantStatus = "DELETED"
)

// Values returns all known values for AssistantStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssistantStatus) Values() []AssistantStatus {
	return []AssistantStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
	}
}

type AssistantType string

// Enum values for AssistantType
const (
	AssistantTypeAgent AssistantType = "AGENT"
)

// Values returns all known values for AssistantType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssistantType) Values() []AssistantType {
	return []AssistantType{
		"AGENT",
	}
}

type AssociationType string

// Enum values for AssociationType
const (
	AssociationTypeKnowledgeBase AssociationType = "KNOWLEDGE_BASE"
)

// Values returns all known values for AssociationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssociationType) Values() []AssociationType {
	return []AssociationType{
		"KNOWLEDGE_BASE",
	}
}

type ContentStatus string

// Enum values for ContentStatus
const (
	ContentStatusCreateInProgress ContentStatus = "CREATE_IN_PROGRESS"
	ContentStatusCreateFailed     ContentStatus = "CREATE_FAILED"
	ContentStatusActive           ContentStatus = "ACTIVE"
	ContentStatusDeleteInProgress ContentStatus = "DELETE_IN_PROGRESS"
	ContentStatusDeleteFailed     ContentStatus = "DELETE_FAILED"
	ContentStatusDeleted          ContentStatus = "DELETED"
	ContentStatusUpdateFailed     ContentStatus = "UPDATE_FAILED"
)

// Values returns all known values for ContentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContentStatus) Values() []ContentStatus {
	return []ContentStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
		"UPDATE_FAILED",
	}
}

type FilterField string

// Enum values for FilterField
const (
	FilterFieldName FilterField = "NAME"
)

// Values returns all known values for FilterField. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FilterField) Values() []FilterField {
	return []FilterField{
		"NAME",
	}
}

type FilterOperator string

// Enum values for FilterOperator
const (
	FilterOperatorEquals FilterOperator = "EQUALS"
)

// Values returns all known values for FilterOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FilterOperator) Values() []FilterOperator {
	return []FilterOperator{
		"EQUALS",
	}
}

type KnowledgeBaseStatus string

// Enum values for KnowledgeBaseStatus
const (
	KnowledgeBaseStatusCreateInProgress KnowledgeBaseStatus = "CREATE_IN_PROGRESS"
	KnowledgeBaseStatusCreateFailed     KnowledgeBaseStatus = "CREATE_FAILED"
	KnowledgeBaseStatusActive           KnowledgeBaseStatus = "ACTIVE"
	KnowledgeBaseStatusDeleteInProgress KnowledgeBaseStatus = "DELETE_IN_PROGRESS"
	KnowledgeBaseStatusDeleteFailed     KnowledgeBaseStatus = "DELETE_FAILED"
	KnowledgeBaseStatusDeleted          KnowledgeBaseStatus = "DELETED"
)

// Values returns all known values for KnowledgeBaseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseStatus) Values() []KnowledgeBaseStatus {
	return []KnowledgeBaseStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETED",
	}
}

type KnowledgeBaseType string

// Enum values for KnowledgeBaseType
const (
	KnowledgeBaseTypeExternal KnowledgeBaseType = "EXTERNAL"
	KnowledgeBaseTypeCustom   KnowledgeBaseType = "CUSTOM"
)

// Values returns all known values for KnowledgeBaseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseType) Values() []KnowledgeBaseType {
	return []KnowledgeBaseType{
		"EXTERNAL",
		"CUSTOM",
	}
}

type RecommendationSourceType string

// Enum values for RecommendationSourceType
const (
	RecommendationSourceTypeIssueDetection RecommendationSourceType = "ISSUE_DETECTION"
	RecommendationSourceTypeRuleEvaluation RecommendationSourceType = "RULE_EVALUATION"
	RecommendationSourceTypeOther          RecommendationSourceType = "OTHER"
)

// Values returns all known values for RecommendationSourceType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationSourceType) Values() []RecommendationSourceType {
	return []RecommendationSourceType{
		"ISSUE_DETECTION",
		"RULE_EVALUATION",
		"OTHER",
	}
}

type RecommendationTriggerType string

// Enum values for RecommendationTriggerType
const (
	RecommendationTriggerTypeQuery RecommendationTriggerType = "QUERY"
)

// Values returns all known values for RecommendationTriggerType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationTriggerType) Values() []RecommendationTriggerType {
	return []RecommendationTriggerType{
		"QUERY",
	}
}

type RecommendationType string

// Enum values for RecommendationType
const (
	RecommendationTypeKnowledgeContent RecommendationType = "KNOWLEDGE_CONTENT"
)

// Values returns all known values for RecommendationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationType) Values() []RecommendationType {
	return []RecommendationType{
		"KNOWLEDGE_CONTENT",
	}
}

type RelevanceLevel string

// Enum values for RelevanceLevel
const (
	RelevanceLevelHigh   RelevanceLevel = "HIGH"
	RelevanceLevelMedium RelevanceLevel = "MEDIUM"
	RelevanceLevelLow    RelevanceLevel = "LOW"
)

// Values returns all known values for RelevanceLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelevanceLevel) Values() []RelevanceLevel {
	return []RelevanceLevel{
		"HIGH",
		"MEDIUM",
		"LOW",
	}
}
