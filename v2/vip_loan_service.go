package binance

import (
	"context"
	"encoding/json"
	// "github.com/taotao5/go-binance/v2/common"
)

// CreateWithdrawService submits a withdraw request.
//
// See https://binance-docs.github.io/apidocs/spot/en/#withdraw
type ListVipLoanOrdersService struct {
	c         *Client
	timesTamp int64
}

// Asset sets the asset parameter (MANDATORY).
func (s *ListVipLoanOrdersService) TimesTamp(v int64) *ListVipLoanOrdersService {
	s.timesTamp = v
	return s
}

// Do sends the request.
func (s *ListVipLoanOrdersService) Do(ctx context.Context) (*ListVipLoanOrdersResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/loan/vip/ongoing/orders",
		secType:  secTypeSigned,
	}
	r.setParam("timestamp", s.timesTamp)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &ListVipLoanOrdersResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 定义 JSON 的结构体
type Row struct {
	OrderID                          string `json:"orderId"`
	LoanCoin                         string `json:"loanCoin"`
	TotalDebt                        string `json:"totalDebt"`
	LoanRate                         string `json:"loanRate"`
	ResidualInterest                 string `json:"residualInterest"`
	CollateralAccountID              string `json:"collateralAccountId"`
	CollateralCoin                   string `json:"collateralCoin"`
	TotalCollateralValueAfterHaircut string `json:"totalCollateralValueAfterHaircut"`
	LockedCollateralValue            string `json:"lockedCollateralValue"`
	CurrentLTV                       string `json:"currentLTV"`
	ExpirationTime                   string `json:"expirationTime"`
	LoanDate                         string `json:"loanDate"`
	LoanTerm                         string `json:"loanTerm"`
}

type ListVipLoanOrdersResponse struct {
	Rows  []Row `json:"rows"`
	Total int   `json:"total"`
}

type CreateVipLoanRepayService struct {
	c       *Client
	orderId int64
	amount  string
}

func (s *CreateVipLoanRepayService) OrderId(v int64) *CreateVipLoanRepayService {
	s.orderId = v
	return s
}

func (s *CreateVipLoanRepayService) Amount(v string) *CreateVipLoanRepayService {
	s.amount = v
	return s
}

// Do sends the request.
func (s *CreateVipLoanRepayService) Do(ctx context.Context) (*CreateVipLoanRepayResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/loan/vip/repay",
		secType:  secTypeSigned,
	}

	m := params{
		"orderId": s.orderId,
		"amount":  s.amount,
	}

	r.setFormParams(m)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateVipLoanRepayResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type CreateVipLoanRepayResponse struct {
	LoanCoin           string `json:"loanCoin"`
	RepayAmount        string `json:"repayAmount"`
	RemainingPrincipal string `json:"remainingPrincipal"`
	RemainingInterest  string `json:"remainingInterest"`
	CollateralCoin     string `json:"collateralCoin"`
	CurrentLTV         string `json:"currentLTV"`
	RepayStatus        string `json:"repayStatus"`
}

type QueryVipLoanInterestRateService struct {
	c         *Client
	timesTamp int64
	loanCoin  string
}

// Asset sets the asset parameter (MANDATORY).
func (s *QueryVipLoanInterestRateService) TimesTamp(v int64) *QueryVipLoanInterestRateService {
	s.timesTamp = v
	return s
}

func (s *QueryVipLoanInterestRateService) LoanCoin(v string) *QueryVipLoanInterestRateService {
	s.loanCoin = v
	return s
}

// Do sends the request.
func (s *QueryVipLoanInterestRateService) Do(ctx context.Context) ([]*AssetInterestRate, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/loan/vip/request/interestRate",
		secType:  secTypeSigned,
	}

	r.setParam("timestamp", s.timesTamp)

	r.setParam("loanCoin", s.loanCoin)

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := make([]*AssetInterestRate, 0)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// 定义 JSON 的结构体
type AssetInterestRate struct {
	Asset                      string `json:"asset"`
	FlexibleDailyInterestRate  string `json:"flexibleDailyInterestRate"`
	FlexibleYearlyInterestRate string `json:"flexibleYearlyInterestRate"`
	Time                       string `json:"time"`
}
