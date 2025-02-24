package binance

import (
	"context"
	"encoding/json"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetAssetDetailService struct {
	c     *Client
	asset string
}

func (s *GetAssetDetailService) Asset(v string) *GetAssetDetailService {
	s.asset = v
	return s
}

// Do sends the request.
func (s *GetAssetDetailService) Do(ctx context.Context) ([]*AssetDetail, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return []*AssetDetail{}, err
	}

	res := make([]*AssetDetail, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AssetDetail{}, err
	}
	return res, nil

}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`
	Locked       string `json:"locked"`
	Freeze       string `json:"freeze"`
	Withdrawing  string `json:"withdrawing"`
	BtcValuation string `json:"btcValuation"`
}

type CreateAssetTransferService struct {
	c      *Client
	ttype  string
	asset  string
	amount string
}

func (s *CreateAssetTransferService) Asset(v string) *CreateAssetTransferService {
	s.asset = v
	return s
}

func (s *CreateAssetTransferService) Type(v string) *CreateAssetTransferService {
	s.ttype = v
	return s
}

func (s *CreateAssetTransferService) Amount(v string) *CreateAssetTransferService {
	s.amount = v
	return s
}

// Do sends the request.
func (s *CreateAssetTransferService) Do(ctx context.Context) (*CreateAssetTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}
	r.setParam("type", s.ttype)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateAssetTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil

}

// CreateWithdrawResponse represents a response from CreateWithdrawService.
type CreateAssetTransferResponse struct {
	TranId int64 `json:"tranId"`
}
