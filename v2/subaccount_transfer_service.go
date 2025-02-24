package binance

import (
	"context"
	"encoding/json"
)

type CreateSubAccountTransferService struct {
	c               *Client
	fromEmail       *string
	toEmail         *string
	fromAccountType string
	toAccountType   string
	asset           string
	amount          string
}

func (s *CreateSubAccountTransferService) Asset(v string) *CreateSubAccountTransferService {
	s.asset = v
	return s
}

func (s *CreateSubAccountTransferService) Amount(v string) *CreateSubAccountTransferService {
	s.amount = v
	return s
}

func (s *CreateSubAccountTransferService) FromEmail(v string) *CreateSubAccountTransferService {
	s.fromEmail = &v
	return s
}

func (s *CreateSubAccountTransferService) ToEmail(v string) *CreateSubAccountTransferService {
	s.toEmail = &v
	return s
}

func (s *CreateSubAccountTransferService) FromAccountType(v string) *CreateSubAccountTransferService {
	s.fromAccountType = v
	return s
}
func (s *CreateSubAccountTransferService) ToAccountType(v string) *CreateSubAccountTransferService {
	s.toAccountType = v
	return s
}

// Do sends the request.
func (s *CreateSubAccountTransferService) Do(ctx context.Context) (*CreateSubAccountTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	m := params{
		"fromAccountType": s.fromAccountType,
		"toAccountType":   s.toAccountType,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	if s.fromEmail != nil {
		m["fromEmail"] = *s.fromEmail
	}

	if s.toEmail != nil {
		m["toEmail"] = *s.toEmail
	}

	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateSubAccountTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil

}

// CreateWithdrawResponse represents a response from CreateWithdrawService.
type CreateSubAccountTransferResponse struct {
	TranId int64 `json:"tranId"`
}
