package binance

import (
	"context"
	"encoding/json"
	"github.com/taotao5/go-binance/v2/common"
)

// CreateWithdrawService submits a withdraw request.
//
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw
type CreateWithdrawService struct {
	c                  *Client
	coin               string
	withdrawOrderID    *string
	network            *string
	address            string
	addressTag         *string
	amount             string
	transactionFeeFlag *bool
	name               *string
	walletType         *string
}

// Asset sets the asset parameter (MANDATORY).
func (s *CreateWithdrawService) Coin(v string) *CreateWithdrawService {
	s.coin = v
	return s
}

// WithdrawOrderID sets the withdrawOrderID parameter.
func (s *CreateWithdrawService) WithdrawOrderID(v string) *CreateWithdrawService {
	s.withdrawOrderID = &v
	return s
}

// Network sets the network parameter.
func (s *CreateWithdrawService) Network(v string) *CreateWithdrawService {
	s.network = &v
	return s
}

// Address sets the address parameter (MANDATORY).
func (s *CreateWithdrawService) Address(v string) *CreateWithdrawService {
	s.address = v
	return s
}

// AddressTag sets the addressTag parameter.
func (s *CreateWithdrawService) AddressTag(v string) *CreateWithdrawService {
	s.addressTag = &v
	return s
}

// Amount sets the amount parameter (MANDATORY).
func (s *CreateWithdrawService) Amount(v string) *CreateWithdrawService {
	s.amount = v
	return s
}

// TransactionFeeFlag sets the transactionFeeFlag parameter.
func (s *CreateWithdrawService) TransactionFeeFlag(v bool) *CreateWithdrawService {
	s.transactionFeeFlag = &v
	return s
}

// Name sets the name parameter.
func (s *CreateWithdrawService) Name(v string) *CreateWithdrawService {
	s.name = &v
	return s
}

func (s *CreateWithdrawService) WalletType(v string) *CreateWithdrawService {
	s.walletType = &v
	return s
}

// Do sends the request.
func (s *CreateWithdrawService) Do(ctx context.Context) (*CreateWithdrawResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/capital/withdraw/apply",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if v := s.withdrawOrderID; v != nil {
		r.setParam("withdrawOrderId", *v)
	}
	if v := s.network; v != nil {
		r.setParam("network", *v)
	}
	if v := s.addressTag; v != nil {
		r.setParam("addressTag", *v)
	}
	if v := s.transactionFeeFlag; v != nil {
		r.setParam("transactionFeeFlag", *v)
	}
	if v := s.name; v != nil {
		r.setParam("name", *v)
	}
	if v := s.walletType; v != nil {
		r.setParam("walletType", *v)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateWithdrawResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateWithdrawResponse represents a response from CreateWithdrawService.
type CreateWithdrawResponse struct {
	ID string `json:"id"`
}

// ListWithdrawsService fetches withdraw history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw-history-supporting-network-user_data
type ListWithdrawsService struct {
	c               *Client
	coin            *string
	withdrawOrderId *string
	status          *int
	startTime       *int64
	endTime         *int64
}

// Asset sets the asset parameter.
func (s *ListWithdrawsService) Coin(coin string) *ListWithdrawsService {
	s.coin = &coin
	return s
}

func (s *ListWithdrawsService) WithdrawOrderId(withdrawOrderId string) *ListWithdrawsService {
	s.withdrawOrderId = &withdrawOrderId
	return s
}

// Status sets the status parameter.
func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *ListWithdrawsService) Do(ctx context.Context) (res []*Withdraw, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/withdraw/history",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.withdrawOrderId != nil {
		r.setParam("withdrawOrderId", *s.withdrawOrderId)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make([]*Withdraw, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// Withdraw represents a single withdraw entry.

type Withdraw struct {
	ID              string `json:"id"`
	Amount          string `json:"amount"`
	TransactionFee  string `json:"transactionFee"`
	Coin            string `json:"coin"`
	Status          int    `json:"status"`
	Address         string `json:"address"`
	TxID            string `json:"txId"`
	ApplyTime       string `json:"applyTime"` // UTC时间
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"` // 1:站内转账, 0:站外转账
	WithdrawOrderID string `json:"withdrawOrderId"`
	Info            string `json:"info"` // 提币失败原因
	ConfirmNo       int    `json:"confirmNo"`
	TxKey           string `json:"txKey"`
	CompleteTime    string `json:"completeTime"` // 提现完成，成功下账时间(UTC)
}

type CreateQueryAllCoinService struct {
	c *Client
}

type NetworkList struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc"`
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`
	WithdrawDesc            string `json:"withdrawDesc"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"`
	EstimatedArrivalTime    int    `json:"estimatedArrivalTime"`
	Busy                    bool   `json:"busy"`
}

type CoinData struct {
	Coin              string        `json:"coin"`
	DepositAllEnable  bool          `json:"depositAllEnable"`
	Free              string        `json:"free"`
	Freeze            string        `json:"freeze"`
	Ipoable           string        `json:"ipoable"`
	Ipoing            string        `json:"ipoing"`
	IsLegalMoney      bool          `json:"isLegalMoney"`
	Locked            string        `json:"locked"`
	Name              string        `json:"name"`
	NetworkList       []NetworkList `json:"networkList"`
	Storage           string        `json:"storage"`
	Trading           bool          `json:"trading"`
	WithdrawAllEnable bool          `json:"withdrawAllEnable"`
	Withdrawing       string        `json:"withdrawing"`
}

func (s *CreateQueryAllCoinService) Do(ctx context.Context, opts ...RequestOption) (res []*CoinData, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*CoinData{}, err
	}
	res = make([]*CoinData, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*CoinData{}, err
	}
	return res, nil
}
