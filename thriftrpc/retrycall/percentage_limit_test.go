// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package retrycall

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex-tests/thriftrpc"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/transport"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPercentageLimit(t *testing.T) {
	testPercentageLimit(t, true)
	testPercentageLimit(t, false)
}

func testPercentageLimit(t *testing.T, enable bool) {
	backupPolicy := &retry.BackupPolicy{
		RetryDelayMS: 100,
		StopPolicy: retry.StopPolicy{
			MaxRetryTimes:    2,
			DisableChainStop: false,
			CBPolicy: retry.CBPolicy{
				ErrorRate: 0.1,
			},
		},
	}

	var rc *retry.Container
	if enable {
		rc = retry.NewRetryContainerWithPercentageLimit()
	} else {
		rc = retry.NewRetryContainer()
	}
	rc.NotifyPolicyChange("circuitBreakTest", retry.BuildBackupRequest(backupPolicy))

	totalCnt, retryCnt := int32(0), int32(0)
	cli := getKitexClient(
		transport.TTHeader,
		client.WithRetryContainer(rc),
		client.WithHostPorts(":9002"),
		client.WithMiddleware(func(next endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, req, resp interface{}) (err error) {
				atomic.AddInt32(&totalCnt, 1)
				if retry.IsLocalRetryRequest(ctx) {
					atomic.AddInt32(&retryCnt, 1)
				}
				return next(ctx, req, resp)
			}
		}),
	)
	ctx, stReq := thriftrpc.CreateSTRequest(context.Background())
	ctx = metainfo.WithPersistentValue(ctx, sleepTimeMsKey, "200") //ms

	n := 200
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			_, _ = cli.CircuitBreakTest(ctx, stReq, callopt.WithRPCTimeout(400*time.Millisecond))
		}()
	}
	wg.Wait()
	t.Logf("enable = %v, total = %v, retry = %v", enable, atomic.LoadInt32(&totalCnt), atomic.LoadInt32(&retryCnt))
}
