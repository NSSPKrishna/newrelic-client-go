// Code generated by tutone: DO NOT EDIT
package logconfigurations

import (
	"github.com/newrelic/newrelic-client-go/v2/pkg/nrtime"
)

// LogConfigurationsCreateDataPartitionRuleErrorType - Expected error types as result of creating a new data partition rule.
type LogConfigurationsCreateDataPartitionRuleErrorType string

var LogConfigurationsCreateDataPartitionRuleErrorTypeTypes = struct {
	// A data partition rule with the provided name already exists.
	DUPLICATE_DATA_PARTITION_RULE_NAME LogConfigurationsCreateDataPartitionRuleErrorType
	// The provided data partition does not match the validation requirements
	INVALID_DATA_PARTITION_INPUT LogConfigurationsCreateDataPartitionRuleErrorType
	// Customer has reached the maximum number of allowed data partition rules.
	MAX_DATA_PARTITION_RULES LogConfigurationsCreateDataPartitionRuleErrorType
}{
	// A data partition rule with the provided name already exists.
	DUPLICATE_DATA_PARTITION_RULE_NAME: "DUPLICATE_DATA_PARTITION_RULE_NAME",
	// The provided data partition does not match the validation requirements
	INVALID_DATA_PARTITION_INPUT: "INVALID_DATA_PARTITION_INPUT",
	// Customer has reached the maximum number of allowed data partition rules.
	MAX_DATA_PARTITION_RULES: "MAX_DATA_PARTITION_RULES",
}

// LogConfigurationsDataPartitionRuleMatchingOperator - The matching method for the rule to allocate the data partition data.
// Select EQUALS to target logs that match your criteria exactly, or select LIKE to apply a fuzzy match.
type LogConfigurationsDataPartitionRuleMatchingOperator string

var LogConfigurationsDataPartitionRuleMatchingOperatorTypes = struct {
	// When applying the rule will allocate data for those attributes that are an exact match with the provided value.
	EQUALS LogConfigurationsDataPartitionRuleMatchingOperator
	// When applying the rule will allocate data for those attributes that contain the provided value.
	LIKE LogConfigurationsDataPartitionRuleMatchingOperator
}{
	// When applying the rule will allocate data for those attributes that are an exact match with the provided value.
	EQUALS: "EQUALS",
	// When applying the rule will allocate data for those attributes that contain the provided value.
	LIKE: "LIKE",
}

// LogConfigurationsDataPartitionRuleMutationErrorType - Expected default error types as result of mutating an existing data partition rule.
type LogConfigurationsDataPartitionRuleMutationErrorType string

var LogConfigurationsDataPartitionRuleMutationErrorTypeTypes = struct {
	// Number format error. ID should be convertible to int.
	INVALID_ID LogConfigurationsDataPartitionRuleMutationErrorType
	// Couldn't find the provided data partition rule.
	NOT_FOUND LogConfigurationsDataPartitionRuleMutationErrorType
}{
	// Number format error. ID should be convertible to int.
	INVALID_ID: "INVALID_ID",
	// Couldn't find the provided data partition rule.
	NOT_FOUND: "NOT_FOUND",
}

// LogConfigurationsDataPartitionRuleRetentionPolicyType - The retention policy for the data partition data.
type LogConfigurationsDataPartitionRuleRetentionPolicyType string

var LogConfigurationsDataPartitionRuleRetentionPolicyTypeTypes = struct {
	// The alternative data retention policy, 30 days of data retention since the log data is ingested.
	SECONDARY LogConfigurationsDataPartitionRuleRetentionPolicyType
	// The maximum retention period associated with the account. This is determined by the customer’s subscription/contract with New Relic.
	STANDARD LogConfigurationsDataPartitionRuleRetentionPolicyType
}{
	// The alternative data retention policy, 30 days of data retention since the log data is ingested.
	SECONDARY: "SECONDARY",
	// The maximum retention period associated with the account. This is determined by the customer’s subscription/contract with New Relic.
	STANDARD: "STANDARD",
}

// LogConfigurationsObfuscationMethod - Methods for replacing obfuscated values.
type LogConfigurationsObfuscationMethod string

var LogConfigurationsObfuscationMethodTypes = struct {
	// Replace the matched data with a SHA256 hash.
	HASH_SHA256 LogConfigurationsObfuscationMethod
	// Replace the matched data with a static value.
	MASK LogConfigurationsObfuscationMethod
}{
	// Replace the matched data with a SHA256 hash.
	HASH_SHA256: "HASH_SHA256",
	// Replace the matched data with a static value.
	MASK: "MASK",
}

// LogConfigurationsParsingRuleMutationErrorType - Expected default error types as result of mutating an existing parsing rule.
type LogConfigurationsParsingRuleMutationErrorType string

var LogConfigurationsParsingRuleMutationErrorTypeTypes = struct {
	// Invalid Grok
	INVALID_GROK LogConfigurationsParsingRuleMutationErrorType
	// Number format error. ID should be convertible to int.
	INVALID_ID LogConfigurationsParsingRuleMutationErrorType
	// Invalid NRQL
	INVALID_NRQL LogConfigurationsParsingRuleMutationErrorType
	// Couldn't find the specified parsing rule.
	NOT_FOUND LogConfigurationsParsingRuleMutationErrorType
}{
	// Invalid Grok
	INVALID_GROK: "INVALID_GROK",
	// Number format error. ID should be convertible to int.
	INVALID_ID: "INVALID_ID",
	// Invalid NRQL
	INVALID_NRQL: "INVALID_NRQL",
	// Couldn't find the specified parsing rule.
	NOT_FOUND: "NOT_FOUND",
}

// Account - The `Account` object provides general data about the account, as well as
// being the entry point into more detailed data about a single account.
//
// Account configuration data is queried through this object, as well as
// telemetry data that is specific to a single account.
type Account struct {
	//
	ID int `json:"id,omitempty"`
	//
	LicenseKey string `json:"licenseKey,omitempty"`
	// This field provides access to LogConfigurations data.
	LogConfigurations LogConfigurationsAccountStitchedFields `json:"logConfigurations,omitempty"`
	//
	Name string `json:"name,omitempty"`
}

// Actor - The `Actor` object contains fields that are scoped to the API user's access level.
type Actor struct {
	// The `account` field is the entry point into data that is scoped to a single account.
	Account Account `json:"account,omitempty"`
}

// LogConfigurationsAccountStitchedFields -
type LogConfigurationsAccountStitchedFields struct {
	// Look up for all data partition rules for a given account.
	DataPartitionRules []LogConfigurationsDataPartitionRule `json:"dataPartitionRules"`
	// Look up for all obfuscation expressions for a given account
	ObfuscationExpressions []LogConfigurationsObfuscationExpression `json:"obfuscationExpressions"`
	// Look up for all obfuscation rules for a given account.
	ObfuscationRules []LogConfigurationsObfuscationRule `json:"obfuscationRules"`
	// Look up for all parsing rules for a given account.
	ParsingRules []*LogConfigurationsParsingRule `json:"parsingRules"`
	// Test a Grok pattern against a list of log lines.
	TestGrok []LogConfigurationsGrokTestResult `json:"testGrok"`
}

// LogConfigurationsCreateDataPartitionRuleError - Expected errors as a result of creating a new data partition rule.
type LogConfigurationsCreateDataPartitionRuleError struct {
	// The message with the error cause.
	Message string `json:"message,omitempty"`
	// Type of error.
	Type LogConfigurationsCreateDataPartitionRuleErrorType `json:"type,omitempty"`
}

// LogConfigurationsCreateDataPartitionRuleInput - A new data partition rule.
type LogConfigurationsCreateDataPartitionRuleInput struct {
	// The description of the data partition rule.
	Description string `json:"description,omitempty"`
	// Whether or not this data partition rule is enabled.
	Enabled bool `json:"enabled"`
	// The matching criteria of the data partition rule.
	MatchingCriteria *LogConfigurationsDataPartitionRuleMatchingCriteriaInput `json:"matchingCriteria,omitempty"`
	// The retention policy of the data partition data.
	RetentionPolicy LogConfigurationsDataPartitionRuleRetentionPolicyType `json:"retentionPolicy"`
	// The name of the data partition where logs will be allocated once the rule is enabled.
	TargetDataPartition LogConfigurationsLogDataPartitionName `json:"targetDataPartition"`
}

// LogConfigurationsCreateDataPartitionRuleResponse - The result after creating a new data partition rule.
type LogConfigurationsCreateDataPartitionRuleResponse struct {
	// List of errors, if any.
	Errors []LogConfigurationsCreateDataPartitionRuleError `json:"errors,omitempty"`
	// The created data partition rule.
	Rule LogConfigurationsDataPartitionRule `json:"rule,omitempty"`
}

// LogConfigurationsCreateObfuscationActionInput - Input for creating an obfuscation action on a rule being created.
type LogConfigurationsCreateObfuscationActionInput struct {
	// Attribute names for action. An empty list applies the action to all the attributes.
	Attributes []string `json:"attributes"`
	// Expression Id for action.
	ExpressionId string `json:"expressionId"`
	// Obfuscation method to use.
	Method LogConfigurationsObfuscationMethod `json:"method"`
}

// LogConfigurationsCreateObfuscationExpressionInput - Input for creating an obfuscation expression.
type LogConfigurationsCreateObfuscationExpressionInput struct {
	// Description of expression.
	Description string `json:"description,omitempty"`
	// Name of expression.
	Name string `json:"name"`
	// Regex of expression.
	Regex string `json:"regex"`
}

// LogConfigurationsCreateObfuscationRuleInput - Input for creating an obfuscation rule.
type LogConfigurationsCreateObfuscationRuleInput struct {
	// Actions for the rule. The actions will be applied in the order specified by this list.
	Actions []LogConfigurationsCreateObfuscationActionInput `json:"actions,omitempty"`
	// Description of rule.
	Description string `json:"description,omitempty"`
	// Whether the rule should be applied or not to incoming data.
	Enabled bool `json:"enabled"`
	// NRQL for determining whether a given log record should have obfuscation actions applied.
	Filter NRQL `json:"filter"`
	// Name of rule.
	Name string `json:"name"`
}

// LogConfigurationsCreateParsingRuleResponse - The result after creating a new parsing rule.
type LogConfigurationsCreateParsingRuleResponse struct {
	// List of errors, if any.
	Errors []LogConfigurationsParsingRuleMutationError `json:"errors,omitempty"`
	// The created parsing rule.
	Rule *LogConfigurationsParsingRule `json:"rule,omitempty"`
}

// LogConfigurationsDataPartitionRule - The data partition rule for an account.
type LogConfigurationsDataPartitionRule struct {
	// Identifies the date and time when the rule was created.
	CreatedAt nrtime.DateTime `json:"createdAt"`
	// Identifies the user who has created the rule.
	CreatedBy UserReference `json:"createdBy,omitempty"`
	// Whether or not this data partition rule is deleted. Deleting a data partition rule does not delete the already persisted data. This data will be retained for a given period of time specified in the retention policy field.
	Deleted bool `json:"deleted"`
	// A description of what this data partition rule represents.
	Description string `json:"description,omitempty"`
	// Whether or not this data partition rule is enabled.
	Enabled bool `json:"enabled"`
	// Unique data partition rule identifier.
	ID string `json:"id"`
	// The matching criteria for this data partition rule. Logs matching this criteria will be routed to the specified data partition once the rule is enabled.
	MatchingCriteria LogConfigurationsDataPartitionRuleMatchingCriteria `json:"matchingCriteria"`
	// The retention policy of the data partition data.
	RetentionPolicy LogConfigurationsDataPartitionRuleRetentionPolicyType `json:"retentionPolicy"`
	// The name of the data partition.
	TargetDataPartition LogConfigurationsLogDataPartitionName `json:"targetDataPartition"`
	// Identifies the date and time when the rule was last updated.
	UpdatedAt nrtime.DateTime `json:"updatedAt,omitempty"`
	// Identifies the user who has last updated the rule.
	UpdatedBy UserReference `json:"updatedBy,omitempty"`
}

// LogConfigurationsDataPartitionRuleMatchingCriteria - The data partition rule matching criteria.
type LogConfigurationsDataPartitionRuleMatchingCriteria struct {
	// The attribute name against which this matching condition will be evaluated.
	AttributeName string `json:"attributeName"`
	// The matching expression of the data partition rule definition.
	MatchingExpression string `json:"matchingExpression"`
	// The matching method of the data partition rule definition.
	MatchingOperator LogConfigurationsDataPartitionRuleMatchingOperator `json:"matchingOperator"`
}

// LogConfigurationsDataPartitionRuleMatchingCriteriaInput - The data partition rule matching criteria.
type LogConfigurationsDataPartitionRuleMatchingCriteriaInput struct {
	// The attribute name against which this matching condition will be evaluated.
	AttributeName string `json:"attributeName"`
	// The matching expression of the data partition rule definition.
	MatchingExpression string `json:"matchingExpression"`
	// The matching method of the data partition rule definition.
	MatchingMethod LogConfigurationsDataPartitionRuleMatchingOperator `json:"matchingMethod"`
}

// LogConfigurationsDataPartitionRuleMutationError - An object that contains expected errors as a result of mutating an existing data partition rule.
type LogConfigurationsDataPartitionRuleMutationError struct {
	// The message with the error cause.
	Message string `json:"message,omitempty"`
	// Type of error.
	Type LogConfigurationsDataPartitionRuleMutationErrorType `json:"type,omitempty"`
}

// LogConfigurationsDeleteDataPartitionRuleResponse - The result after deleting a data partition rule.
type LogConfigurationsDeleteDataPartitionRuleResponse struct {
	// List of errors, if any.
	Errors []LogConfigurationsDataPartitionRuleMutationError `json:"errors,omitempty"`
}

// LogConfigurationsDeleteParsingRuleResponse - The result after deleting a parsing rule.
type LogConfigurationsDeleteParsingRuleResponse struct {
	// List of errors, if any.
	Errors []LogConfigurationsParsingRuleMutationError `json:"errors,omitempty"`
}

// LogConfigurationsGrokTestExtractedAttribute - An attribute that was extracted from a Grok test.
type LogConfigurationsGrokTestExtractedAttribute struct {
	// The attribute name.
	Name string `json:"name"`
	// A string representation of the extracted value (which might not be a String).
	Value string `json:"value"`
}

// LogConfigurationsGrokTestResult - The result of testing Grok on a log line.
type LogConfigurationsGrokTestResult struct {
	// Any attributes that were extracted.
	Attributes []LogConfigurationsGrokTestExtractedAttribute `json:"attributes"`
	// The log line that was tested against.
	LogLine string `json:"logLine"`
	// Whether the Grok pattern matched.
	Matched bool `json:"matched"`
}

// LogConfigurationsObfuscationAction - Application of an obfuscation expression with specific a replacement method.
type LogConfigurationsObfuscationAction struct {
	// Log record attributes to apply this expression to. An empty list applies the action to all the attributes.
	Attributes []string `json:"attributes"`
	// Obfuscation expression applied by this action.
	Expression LogConfigurationsObfuscationExpression `json:"expression"`
	// The id of the obfuscation action.
	ID string `json:"id"`
	// How to obfuscate matches for the applied expression.
	Method LogConfigurationsObfuscationMethod `json:"method"`
}

// LogConfigurationsObfuscationExpression - Reusable obfuscation expression.
type LogConfigurationsObfuscationExpression struct {
	// Identifies the date and time when the expression was created.
	CreatedAt nrtime.DateTime `json:"createdAt"`
	// Identifies the user who has created the expression.
	CreatedBy UserReference `json:"createdBy,omitempty"`
	// Description of the expression.
	Description string `json:"description,omitempty"`
	// The id of the obfuscation expression.
	ID string `json:"id"`
	// Name of the expression.
	Name string `json:"name"`
	// Regular expression for this obfuscation expression. Capture groups will be obscured on matching.
	Regex string `json:"regex"`
	// Identifies the date and time when the expression was last updated.
	UpdatedAt nrtime.DateTime `json:"updatedAt"`
	// Identifies the user who has last updated the expression.
	UpdatedBy UserReference `json:"updatedBy,omitempty"`
}

// LogConfigurationsObfuscationRule - Rule for identifying a set of log data to apply specific obfuscation actions to.
type LogConfigurationsObfuscationRule struct {
	// Obfuscation actions to take if a record passes the matching criteria.
	Actions []LogConfigurationsObfuscationAction `json:"actions"`
	// Identifies the date and time when the rule was created.
	CreatedAt nrtime.DateTime `json:"createdAt"`
	// Identifies the user who has created the rule.
	CreatedBy UserReference `json:"createdBy,omitempty"`
	// Description of the obfuscation rule.
	Description string `json:"description,omitempty"`
	// Whether the rule should be applied to incoming logs
	Enabled bool `json:"enabled"`
	// NRQL filter to determine if a log record should have obfuscation actions applied.
	Filter NRQL `json:"filter"`
	// The id of the obfuscation rule.
	ID string `json:"id"`
	// Name of the obfuscation rule.
	Name string `json:"name"`
	// Identifies the date and time when the rule was last updated.
	UpdatedAt nrtime.DateTime `json:"updatedAt"`
	// Identifies the user who has last updated the rule.
	UpdatedBy UserReference `json:"updatedBy,omitempty"`
}

// LogConfigurationsParsingRule - A parsing rule for an account.
type LogConfigurationsParsingRule struct {
	// The account id associated with the rule.
	AccountID int `json:"accountId"`
	// The parsing rule will apply to value of this attribute.
	Attribute string `json:"attribute"`
	// Identifies the user who has created the rule.
	CreatedBy UserReference `json:"createdBy,omitempty"`
	// Whether or not this rule is deleted.
	Deleted bool `json:"deleted"`
	// A description of what this parsing rule represents.
	Description string `json:"description"`
	// Whether or not this rule is enabled.
	Enabled bool `json:"enabled"`
	// The Grok of what to parse.
	Grok string `json:"grok"`
	// Unique parsing rule identifier.
	ID string `json:"id"`
	// The Lucene to match events to the parsing rule.
	Lucene string `json:"lucene"`
	// The NRQL to match events to the parsing rule.
	NRQL NRQL `json:"nrql"`
	// Identifies the date and time when the rule was last updated.
	UpdatedAt nrtime.DateTime `json:"updatedAt,omitempty"`
	// Identifies the user who has last updated the rule.
	UpdatedBy UserReference `json:"updatedBy,omitempty"`
}

// LogConfigurationsParsingRuleConfiguration - A new parsing rule.
type LogConfigurationsParsingRuleConfiguration struct {
	// The parsing rule will apply to value of this attribute. If field is not provided, value will default to message.
	Attribute string `json:"attribute,omitempty"`
	// A description of what this parsing rule represents.
	Description string `json:"description"`
	// Whether or not this rule is enabled.
	Enabled bool `json:"enabled"`
	// The Grok of what to parse.
	Grok string `json:"grok"`
	// The Lucene to match events to the parsing rule.
	Lucene string `json:"lucene"`
	// The NRQL to match events to the parsing rule.
	NRQL NRQL `json:"nrql"`
}

// LogConfigurationsParsingRuleMutationError - Expected errors as a result of mutating a parsing rule.
type LogConfigurationsParsingRuleMutationError struct {
	// The message with the error cause.
	Message string `json:"message,omitempty"`
	// Type of error.
	Type LogConfigurationsParsingRuleMutationErrorType `json:"type,omitempty"`
}

// LogConfigurationsUpdateDataPartitionRuleInput - An object for updating an existing data partition rule.
type LogConfigurationsUpdateDataPartitionRuleInput struct {
	// The description of the data partition rule.
	Description string `json:"description,omitempty"`
	// Whether or not this data partition rule is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Unique data partition rule identifier.
	ID string `json:"id"`
	// The criteria of the data partition rule.
	MatchingCriteria *LogConfigurationsDataPartitionRuleMatchingCriteriaInput `json:"matchingCriteria,omitempty"`
}

// LogConfigurationsUpdateDataPartitionRuleResponse - An object that represents the result after updating a data partition rule.
type LogConfigurationsUpdateDataPartitionRuleResponse struct {
	// List of errors, if any.
	Errors []LogConfigurationsDataPartitionRuleMutationError `json:"errors,omitempty"`
	// The updated data partition rule.
	Rule LogConfigurationsDataPartitionRule `json:"rule,omitempty"`
}

// LogConfigurationsUpdateObfuscationActionInput - Input for creating an obfuscation action on a rule being updated.
type LogConfigurationsUpdateObfuscationActionInput struct {
	// Attribute names for action. An empty list applies the action to all the attributes.
	Attributes []string `json:"attributes"`
	// Expression Id for action.
	ExpressionId string `json:"expressionId"`
	// Obfuscation method to use.
	Method LogConfigurationsObfuscationMethod `json:"method"`
}

// LogConfigurationsUpdateObfuscationExpressionInput - Input for updating an obfuscation expression.
// Null fields are left untouched by mutation.
type LogConfigurationsUpdateObfuscationExpressionInput struct {
	// Description of expression.
	Description string `json:"description,omitempty"`
	// Expression Id.
	ID string `json:"id"`
	// Name of expression.
	Name string `json:"name,omitempty"`
	// Regex of expression.
	Regex string `json:"regex,omitempty"`
}

// LogConfigurationsUpdateObfuscationRuleInput - Input for updating an obfuscation rule.
// Null fields are left untouched by mutation.
type LogConfigurationsUpdateObfuscationRuleInput struct {
	// Actions for the rule. When non-null, this list of actions is used to replace
	// the existing list of actions of the rule. The actions will be applied in the
	// order specified by this list.
	Actions []LogConfigurationsUpdateObfuscationActionInput `json:"actions,omitempty"`
	// Description of rule.
	Description string `json:"description,omitempty"`
	// Whether the rule should be applied or not to incoming data.
	Enabled bool `json:"enabled,omitempty"`
	// NRQL for determining whether a given log record should have obfuscation actions applied.
	Filter NRQL `json:"filter,omitempty"`
	// Rule Id.
	ID string `json:"id"`
	// Name of rule.
	Name string `json:"name,omitempty"`
}

// LogConfigurationsUpdateParsingRuleResponse - The result after updating a parsing rule.
type LogConfigurationsUpdateParsingRuleResponse struct {
	// List of errors, if any.
	Errors []LogConfigurationsParsingRuleMutationError `json:"errors,omitempty"`
	// The updated parsing rule.
	Rule *LogConfigurationsParsingRule `json:"rule,omitempty"`
}

// UserReference - The `UserReference` object provides basic identifying information about the user.
type UserReference struct {
	//
	Email string `json:"email,omitempty"`
	//
	Gravatar string `json:"gravatar,omitempty"`
	//
	ID int `json:"id,omitempty"`
	//
	Name string `json:"name,omitempty"`
}

type dataPartitionRulesResponse struct {
	Actor Actor `json:"actor"`
}

type obfuscationExpressionsResponse struct {
	Actor Actor `json:"actor"`
}

type obfuscationRulesResponse struct {
	Actor Actor `json:"actor"`
}

type parsingRulesResponse struct {
	Actor Actor `json:"actor"`
}

type testGrokResponse struct {
	Actor Actor `json:"actor"`
}

// LogConfigurationsLogDataPartitionName - The name of a log data partition. Has to start with 'Log_' prefix and can only contain alphanumeric characters and underscores.
type LogConfigurationsLogDataPartitionName string

// NRQL - This scalar represents a NRQL query string.
//
// See the [NRQL Docs](https://docs.newrelic.com/docs/insights/nrql-new-relic-query-language/nrql-resources/nrql-syntax-components-functions) for more information about NRQL syntax.
type NRQL string
