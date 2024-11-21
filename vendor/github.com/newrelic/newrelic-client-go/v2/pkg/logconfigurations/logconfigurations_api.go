// Code generated by tutone: DO NOT EDIT
package logconfigurations

import (
	"context"

	"github.com/newrelic/newrelic-client-go/v2/pkg/errors"
)

// Create a new data partition rule.
func (a *Logconfigurations) LogConfigurationsCreateDataPartitionRule(
	accountID int,
	rule LogConfigurationsCreateDataPartitionRuleInput,
) (*LogConfigurationsCreateDataPartitionRuleResponse, error) {
	return a.LogConfigurationsCreateDataPartitionRuleWithContext(context.Background(),
		accountID,
		rule,
	)
}

// Create a new data partition rule.
func (a *Logconfigurations) LogConfigurationsCreateDataPartitionRuleWithContext(
	ctx context.Context,
	accountID int,
	rule LogConfigurationsCreateDataPartitionRuleInput,
) (*LogConfigurationsCreateDataPartitionRuleResponse, error) {

	resp := LogConfigurationsCreateDataPartitionRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"rule":      rule,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsCreateDataPartitionRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsCreateDataPartitionRuleResponse, nil
}

type LogConfigurationsCreateDataPartitionRuleQueryResponse struct {
	LogConfigurationsCreateDataPartitionRuleResponse LogConfigurationsCreateDataPartitionRuleResponse `json:"LogConfigurationsCreateDataPartitionRule"`
}

const LogConfigurationsCreateDataPartitionRuleMutation = `mutation(
	$accountId: Int!,
	$rule: LogConfigurationsCreateDataPartitionRuleInput!,
) { logConfigurationsCreateDataPartitionRule(
	accountId: $accountId,
	rule: $rule,
) {
	errors {
		message
		type
	}
	rule {
		createdAt
		createdBy {
			email
			gravatar
			id
			name
		}
		deleted
		description
		enabled
		id
		matchingCriteria {
			attributeName
			matchingExpression
			matchingOperator
		}
		nrql
		retentionPolicy
		targetDataPartition
		updatedAt
		updatedBy {
			email
			gravatar
			id
			name
		}
	}
} }`

// Create an obfuscation expression.
func (a *Logconfigurations) LogConfigurationsCreateObfuscationExpression(
	accountID int,
	expression LogConfigurationsCreateObfuscationExpressionInput,
) (*LogConfigurationsObfuscationExpression, error) {
	return a.LogConfigurationsCreateObfuscationExpressionWithContext(context.Background(),
		accountID,
		expression,
	)
}

// Create an obfuscation expression.
func (a *Logconfigurations) LogConfigurationsCreateObfuscationExpressionWithContext(
	ctx context.Context,
	accountID int,
	expression LogConfigurationsCreateObfuscationExpressionInput,
) (*LogConfigurationsObfuscationExpression, error) {

	resp := LogConfigurationsCreateObfuscationExpressionQueryResponse{}
	vars := map[string]interface{}{
		"accountId":  accountID,
		"expression": expression,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsCreateObfuscationExpressionMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsObfuscationExpression, nil
}

type LogConfigurationsCreateObfuscationExpressionQueryResponse struct {
	LogConfigurationsObfuscationExpression LogConfigurationsObfuscationExpression `json:"LogConfigurationsCreateObfuscationExpression"`
}

const LogConfigurationsCreateObfuscationExpressionMutation = `mutation(
	$accountId: Int!,
	$expression: LogConfigurationsCreateObfuscationExpressionInput!,
) { logConfigurationsCreateObfuscationExpression(
	accountId: $accountId,
	expression: $expression,
) {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	id
	name
	regex
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Create an obfuscation rule.
func (a *Logconfigurations) LogConfigurationsCreateObfuscationRule(
	accountID int,
	rule LogConfigurationsCreateObfuscationRuleInput,
) (*LogConfigurationsObfuscationRule, error) {
	return a.LogConfigurationsCreateObfuscationRuleWithContext(context.Background(),
		accountID,
		rule,
	)
}

// Create an obfuscation rule.
func (a *Logconfigurations) LogConfigurationsCreateObfuscationRuleWithContext(
	ctx context.Context,
	accountID int,
	rule LogConfigurationsCreateObfuscationRuleInput,
) (*LogConfigurationsObfuscationRule, error) {

	resp := LogConfigurationsCreateObfuscationRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"rule":      rule,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsCreateObfuscationRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsObfuscationRule, nil
}

type LogConfigurationsCreateObfuscationRuleQueryResponse struct {
	LogConfigurationsObfuscationRule LogConfigurationsObfuscationRule `json:"LogConfigurationsCreateObfuscationRule"`
}

const LogConfigurationsCreateObfuscationRuleMutation = `mutation(
	$accountId: Int!,
	$rule: LogConfigurationsCreateObfuscationRuleInput!,
) { logConfigurationsCreateObfuscationRule(
	accountId: $accountId,
	rule: $rule,
) {
	actions {
		attributes
		expression {
			createdAt
			createdBy {
				email
				gravatar
				id
				name
			}
			description
			id
			name
			regex
			updatedAt
			updatedBy {
				email
				gravatar
				id
				name
			}
		}
		id
		method
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	enabled
	filter
	id
	name
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Create a new parsing rule.
func (a *Logconfigurations) LogConfigurationsCreateParsingRule(
	accountID int,
	rule LogConfigurationsParsingRuleConfiguration,
) (*LogConfigurationsCreateParsingRuleResponse, error) {
	return a.LogConfigurationsCreateParsingRuleWithContext(context.Background(),
		accountID,
		rule,
	)
}

// Create a new parsing rule.
func (a *Logconfigurations) LogConfigurationsCreateParsingRuleWithContext(
	ctx context.Context,
	accountID int,
	rule LogConfigurationsParsingRuleConfiguration,
) (*LogConfigurationsCreateParsingRuleResponse, error) {

	resp := LogConfigurationsCreateParsingRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"rule":      rule,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsCreateParsingRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsCreateParsingRuleResponse, nil
}

type LogConfigurationsCreateParsingRuleQueryResponse struct {
	LogConfigurationsCreateParsingRuleResponse LogConfigurationsCreateParsingRuleResponse `json:"LogConfigurationsCreateParsingRule"`
}

const LogConfigurationsCreateParsingRuleMutation = `mutation(
	$accountId: Int!,
	$rule: LogConfigurationsParsingRuleConfiguration!,
) { logConfigurationsCreateParsingRule(
	accountId: $accountId,
	rule: $rule,
) {
	errors {
		message
		type
	}
	rule {
		accountId
		attribute
		createdBy {
			email
			gravatar
			id
			name
		}
		deleted
		description
		enabled
		grok
		id
		lucene
		nrql
		updatedAt
		updatedBy {
			email
			gravatar
			id
			name
		}
	}
} }`

// Delete an existing data partition rule.
// This operation will result in data to be allocated in the main NRDB storage (Log) if no other data partition rule exists intercepting the logs matching this rule.
// A deleted data partition rule can be recreated using the same name.
func (a *Logconfigurations) LogConfigurationsDeleteDataPartitionRule(
	accountID int,
	iD string,
) (*LogConfigurationsDeleteDataPartitionRuleResponse, error) {
	return a.LogConfigurationsDeleteDataPartitionRuleWithContext(context.Background(),
		accountID,
		iD,
	)
}

// Delete an existing data partition rule.
// This operation will result in data to be allocated in the main NRDB storage (Log) if no other data partition rule exists intercepting the logs matching this rule.
// A deleted data partition rule can be recreated using the same name.
func (a *Logconfigurations) LogConfigurationsDeleteDataPartitionRuleWithContext(
	ctx context.Context,
	accountID int,
	iD string,
) (*LogConfigurationsDeleteDataPartitionRuleResponse, error) {

	resp := LogConfigurationsDeleteDataPartitionRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"id":        iD,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsDeleteDataPartitionRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsDeleteDataPartitionRuleResponse, nil
}

type LogConfigurationsDeleteDataPartitionRuleQueryResponse struct {
	LogConfigurationsDeleteDataPartitionRuleResponse LogConfigurationsDeleteDataPartitionRuleResponse `json:"LogConfigurationsDeleteDataPartitionRule"`
}

const LogConfigurationsDeleteDataPartitionRuleMutation = `mutation(
	$accountId: Int!,
	$id: ID!,
) { logConfigurationsDeleteDataPartitionRule(
	accountId: $accountId,
	id: $id,
) {
	errors {
		message
		type
	}
} }`

// Delete an obfuscation expression.
func (a *Logconfigurations) LogConfigurationsDeleteObfuscationExpression(
	accountID int,
	iD string,
) (*LogConfigurationsObfuscationExpression, error) {
	return a.LogConfigurationsDeleteObfuscationExpressionWithContext(context.Background(),
		accountID,
		iD,
	)
}

// Delete an obfuscation expression.
func (a *Logconfigurations) LogConfigurationsDeleteObfuscationExpressionWithContext(
	ctx context.Context,
	accountID int,
	iD string,
) (*LogConfigurationsObfuscationExpression, error) {

	resp := LogConfigurationsDeleteObfuscationExpressionQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"id":        iD,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsDeleteObfuscationExpressionMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsObfuscationExpression, nil
}

type LogConfigurationsDeleteObfuscationExpressionQueryResponse struct {
	LogConfigurationsObfuscationExpression LogConfigurationsObfuscationExpression `json:"LogConfigurationsDeleteObfuscationExpression"`
}

const LogConfigurationsDeleteObfuscationExpressionMutation = `mutation(
	$accountId: Int!,
	$id: ID!,
) { logConfigurationsDeleteObfuscationExpression(
	accountId: $accountId,
	id: $id,
) {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	id
	name
	regex
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Delete an obfuscation rule.
func (a *Logconfigurations) LogConfigurationsDeleteObfuscationRule(
	accountID int,
	iD string,
) (*LogConfigurationsObfuscationRule, error) {
	return a.LogConfigurationsDeleteObfuscationRuleWithContext(context.Background(),
		accountID,
		iD,
	)
}

// Delete an obfuscation rule.
func (a *Logconfigurations) LogConfigurationsDeleteObfuscationRuleWithContext(
	ctx context.Context,
	accountID int,
	iD string,
) (*LogConfigurationsObfuscationRule, error) {

	resp := LogConfigurationsDeleteObfuscationRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"id":        iD,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsDeleteObfuscationRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsObfuscationRule, nil
}

type LogConfigurationsDeleteObfuscationRuleQueryResponse struct {
	LogConfigurationsObfuscationRule LogConfigurationsObfuscationRule `json:"LogConfigurationsDeleteObfuscationRule"`
}

const LogConfigurationsDeleteObfuscationRuleMutation = `mutation(
	$accountId: Int!,
	$id: ID!,
) { logConfigurationsDeleteObfuscationRule(
	accountId: $accountId,
	id: $id,
) {
	actions {
		attributes
		expression {
			createdAt
			createdBy {
				email
				gravatar
				id
				name
			}
			description
			id
			name
			regex
			updatedAt
			updatedBy {
				email
				gravatar
				id
				name
			}
		}
		id
		method
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	enabled
	filter
	id
	name
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Delete an existing parsing rule.
func (a *Logconfigurations) LogConfigurationsDeleteParsingRule(
	accountID int,
	iD string,
) (*LogConfigurationsDeleteParsingRuleResponse, error) {
	return a.LogConfigurationsDeleteParsingRuleWithContext(context.Background(),
		accountID,
		iD,
	)
}

// Delete an existing parsing rule.
func (a *Logconfigurations) LogConfigurationsDeleteParsingRuleWithContext(
	ctx context.Context,
	accountID int,
	iD string,
) (*LogConfigurationsDeleteParsingRuleResponse, error) {

	resp := LogConfigurationsDeleteParsingRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"id":        iD,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsDeleteParsingRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsDeleteParsingRuleResponse, nil
}

type LogConfigurationsDeleteParsingRuleQueryResponse struct {
	LogConfigurationsDeleteParsingRuleResponse LogConfigurationsDeleteParsingRuleResponse `json:"LogConfigurationsDeleteParsingRule"`
}

const LogConfigurationsDeleteParsingRuleMutation = `mutation(
	$accountId: Int!,
	$id: ID!,
) { logConfigurationsDeleteParsingRule(
	accountId: $accountId,
	id: $id,
) {
	errors {
		message
		type
	}
} }`

// Update an existing data partition rule.
func (a *Logconfigurations) LogConfigurationsUpdateDataPartitionRule(
	accountID int,
	rule LogConfigurationsUpdateDataPartitionRuleInput,
) (*LogConfigurationsUpdateDataPartitionRuleResponse, error) {
	return a.LogConfigurationsUpdateDataPartitionRuleWithContext(context.Background(),
		accountID,
		rule,
	)
}

// Update an existing data partition rule.
func (a *Logconfigurations) LogConfigurationsUpdateDataPartitionRuleWithContext(
	ctx context.Context,
	accountID int,
	rule LogConfigurationsUpdateDataPartitionRuleInput,
) (*LogConfigurationsUpdateDataPartitionRuleResponse, error) {

	resp := LogConfigurationsUpdateDataPartitionRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"rule":      rule,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsUpdateDataPartitionRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsUpdateDataPartitionRuleResponse, nil
}

type LogConfigurationsUpdateDataPartitionRuleQueryResponse struct {
	LogConfigurationsUpdateDataPartitionRuleResponse LogConfigurationsUpdateDataPartitionRuleResponse `json:"LogConfigurationsUpdateDataPartitionRule"`
}

const LogConfigurationsUpdateDataPartitionRuleMutation = `mutation(
	$accountId: Int!,
	$rule: LogConfigurationsUpdateDataPartitionRuleInput,
) { logConfigurationsUpdateDataPartitionRule(
	accountId: $accountId,
	rule: $rule,
) {
	errors {
		message
		type
	}
	rule {
		createdAt
		createdBy {
			email
			gravatar
			id
			name
		}
		deleted
		description
		enabled
		id
		matchingCriteria {
			attributeName
			matchingExpression
			matchingOperator
		}
		nrql
		retentionPolicy
		targetDataPartition
		updatedAt
		updatedBy {
			email
			gravatar
			id
			name
		}
	}
} }`

// Update an existing data partition rule.
func (a *Logconfigurations) LogConfigurationsUpdateObfuscationExpression(
	accountID int,
	expression LogConfigurationsUpdateObfuscationExpressionInput,
) (*LogConfigurationsObfuscationExpression, error) {
	return a.LogConfigurationsUpdateObfuscationExpressionWithContext(context.Background(),
		accountID,
		expression,
	)
}

// Update an existing data partition rule.
func (a *Logconfigurations) LogConfigurationsUpdateObfuscationExpressionWithContext(
	ctx context.Context,
	accountID int,
	expression LogConfigurationsUpdateObfuscationExpressionInput,
) (*LogConfigurationsObfuscationExpression, error) {

	resp := LogConfigurationsUpdateObfuscationExpressionQueryResponse{}
	vars := map[string]interface{}{
		"accountId":  accountID,
		"expression": expression,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsUpdateObfuscationExpressionMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsObfuscationExpression, nil
}

type LogConfigurationsUpdateObfuscationExpressionQueryResponse struct {
	LogConfigurationsObfuscationExpression LogConfigurationsObfuscationExpression `json:"LogConfigurationsUpdateObfuscationExpression"`
}

const LogConfigurationsUpdateObfuscationExpressionMutation = `mutation(
	$accountId: Int!,
	$expression: LogConfigurationsUpdateObfuscationExpressionInput!,
) { logConfigurationsUpdateObfuscationExpression(
	accountId: $accountId,
	expression: $expression,
) {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	id
	name
	regex
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Update an existing data partition rule.
func (a *Logconfigurations) LogConfigurationsUpdateObfuscationRule(
	accountID int,
	rule LogConfigurationsUpdateObfuscationRuleInput,
) (*LogConfigurationsObfuscationRule, error) {
	return a.LogConfigurationsUpdateObfuscationRuleWithContext(context.Background(),
		accountID,
		rule,
	)
}

// Update an existing data partition rule.
func (a *Logconfigurations) LogConfigurationsUpdateObfuscationRuleWithContext(
	ctx context.Context,
	accountID int,
	rule LogConfigurationsUpdateObfuscationRuleInput,
) (*LogConfigurationsObfuscationRule, error) {

	resp := LogConfigurationsUpdateObfuscationRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"rule":      rule,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsUpdateObfuscationRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsObfuscationRule, nil
}

type LogConfigurationsUpdateObfuscationRuleQueryResponse struct {
	LogConfigurationsObfuscationRule LogConfigurationsObfuscationRule `json:"LogConfigurationsUpdateObfuscationRule"`
}

const LogConfigurationsUpdateObfuscationRuleMutation = `mutation(
	$accountId: Int!,
	$rule: LogConfigurationsUpdateObfuscationRuleInput!,
) { logConfigurationsUpdateObfuscationRule(
	accountId: $accountId,
	rule: $rule,
) {
	actions {
		attributes
		expression {
			createdAt
			createdBy {
				email
				gravatar
				id
				name
			}
			description
			id
			name
			regex
			updatedAt
			updatedBy {
				email
				gravatar
				id
				name
			}
		}
		id
		method
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	enabled
	filter
	id
	name
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} }`

// Update an existing parsing rule.
func (a *Logconfigurations) LogConfigurationsUpdateParsingRule(
	accountID int,
	iD string,
	rule LogConfigurationsParsingRuleConfiguration,
) (*LogConfigurationsUpdateParsingRuleResponse, error) {
	return a.LogConfigurationsUpdateParsingRuleWithContext(context.Background(),
		accountID,
		iD,
		rule,
	)
}

// Update an existing parsing rule.
func (a *Logconfigurations) LogConfigurationsUpdateParsingRuleWithContext(
	ctx context.Context,
	accountID int,
	iD string,
	rule LogConfigurationsParsingRuleConfiguration,
) (*LogConfigurationsUpdateParsingRuleResponse, error) {

	resp := LogConfigurationsUpdateParsingRuleQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"id":        iD,
		"rule":      rule,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, LogConfigurationsUpdateParsingRuleMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.LogConfigurationsUpdateParsingRuleResponse, nil
}

type LogConfigurationsUpdateParsingRuleQueryResponse struct {
	LogConfigurationsUpdateParsingRuleResponse LogConfigurationsUpdateParsingRuleResponse `json:"LogConfigurationsUpdateParsingRule"`
}

const LogConfigurationsUpdateParsingRuleMutation = `mutation(
	$accountId: Int!,
	$id: ID!,
	$rule: LogConfigurationsParsingRuleConfiguration!,
) { logConfigurationsUpdateParsingRule(
	accountId: $accountId,
	id: $id,
	rule: $rule,
) {
	errors {
		message
		type
	}
	rule {
		accountId
		attribute
		createdBy {
			email
			gravatar
			id
			name
		}
		deleted
		description
		enabled
		grok
		id
		lucene
		nrql
		updatedAt
		updatedBy {
			email
			gravatar
			id
			name
		}
	}
} }`

// Look up for all data partition rules for a given account.
func (a *Logconfigurations) GetDataPartitionRules(
	accountID int,
) (*[]LogConfigurationsDataPartitionRule, error) {
	return a.GetDataPartitionRulesWithContext(context.Background(),
		accountID,
	)
}

// Look up for all data partition rules for a given account.
func (a *Logconfigurations) GetDataPartitionRulesWithContext(
	ctx context.Context,
	accountID int,
) (*[]LogConfigurationsDataPartitionRule, error) {

	resp := dataPartitionRulesResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getDataPartitionRulesQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Account.LogConfigurations.DataPartitionRules) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Account.LogConfigurations.DataPartitionRules, nil
}

const getDataPartitionRulesQuery = `query(
	$accountID: Int!,
) { actor { account(id: $accountID) { logConfigurations { dataPartitionRules {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	deleted
	description
	enabled
	id
	matchingCriteria {
		attributeName
		matchingExpression
		matchingOperator
	}
	nrql
	retentionPolicy
	targetDataPartition
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} } } } }`

// Look up for all obfuscation expressions for a given account
func (a *Logconfigurations) GetObfuscationExpressions(
	accountID int,
) (*[]LogConfigurationsObfuscationExpression, error) {
	return a.GetObfuscationExpressionsWithContext(context.Background(),
		accountID,
	)
}

// Look up for all obfuscation expressions for a given account
func (a *Logconfigurations) GetObfuscationExpressionsWithContext(
	ctx context.Context,
	accountID int,
) (*[]LogConfigurationsObfuscationExpression, error) {

	resp := obfuscationExpressionsResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getObfuscationExpressionsQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Account.LogConfigurations.ObfuscationExpressions) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Account.LogConfigurations.ObfuscationExpressions, nil
}

const getObfuscationExpressionsQuery = `query(
	$accountID: Int!,
) { actor { account(id: $accountID) { logConfigurations { obfuscationExpressions {
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	id
	name
	regex
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} } } } }`

// Look up for all obfuscation rules for a given account.
func (a *Logconfigurations) GetObfuscationRules(
	accountID int,
) (*[]LogConfigurationsObfuscationRule, error) {
	return a.GetObfuscationRulesWithContext(context.Background(),
		accountID,
	)
}

// Look up for all obfuscation rules for a given account.
func (a *Logconfigurations) GetObfuscationRulesWithContext(
	ctx context.Context,
	accountID int,
) (*[]LogConfigurationsObfuscationRule, error) {

	resp := obfuscationRulesResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getObfuscationRulesQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Account.LogConfigurations.ObfuscationRules) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Account.LogConfigurations.ObfuscationRules, nil
}

const getObfuscationRulesQuery = `query(
	$accountID: Int!,
) { actor { account(id: $accountID) { logConfigurations { obfuscationRules {
	actions {
		attributes
		expression {
			createdAt
			createdBy {
				email
				gravatar
				id
				name
			}
			description
			id
			name
			regex
			updatedAt
			updatedBy {
				email
				gravatar
				id
				name
			}
		}
		id
		method
	}
	createdAt
	createdBy {
		email
		gravatar
		id
		name
	}
	description
	enabled
	filter
	id
	name
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} } } } }`

// Look up for all parsing rules for a given account.
func (a *Logconfigurations) GetParsingRules(
	accountID int,
) (*[]*LogConfigurationsParsingRule, error) {
	return a.GetParsingRulesWithContext(context.Background(),
		accountID,
	)
}

// Look up for all parsing rules for a given account.
func (a *Logconfigurations) GetParsingRulesWithContext(
	ctx context.Context,
	accountID int,
) (*[]*LogConfigurationsParsingRule, error) {

	resp := parsingRulesResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getParsingRulesQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Account.LogConfigurations.ParsingRules) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Account.LogConfigurations.ParsingRules, nil
}

const getParsingRulesQuery = `query(
	$accountID: Int!,
) { actor { account(id: $accountID) { logConfigurations { parsingRules {
	accountId
	attribute
	createdBy {
		email
		gravatar
		id
		name
	}
	deleted
	description
	enabled
	grok
	id
	lucene
	nrql
	updatedAt
	updatedBy {
		email
		gravatar
		id
		name
	}
} } } } }`

// Test a Grok pattern against a list of log lines.
func (a *Logconfigurations) GetTestGrok(
	accountID int,
	grok string,
	logLines []string,
) (*[]LogConfigurationsGrokTestResult, error) {
	return a.GetTestGrokWithContext(context.Background(),
		accountID,
		grok,
		logLines,
	)
}

// Test a Grok pattern against a list of log lines.
func (a *Logconfigurations) GetTestGrokWithContext(
	ctx context.Context,
	accountID int,
	grok string,
	logLines []string,
) (*[]LogConfigurationsGrokTestResult, error) {

	resp := testGrokResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
		"grok":      grok,
		"logLines":  logLines,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getTestGrokQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Account.LogConfigurations.TestGrok) == 0 {
		return nil, errors.NewNotFound("")
	}

	return &resp.Actor.Account.LogConfigurations.TestGrok, nil
}

const getTestGrokQuery = `query(
	$accountID: Int!,
	$grok: String!,
	$logLines: [String!]!,
) { actor { account(id: $accountID) { logConfigurations { testGrok(
	grok: $grok,
	logLines: $logLines,
) {
	attributes {
		name
		value
	}
	logLine
	matched
} } } } }`