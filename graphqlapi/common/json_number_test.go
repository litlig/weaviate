/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */package common

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input          interface{}
	expectedOutput float64
}

func TestJSONNumberResolver(t *testing.T) {
	tests := []testCase{
		testCase{
			input:          json.Number("10"),
			expectedOutput: 10.0,
		},
		testCase{
			input:          int(10),
			expectedOutput: 10.0,
		},
		testCase{
			input:          int64(10),
			expectedOutput: 10.0,
		},
		testCase{
			input:          float64(10),
			expectedOutput: 10.0,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%#v -> %#v", test.input, test.expectedOutput)
		t.Run(name, func(t *testing.T) {
			result, err := JSONNumberResolver(resolveParams(test.input))
			assert.Nil(t, err, "should not error")
			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

func resolveParams(input interface{}) graphql.ResolveParams {
	return graphql.ResolveParams{
		Source: map[string]interface{}{
			"myField": input,
		},
		Info: graphql.ResolveInfo{FieldName: "myField"},
	}
}
