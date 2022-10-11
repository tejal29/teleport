// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dynamo

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	apiRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dynamo_requests_total",
			Help: "Requests to the DynamoDB API",
		},
		[]string{"type", "operation"},
	)
	apiRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "dynamo_requests",
			Help: "Requests to the DynamoDB API by result",
		},
		[]string{"type", "operation", "result", "throughput_exceeded"},
	)
	apiRequestLatencies = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "dynamo_requests_seconds",
			Help: "Request latency for the DynamoDB API",
			// lowest bucket start of upper bound 0.001 sec (1 ms) with factor 2
			// highest bucket start of 0.001 sec * 2^15 == 32.768 sec
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 16),
		},
		[]string{"type", "operation"},
	)
	apiRequestCapacityConsumed = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "dynamo_requests_capacity_consumed",
			Help: "Request capacity consumed for the DynamoDB API",
			// lowest bucket start of upper bound 0.001 sec (1 ms) with factor 2
			// highest bucket start of 0.001 sec * 2^15 == 32.768 sec
			// TODO: fix this for expected consumed r/w units
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 16),
		},
		[]string{"type", "operation", "target", "rw"},
	)

	dynamoCollectors = []prometheus.Collector{
		apiRequests,
		apiRequestsTotal,
		apiRequestLatencies,
	}
)

// TableType indicates which type of table metrics are being calculated for
type TableType string

const (
	// Backend is a table used to store backend data.
	Backend TableType = "backend"
	// Events is a table used to store audit events.
	Events TableType = "events"
)

// recordMetrics updates the set of dynamo api metrics
func recordMetrics(tableType TableType, operation string, err error, latency float64, consumedCapacities ...*dynamodb.ConsumedCapacity) {
	labels := []string{string(tableType), operation}
	apiRequestsTotal.WithLabelValues(labels...).Inc()
	apiRequestLatencies.WithLabelValues(labels...).Observe(latency)

	result := "success"
	if err != nil {
		result = "error"
	}
	exceeded := strconv.FormatBool(provisionedThroughputExceeded(err))
	apiRequests.WithLabelValues(append(labels, result, exceeded)...).Inc()

	// record capacity consumed
	tableReadCapacityUnits := 0.0
	tableWriteCapacityUnits := 0.0
	indexReadCapacityUnits := 0.0
	indexWriteCapacityUnits := 0.0
	for _, consumedCapacity := range consumedCapacities {
		// compute capacity r/w units consumed by table
		tableReadCapacityUnits += *consumedCapacity.Table.ReadCapacityUnits
		tableWriteCapacityUnits += *consumedCapacity.Table.WriteCapacityUnits

		// compute capacity r/w units consumed by indexes
		for _, indexConsumedCapacity := range consumedCapacity.LocalSecondaryIndexes {
			indexReadCapacityUnits += *indexConsumedCapacity.ReadCapacityUnits
			indexWriteCapacityUnits += *indexConsumedCapacity.WriteCapacityUnits
		}
		for _, indexConsumedCapacity := range consumedCapacity.GlobalSecondaryIndexes {
			indexReadCapacityUnits += *indexConsumedCapacity.ReadCapacityUnits
			indexWriteCapacityUnits += *indexConsumedCapacity.WriteCapacityUnits
		}
	}
	recordConsumedCapacity := func(target, kind string, value float64) {
		if value > 0 {
			apiRequestCapacityConsumed.WithLabelValues(append(labels, target, kind)...).Observe(value)
		}
	}
	recordConsumedCapacity("table", "read", tableReadCapacityUnits)
	recordConsumedCapacity("table", "write", tableWriteCapacityUnits)
	recordConsumedCapacity("index", "read", indexReadCapacityUnits)
	recordConsumedCapacity("index", "write", indexWriteCapacityUnits)
}

func provisionedThroughputExceeded(err error) bool {
	if err == nil {
		return false
	}
	aerr, ok := err.(awserr.Error)
	if !ok {
		return false
	}
	return aerr.Code() == dynamodb.ErrCodeProvisionedThroughputExceededException
}
