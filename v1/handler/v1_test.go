package handler

import (
	"context"
	"testing"

	balance "github.com/m3o/services/balance/proto"
	balancefake "github.com/m3o/services/balance/proto/fakes"
	publicapi "github.com/m3o/services/publicapi/proto"
	publicapifake "github.com/m3o/services/publicapi/proto/fakes"
	usage "github.com/m3o/services/usage/proto"
	usagefake "github.com/m3o/services/usage/proto/fakes"
	muclient "github.com/micro/micro/v3/service/client"
	. "github.com/onsi/gomega"
)

func TestCallVerification(t *testing.T) {
	g := NewWithT(t)

	tcs := []struct {
		key            *apiKeyRecord
		name           string
		err            error
		priceRet       string
		quota          int64
		price          int64
		reqCount       int64
		totalFreeCount int64
	}{
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"helloworld"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:     "free",
			err:      nil,
			priceRet: "free",
			quota:    0,
			reqCount: 0,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"helloworld"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:           "free monthly cap exceeded",
			err:            errBlocked("Monthly usage cap exceeded"),
			priceRet:       "",
			totalFreeCount: 1000000,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"helloworld"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:     "free quota",
			err:      nil,
			priceRet: "0",
			quota:    1,
			price:    1,
			reqCount: 0,
		},
		{
			key: &apiKeyRecord{
				ID:            "id",
				ApiKey:        "apikey",
				Scopes:        []string{"helloworld"},
				UserID:        "userid",
				AccID:         "accid",
				Namespace:     "micro",
				Token:         "token",
				Status:        keyStatusBlocked,
				StatusMessage: "foobar",
			},
			name:     "free quota blocked key",
			err:      errBlocked("foobar"),
			priceRet: "",
			quota:    1,
			price:    1,
			reqCount: 0,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"helloworld"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:           "free quota monthly cap exceeded but paid for api quota",
			err:            nil,
			priceRet:       "0",
			quota:          1,
			price:          1,
			reqCount:       0,
			totalFreeCount: 1000000,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"helloworld"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:           "free quota monthly cap exceeded but paid for api",
			err:            nil,
			priceRet:       "1",
			price:          1,
			reqCount:       0,
			totalFreeCount: 1000000,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"foo"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:           "block no scope",
			err:            errBlocked("Insufficient privileges"),
			priceRet:       "",
			totalFreeCount: 1000000,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"*"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:     "Wilcard scope",
			err:      nil,
			priceRet: "free",
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"*"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:     "Insufficient funds",
			err:      errBlocked("Insufficient funds"),
			priceRet: "",
			price:    10000,
		},
		{
			key: &apiKeyRecord{
				ID:        "id",
				ApiKey:    "apikey",
				Scopes:    []string{"*"},
				UserID:    "userid",
				AccID:     "accid",
				Namespace: "micro",
				Token:     "token",
				Status:    keyStatusActive,
			},
			name:     "Insufficient funds but still in quota",
			err:      nil,
			priceRet: "0",
			price:    10000,
			quota:    10,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fakepapi := publicapifake.FakePublicapiService{
				ListStub: func(ctx context.Context, request *publicapi.ListRequest, option ...muclient.CallOption) (*publicapi.ListResponse, error) {
					return &publicapi.ListResponse{Apis: []*publicapi.PublicAPI{
						{
							Id:      "helloworld",
							Name:    "helloworld",
							Pricing: map[string]int64{"Helloworld.Call": tc.price},
							Quotas:  map[string]int64{"Helloworld.Call": tc.quota},
						},
					}}, nil
				},
			}
			fakebalance := balancefake.FakeBalanceService{
				CurrentStub: func(ctx context.Context, request *balance.CurrentRequest, option ...muclient.CallOption) (*balance.CurrentResponse, error) {
					return &balance.CurrentResponse{CurrentBalance: 10}, nil
				},
			}

			fakeusage := usagefake.FakeUsageService{
				ReadMonthlyTotalStub: func(ctx context.Context, request *usage.ReadMonthlyTotalRequest, option ...muclient.CallOption) (*usage.ReadMonthlyTotalResponse, error) {
					return &usage.ReadMonthlyTotalResponse{
						Requests: 2,
						EndpointRequests: map[string]int64{
							"helloworld$Helloworld.Call": tc.reqCount,
							"total":                      125,
							"totalfree":                  tc.totalFreeCount,
						},
					}, nil
				},
			}

			v1api := &V1{
				publicapiCache: &publicapiCache{
					papi: &fakepapi,
				},
				balanceCache: &balanceCache{
					balsvc: &fakebalance,
				},
				usageCache: &usageCache{
					usagesvc: &fakeusage,
				},
			}
			v1api.publicapiCache.init()

			pr, err := v1api.verifyCallAllowed(context.Background(), tc.key, "/v1/helloworld/Call")
			g.Expect(pr).To(Equal(tc.priceRet))
			if tc.err == nil {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(Equal(tc.err))
			}

		})
	}

}
