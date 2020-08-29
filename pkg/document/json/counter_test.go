/*
 * Copyright 2020 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package json_test

import (
	"testing"
	time2 "time"

	"github.com/stretchr/testify/assert"

	"github.com/yorkie-team/yorkie/pkg/document/json"
	"github.com/yorkie-team/yorkie/pkg/document/time"
)

func TestCounter(t *testing.T) {
	t.Run("increase test", func(t *testing.T) {
		var x int = 5
		var y int64 = 10
		var z float64 = 3.14
		integer := json.NewCounter(x, time.InitialTicket)
		long := json.NewCounter(y, time.InitialTicket)
		double := json.NewCounter(z, time.InitialTicket)

		integerOperand := integer.DeepCopy().(*json.Counter)
		longOperand := long.DeepCopy().(*json.Counter)
		doubleOperand := double.DeepCopy().(*json.Counter)

		// normal process test
		integer.Increase(integerOperand)
		integer.Increase(longOperand)
		integer.Increase(doubleOperand)
		assert.Equal(t, integer.Marshal(), "23")

		long.Increase(integerOperand)
		long.Increase(longOperand)
		long.Increase(doubleOperand)
		assert.Equal(t, long.Marshal(), "28")

		double.Increase(integerOperand)
		double.Increase(longOperand)
		double.Increase(doubleOperand)
		assert.Equal(t, double.Marshal(), "21.280000")

		// error process test
		// TODO it should be modified to error check
		// when 'Remove panic from server code (#50)' is completed.
		unsupportedTypePanicTest := func() {
			r := recover()
			assert.NotNil(t, r)
			assert.Equal(t, r, "unsupported type")
		}
		unsupportedTest := func(v interface{}) {
			defer unsupportedTypePanicTest()
			json.NewCounter(v, time.InitialTicket)
		}
		unsupportedTest("str")
		unsupportedTest(true)
		unsupportedTest([]byte{2})
		unsupportedTest(time2.Now())

		assert.Equal(t, integer.Marshal(), "23")
		assert.Equal(t, long.Marshal(), "28")
		assert.Equal(t, double.Marshal(), "21.280000")
	})
}
