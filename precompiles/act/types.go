package act

import (
	"math/big"
	"time"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"

	precommon "github.com/warden-protocol/wardenprotocol/precompiles/common"
	types "github.com/warden-protocol/wardenprotocol/warden/x/act/types/v1beta1"
)

// ActionsInput needed to unmarshal Pagination field and pass it to types.QueryActionsRequest
type ActionsInput struct {
	Pagination query.PageRequest `abi:"pagination"`
}

// FromResponse needed to map QueryActionsResponse to ActionsResponse
func (r *ActionsResponse) FromResponse(res *types.QueryActionsResponse) (ActionsResponse, error) {
	if res != nil {
		actions := make([]Action, 0, len(res.Actions))
		for _, action := range res.Actions {
			mappedAction, err := mapAction(action)
			if err != nil {
				return ActionsResponse{}, err
			}

			actions = append(actions, mappedAction)
		}

		r.Actions = actions
		r.Pagination = mapPageResponse(res.Pagination)
	}

	return *r, nil
}

// FromResponse needed to map QueryActionByIdResponse to ActionByIdResponse
func (r *ActionByIdResponse) FromResponse(res *types.QueryActionByIdResponse) (ActionByIdResponse, error) {
	if res != nil && res.Action != nil {
		mappedAction, err := mapAction(*res.Action)
		if err != nil {
			return ActionByIdResponse{}, err
		}

		r.Action = mappedAction
	}

	return *r, nil
}

// ActionsByAddressInput needed to unmarshal Pagination field and pass it to types.QueryActionsByAddressRequest
type ActionsByAddressInput struct {
	Pagination query.PageRequest `abi:"pagination"`
	Address    common.Address    `abi:"addr"`
	Status     int32             `abi:"status"`
}

// FromResponse needed to map QueryActionsByAddressResponse to ActionsByAddressResponse
func (r *ActionsByAddressResponse) FromResponse(res *types.QueryActionsByAddressResponse) (ActionsByAddressResponse, error) {
	if res != nil {
		actions := make([]Action, 0, len(res.Actions))
		for _, action := range res.Actions {
			mappedAction, err := mapAction(action)
			if err != nil {
				return ActionsByAddressResponse{}, err
			}

			actions = append(actions, mappedAction)
		}

		r.Actions = actions
		r.Pagination = mapPageResponse(res.Pagination)
	}

	return *r, nil
}

// TemplatesInput needed to unmarshal Pagination field and pass it to types.QueryTemplatesRequest
type TemplatesInput struct {
	Pagination query.PageRequest
	Creator    string `json:"creator"`
}

// FromResponse needed to map QueryTemplatesResponse to TemplatesResponse
func (r *TemplatesResponse) FromResponse(res *types.QueryTemplatesResponse) (TemplatesResponse, error) {
	if res != nil {
		templates := make([]Template, 0, len(res.Templates))
		for _, action := range res.Templates {
			mappedTemplate, err := mapTemplate(action)
			if err != nil {
				return TemplatesResponse{}, err
			}

			templates = append(templates, mappedTemplate)
		}

		r.Templates = templates
		r.Pagination = mapPageResponse(res.Pagination)
	}

	return *r, nil
}

// FromResponse needed to map QueryTemplateByIdResponse to TemplateByIdResponse
func (r *TemplateByIdResponse) FromResponse(res *types.QueryTemplateByIdResponse) (TemplateByIdResponse, error) {
	if res != nil && res.Template != nil {
		mappedTemplate, err := mapTemplate(*res.Template)
		if err != nil {
			return TemplateByIdResponse{}, err
		}

		r.Template = mappedTemplate
	}

	return *r, nil
}

func mapAction(action types.Action) (Action, error) {
	mentions := make([]common.Address, 0, len(action.Mentions))
	for _, mention := range action.Mentions {
		mentionAddress, err := precommon.AddressFromBech32Str(mention)

		if err != nil {
			return Action{}, err
		}

		mentions = append(mentions, mentionAddress)
	}

	creator, err := precommon.AddressFromBech32Str(action.Creator)

	if err != nil {
		return Action{}, err
	}

	mappedVotes, err := mapVotes(action.Votes)
	if err != nil {
		return Action{}, err
	}

	return Action{
		Id:                action.Id,
		Status:            big.NewInt(int64(action.Status)),
		StatusText:        action.Status.String(),
		Msg:               mapAny(action.Msg),
		Result:            mapAny(action.Result),
		Creator:           creator,
		TimeoutHeight:     action.TimeoutHeight,
		CreatedAt:         mapTimestamp(action.CreatedAt),
		UpdatedAt:         mapTimestamp(action.UpdatedAt),
		ApproveExpression: action.ApproveExpression.String(),
		RejectExpression:  action.RejectExpression.String(),
		Mentions:          mentions,
		Votes:             mappedVotes,
	}, nil
}

func mapVotes(values []*types.ActionVote) ([]ActionVote, error) {
	result := make([]ActionVote, 0, len(values))
	for _, v := range values {
		if v != nil {
			mappedTemplate, err := mapVote(*v)
			if err != nil {
				return nil, err
			}

			result = append(result, mappedTemplate)
		}
	}
	return result, nil
}

func mapVote(value types.ActionVote) (ActionVote, error) {
	participant, err := precommon.AddressFromBech32Str(value.Participant)

	if err != nil {
		return ActionVote{}, err
	}

	return ActionVote{
		Participant:  participant,
		VotedAt:      mapTimestamp(value.VotedAt),
		VoteType:     int32(value.VoteType),
		VoteTypeText: value.VoteType.String(),
	}, nil
}

func mapTemplate(value types.Template) (Template, error) {
	creator, err := precommon.AddressFromBech32Str(value.Creator)

	if err != nil {
		return Template{}, err
	}

	return Template{
		Id:         value.Id,
		Creator:    creator,
		Name:       value.Name,
		Expression: value.Expression.String(),
	}, nil
}

func mapAny(any *cdctypes.Any) AnyType {
	if any == nil {
		return AnyType{}
	}

	return AnyType{
		any.TypeUrl,
		any.Value,
	}
}

func mapTimestamp(value time.Time) Timestamp {
	return Timestamp{
		Secs:  uint64(value.Unix()),
		Nanos: uint64(value.Nanosecond()),
	}
}

func mapPageResponse(value *query.PageResponse) TypesPageResponse {
	if value == nil {
		return TypesPageResponse{}
	}

	return TypesPageResponse{
		NextKey: value.NextKey,
		Total:   value.Total,
	}
}