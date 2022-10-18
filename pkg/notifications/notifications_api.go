// Code generated by tutone: DO NOT EDIT
package notifications

import (
	"context"

	"github.com/newrelic/newrelic-client-go/v2/pkg/ai"
)

// Create a Channel
func (a *Notifications) AiNotificationsCreateChannel(
	accountID int,
	channel AiNotificationsChannelInput,
) (*AiNotificationsChannelResponse, error) {
	return a.AiNotificationsCreateChannelWithContext(context.Background(),
		accountID,
		channel,
	)
}

// Create a Channel
func (a *Notifications) AiNotificationsCreateChannelWithContext(
	ctx context.Context,
	accountID int,
	channel AiNotificationsChannelInput,
) (*AiNotificationsChannelResponse, error) {

	resp := AiNotificationsCreateChannelQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"channel":   channel,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, AiNotificationsCreateChannelMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.AiNotificationsChannelResponse, nil
}

type AiNotificationsCreateChannelQueryResponse struct {
	AiNotificationsChannelResponse AiNotificationsChannelResponse `json:"AiNotificationsCreateChannel"`
}

const AiNotificationsCreateChannelMutation = `mutation(
	$accountId: Int!,
	$channel: AiNotificationsChannelInput!,
) { aiNotificationsCreateChannel(
	accountId: $accountId,
	channel: $channel,
) {
	channel {
		accountId
		active
		createdAt
		destinationId
		id
		name
		product
		properties {
			displayValue
			key
			label
			value
		}
		status
		type
		updatedAt
		updatedBy
	}
    error {
      ... on AiNotificationsConstraintsError {
        constraints {
          dependencies
          name
        }
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          field
          message
        }
      }
      ... on AiNotificationsSuggestionError {
        description
        details
        type
      }
      ... on AiNotificationsResponseError {
        description
        details
        type
      }
    }
    errors {
      ... on AiNotificationsSuggestionError {
        description
        type
        details
      }
      ... on AiNotificationsResponseError {
        description
        type
        details
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          message
          field
        }
      }
      ... on AiNotificationsConstraintsError {
        constraints {
          dependencies
          name
        }
      }
    }
} }`

// Create a Destination
func (a *Notifications) AiNotificationsCreateDestination(
	accountID int,
	destination AiNotificationsDestinationInput,
) (*AiNotificationsDestinationResponse, error) {
	return a.AiNotificationsCreateDestinationWithContext(context.Background(),
		accountID,
		destination,
	)
}

// Create a Destination
func (a *Notifications) AiNotificationsCreateDestinationWithContext(
	ctx context.Context,
	accountID int,
	destination AiNotificationsDestinationInput,
) (*AiNotificationsDestinationResponse, error) {

	resp := AiNotificationsCreateDestinationQueryResponse{}
	vars := map[string]interface{}{
		"accountId":   accountID,
		"destination": destination,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, AiNotificationsCreateDestinationMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.AiNotificationsDestinationResponse, nil
}

type AiNotificationsCreateDestinationQueryResponse struct {
	AiNotificationsDestinationResponse AiNotificationsDestinationResponse `json:"AiNotificationsCreateDestination"`
}

const AiNotificationsCreateDestinationMutation = `mutation(
	$accountId: Int!,
	$destination: AiNotificationsDestinationInput!,
) { aiNotificationsCreateDestination(
	accountId: $accountId,
	destination: $destination,
) {
	destination {
		accountId
		active
		auth {
			... on AiNotificationsBasicAuth {
			  authType
			  user
			}
			... on AiNotificationsOAuth2Auth {
			  accessTokenUrl
			  scope
			  refreshable
			  refreshInterval
			  prefix
			  clientId
			  authorizationUrl
			  authType
			}
			... on AiNotificationsTokenAuth {
			  authType
			  prefix
			}
		}
		createdAt
		id
		isUserAuthenticated
		lastSent
		name
		properties {
			displayValue
			key
			label
			value
		}
		status
		type
		updatedAt
		updatedBy
	}
	errors {
      ... on AiNotificationsConstraintsError {
        constraints {
          dependencies
          name
        }
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          field
          message
        }
      }
      ... on AiNotificationsResponseError {
        description
        details
        type
      }
      ... on AiNotificationsSuggestionError {
        description
        type
        details
      }
    }
    error {
      ... on AiNotificationsSuggestionError {
        description
        type
        details
      }
      ... on AiNotificationsResponseError {
        description
        type
        details
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          message
          field
        }
      }
      ... on AiNotificationsConstraintsError {
        constraints {
          name
          dependencies
        }
      }
    }
} }`

// Delete a Channel
func (a *Notifications) AiNotificationsDeleteChannel(
	accountID int,
	channelId string,
) (*AiNotificationsDeleteResponse, error) {
	return a.AiNotificationsDeleteChannelWithContext(context.Background(),
		accountID,
		channelId,
	)
}

// Delete a Channel
func (a *Notifications) AiNotificationsDeleteChannelWithContext(
	ctx context.Context,
	accountID int,
	channelId string,
) (*AiNotificationsDeleteResponse, error) {

	resp := AiNotificationsDeleteChannelQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"channelId": channelId,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, AiNotificationsDeleteChannelMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.AiNotificationsDeleteResponse, nil
}

type AiNotificationsDeleteChannelQueryResponse struct {
	AiNotificationsDeleteResponse AiNotificationsDeleteResponse `json:"AiNotificationsDeleteChannel"`
}

const AiNotificationsDeleteChannelMutation = `mutation(
	$accountId: Int!,
	$channelId: ID!,
) { aiNotificationsDeleteChannel(
	accountId: $accountId,
	channelId: $channelId,
) {
	error {
		description
		details
		type
	}
	errors {
		description
		details
		type
	}
	ids
} }`

// Delete a Destination
func (a *Notifications) AiNotificationsDeleteDestination(
	accountID int,
	destinationId string,
) (*AiNotificationsDeleteResponse, error) {
	return a.AiNotificationsDeleteDestinationWithContext(context.Background(),
		accountID,
		destinationId,
	)
}

// Delete a Destination
func (a *Notifications) AiNotificationsDeleteDestinationWithContext(
	ctx context.Context,
	accountID int,
	destinationId string,
) (*AiNotificationsDeleteResponse, error) {

	resp := AiNotificationsDeleteDestinationQueryResponse{}
	vars := map[string]interface{}{
		"accountId":     accountID,
		"destinationId": destinationId,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, AiNotificationsDeleteDestinationMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.AiNotificationsDeleteResponse, nil
}

type AiNotificationsDeleteDestinationQueryResponse struct {
	AiNotificationsDeleteResponse AiNotificationsDeleteResponse `json:"AiNotificationsDeleteDestination"`
}

const AiNotificationsDeleteDestinationMutation = `mutation(
	$accountId: Int!,
	$destinationId: ID!,
) { aiNotificationsDeleteDestination(
	accountId: $accountId,
	destinationId: $destinationId,
) {
	error {
		description
		details
		type
	}
	errors {
		description
		details
		type
	}
	ids
} }`

// Update a Channel
func (a *Notifications) AiNotificationsUpdateChannel(
	accountID int,
	channel AiNotificationsChannelUpdate,
	channelId string,
) (*AiNotificationsChannelResponse, error) {
	return a.AiNotificationsUpdateChannelWithContext(context.Background(),
		accountID,
		channel,
		channelId,
	)
}

// Update a Channel
func (a *Notifications) AiNotificationsUpdateChannelWithContext(
	ctx context.Context,
	accountID int,
	channel AiNotificationsChannelUpdate,
	channelId string,
) (*AiNotificationsChannelResponse, error) {

	resp := AiNotificationsUpdateChannelQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"channel":   channel,
		"channelId": channelId,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, AiNotificationsUpdateChannelMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.AiNotificationsChannelResponse, nil
}

type AiNotificationsUpdateChannelQueryResponse struct {
	AiNotificationsChannelResponse AiNotificationsChannelResponse `json:"AiNotificationsUpdateChannel"`
}

const AiNotificationsUpdateChannelMutation = `mutation(
	$accountId: Int!,
	$channel: AiNotificationsChannelUpdate!,
	$channelId: ID!,
) { aiNotificationsUpdateChannel(
	accountId: $accountId,
	channel: $channel,
	channelId: $channelId,
) {
	channel {
		accountId
		active
		createdAt
		destinationId
		id
		name
		product
		properties {
			displayValue
			key
			label
			value
		}
		status
		type
		updatedAt
		updatedBy
	}
    error {
      ... on AiNotificationsConstraintsError {
        constraints {
          dependencies
          name
        }
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          field
          message
        }
      }
      ... on AiNotificationsSuggestionError {
        description
        details
        type
      }
      ... on AiNotificationsResponseError {
        description
        details
        type
      }
    }
    errors {
      ... on AiNotificationsSuggestionError {
        description
        type
        details
      }
      ... on AiNotificationsResponseError {
        description
        type
        details
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          message
          field
        }
      }
      ... on AiNotificationsConstraintsError {
        constraints {
          dependencies
          name
        }
      }
    }
} }`

// Update a Destination
func (a *Notifications) AiNotificationsUpdateDestination(
	accountID int,
	destination AiNotificationsDestinationUpdate,
	destinationId string,
) (*AiNotificationsDestinationResponse, error) {
	return a.AiNotificationsUpdateDestinationWithContext(context.Background(),
		accountID,
		destination,
		destinationId,
	)
}

// Update a Destination
func (a *Notifications) AiNotificationsUpdateDestinationWithContext(
	ctx context.Context,
	accountID int,
	destination AiNotificationsDestinationUpdate,
	destinationId string,
) (*AiNotificationsDestinationResponse, error) {

	resp := AiNotificationsUpdateDestinationQueryResponse{}
	vars := map[string]interface{}{
		"accountId":     accountID,
		"destination":   destination,
		"destinationId": destinationId,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, AiNotificationsUpdateDestinationMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.AiNotificationsDestinationResponse, nil
}

type AiNotificationsUpdateDestinationQueryResponse struct {
	AiNotificationsDestinationResponse AiNotificationsDestinationResponse `json:"AiNotificationsUpdateDestination"`
}

const AiNotificationsUpdateDestinationMutation = `mutation(
	$accountId: Int!,
	$destination: AiNotificationsDestinationUpdate!,
	$destinationId: ID!,
) { aiNotificationsUpdateDestination(
	accountId: $accountId,
	destination: $destination,
	destinationId: $destinationId,
) {
	destination {
		accountId
		active
		auth {
			... on AiNotificationsBasicAuth {
			  authType
			  user
			}
			... on AiNotificationsOAuth2Auth {
			  accessTokenUrl
			  scope
			  refreshable
			  refreshInterval
			  prefix
			  clientId
			  authorizationUrl
			  authType
			}
			... on AiNotificationsTokenAuth {
			  authType
			  prefix
			}
		}
		createdAt
		id
		isUserAuthenticated
		lastSent
		name
		properties {
			displayValue
			key
			label
			value
		}
		status
		type
		updatedAt
		updatedBy
	}
	errors {
      ... on AiNotificationsConstraintsError {
        constraints {
          dependencies
          name
        }
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          field
          message
        }
      }
      ... on AiNotificationsResponseError {
        description
        details
        type
      }
      ... on AiNotificationsSuggestionError {
        description
        type
        details
      }
    }
    error {
      ... on AiNotificationsSuggestionError {
        description
        type
        details
      }
      ... on AiNotificationsResponseError {
        description
        type
        details
      }
      ... on AiNotificationsDataValidationError {
        details
        fields {
          message
          field
        }
      }
      ... on AiNotificationsConstraintsError {
        constraints {
          name
          dependencies
        }
      }
    }
} }`

// Fetch a Channel by product
func (a *Notifications) GetChannels(
	accountID int,
	cursor string,
	filters ai.AiNotificationsChannelFilter,
	sorter AiNotificationsChannelSorter,
) (*AiNotificationsChannelsResponse, error) {
	return a.GetChannelsWithContext(context.Background(),
		accountID,
		cursor,
		filters,
		sorter,
	)
}

// Fetch a Channel by product
func (a *Notifications) GetChannelsWithContext(
	ctx context.Context,
	accountID int,
	cursor string,
	filters ai.AiNotificationsChannelFilter,
	sorter AiNotificationsChannelSorter,
) (*AiNotificationsChannelsResponse, error) {

	resp := channelsResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
		"cursor":    cursor,
		"filters":   filters,
		"sorter":    sorter,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getChannelsQuery, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.Actor.Account.AiNotifications.Channels, nil
}

const getChannelsQuery = `query(
	$accountID: Int!, $filters: AiNotificationsChannelFilter,
) { actor { account(id: $accountID) { aiNotifications { channels(filters: $filters) {
	entities {
		accountId
		active
		createdAt
		destinationId
		id
		name
		product
		properties {
			displayValue
			key
			label
			value
		}
		status
		type
		updatedAt
		updatedBy
	}
	error {
		description
		details
		type
	}
	errors {
		description
		details
		type
	}
	nextCursor
	totalCount
} } } } }`

// Fetch a Destinations by type
func (a *Notifications) GetDestinations(
	accountID int,
	cursor string,
	filters ai.AiNotificationsDestinationFilter,
	sorter AiNotificationsDestinationSorter,
) (*AiNotificationsDestinationsResponse, error) {
	return a.GetDestinationsWithContext(context.Background(),
		accountID,
		cursor,
		filters,
		sorter,
	)
}

// Fetch a Destinations by type
func (a *Notifications) GetDestinationsWithContext(
	ctx context.Context,
	accountID int,
	cursor string,
	filters ai.AiNotificationsDestinationFilter,
	sorter AiNotificationsDestinationSorter,
) (*AiNotificationsDestinationsResponse, error) {

	resp := destinationsResponse{}
	vars := map[string]interface{}{
		"accountID": accountID,
		"cursor":    cursor,
		"filters":   filters,
		"sorter":    sorter,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getDestinationsQuery, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.Actor.Account.AiNotifications.Destinations, nil
}

const getDestinationsQuery = `query(
	$accountID: Int!, $filters: AiNotificationsDestinationFilter,
) { actor { account(id: $accountID) { aiNotifications { destinations(filters: $filters) {
	entities {
		accountId
		active
		auth {
			... on AiNotificationsBasicAuth {
			  authType
			  user
			}
			... on AiNotificationsOAuth2Auth {
			  accessTokenUrl
			  scope
			  refreshable
			  refreshInterval
			  prefix
			  clientId
			  authorizationUrl
			  authType
			}
			... on AiNotificationsTokenAuth {
			  authType
			  prefix
			}
		}
		createdAt
		id
		isUserAuthenticated
		lastSent
		name
		properties {
			displayValue
			key
			label
			value
		}
		status
		type
		updatedAt
		updatedBy
	}
	error {
		description
		details
		type
	}
	errors {
		description
		details
		type
	}
	nextCursor
	totalCount
} } } } }`
