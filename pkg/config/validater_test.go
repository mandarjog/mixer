// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	aspectpb "istio.io/api/mixer/v1/config/aspect"
	"istio.io/mixer/pkg/aspect"
	"istio.io/mixer/pkg/attribute"
)

type fakeVFinder struct {
	v map[string]aspect.ConfigValidater
}

func (f *fakeVFinder) Find(name string) (aspect.ConfigValidater, bool) {
	v, found := f.v[name]
	return v, found
}

type lc struct {
	ce *aspect.ConfigErrors
}

func (m *lc) DefaultConfig() (implConfig proto.Message) {
	return &aspectpb.ListCheckerConfig{}
}

// ValidateConfig determines whether the given configuration meets all correctness requirements.
func (m *lc) ValidateConfig(implConfig proto.Message) *aspect.ConfigErrors {
	return m.ce
}

type configTable struct {
	cerr     *aspect.ConfigErrors
	v        map[string]aspect.ConfigValidater
	nerrors  int
	selector string
	strict   bool
	cfg      string
}

func TestConfigValidatorError(t *testing.T) {
	var ct *aspect.ConfigErrors
	evaluator := newFakeExpr()
	cerr := ct.Appendf("Url", "Must have a valid URL")

	ctable := []*configTable{
		{nil,
			map[string]aspect.ConfigValidater{
				"metrics":  &lc{},
				"metrics2": &lc{},
			}, 0, "service.name == “*”", false, sSvcConfig},
		{nil,
			map[string]aspect.ConfigValidater{
				"istio/denychecker": &lc{},
				"metrics2":          &lc{},
			}, 0, "service.name == “*”", false, sGlobalConfig},
		{nil,
			map[string]aspect.ConfigValidater{
				"metrics":  &lc{},
				"metrics2": &lc{},
			}, 1, "service.name == “*”", false, sGlobalConfig},
		{nil,
			map[string]aspect.ConfigValidater{
				"metrics":  &lc{},
				"metrics2": &lc{},
			}, 1, "service.name == “*”", true, sSvcConfig},
		{cerr,
			map[string]aspect.ConfigValidater{
				"metrics": &lc{ce: cerr},
			}, 2, "service.name == “*”", false, sSvcConfig},
		{ct.Append("/:metrics", UnknownValidater("metrics")),
			nil, 3, "\"\"", false, sSvcConfig},
	}

	for idx, ctx := range ctable {
		var ce *aspect.ConfigErrors
		mgr := &fakeVFinder{v: ctx.v}
		p := NewValidater(mgr, mgr, ctx.strict, evaluator)
		if ctx.cfg == sSvcConfig {
			ce = p.validateServiceConfig(fmt.Sprintf(ctx.cfg, ctx.selector), false)
		} else {
			ce = p.validateGlobalConfig(ctx.cfg)
		}
		cok := ce == nil
		ok := ctx.nerrors == 0

		if ok != cok {
			t.Errorf("%d Expected %t Got %t", idx, ok, cok)
		}
		if ce == nil {
			continue
		}

		if len(ce.Multi.Errors) != ctx.nerrors {
			t.Error(idx, "\nExpected:", ctx.cerr.Error(), "\nGot:", ce.Error(), len(ce.Multi.Errors), ctx.nerrors)
		}
	}
}

func TestFullConfigValidator(t *testing.T) {
	//var ccfg []*Combined
	//bag, err := attribute.NewManager().NewTracker().StartRequest(&mixerpb.Attributes{})
	//if err != nil {
	//	t.Error("Unable to get attribute bag")
	//}
	fe := newFakeExpr()

	ctable := []*configTable{
		{nil,
			map[string]aspect.ConfigValidater{
				"istio/denychecker": &lc{},
				"metrics":           &lc{},
				"listchecker":       &lc{},
			}, 1, "service.name == “*”", false, sSvcConfig1},
		{nil,
			map[string]aspect.ConfigValidater{
				"istio/denychecker": &lc{},
				"metrics":           &lc{},
				"listchecker":       &lc{},
			}, 0, "service.name == “*”", false, sSvcConfig2},
	}
	for idx, ctx := range ctable {
		mgr := &fakeVFinder{v: ctx.v}
		p := NewValidater(mgr, mgr, ctx.strict, fe)

		_, ce := p.Validate(ctx.cfg, sGlobalConfig)
		cok := ce == nil
		ok := ctx.nerrors == 0

		if ok != cok {
			t.Errorf("%d Expected %t Got %t", idx, ok, cok)
		}

		if ce != nil && len(ce.Multi.Errors) != ctx.nerrors {
			t.Error(idx, "\nExpected:", ctx.cerr.Error(), "\nGot:", ce.Error(), len(ce.Multi.Errors), ctx.nerrors)
		}

	}
}

func TestConfigParseError(t *testing.T) {
	mgr := &fakeVFinder{}
	evaluator := newFakeExpr()
	p := NewValidater(mgr, mgr, false, evaluator)
	ce := p.validateServiceConfig("<config>  </config>", false)

	if ce == nil || !strings.Contains(ce.Error(), "unmarshal error") {
		t.Error("Expected unmarhal Error", ce)
	}

	ce = p.validateGlobalConfig("<config>  </config>")

	if ce == nil || !strings.Contains(ce.Error(), "unmarshal error") {
		t.Error("Expected unmarhal Error", ce)
	}
}

const sGlobalConfig = `
subject: "namespace:ns"
revision: "2022"
adapters:
  - name: denychecker.1
    kind: listchecker
    impl: istio/denychecker
    params:
      checkattribute: src.ip
      blacklist: true
      unknown_field: true
`

const sSvcConfig1 = `
subject: "namespace:ns"
revision: "2022"
rules:
- selector: service.name == “*”
  aspects:
  - kind: metrics
    params:
      metrics:
      - name: response_time_by_consumer
        value: metric_response_time
        metric_kind: DELTA
        labels:
        - key: target_consumer_id
`
const sSvcConfig2 = `
subject: namespace:ns
revision: "2022"
rules:
- selector: service.name == “*”
  aspects:
  - kind: listchecker
    inputs: {}
    params:
`

const sSvcConfig = `
subject: namespace:ns
revision: "2022"
rules:
- selector: %s
  aspects:
  - kind: metrics
    adapter: ""
    inputs: {}
    params:
      checkattribute: src.ip
      blacklist: true
      unknown_field: true
  rules:
  - selector: src.name == "abc"
    aspects:
    - kind: metrics2
      adapter: ""
      inputs: {}
      params:
        checkattribute: src.ip
        blacklist: true
`

type fakeExpr struct {
}

// newFakeExpr returns the basic
func newFakeExpr() *fakeExpr {
	return &fakeExpr{}
}

func UnboundVariable(vname string) error {
	return fmt.Errorf("unbound variable %s", vname)
}

// Eval evaluates given expression using the attribute bag
func (e *fakeExpr) Eval(mapExpression string, attrs attribute.Bag) (v interface{}, err error) {
	var found bool

	v, found = attrs.String(mapExpression)
	if found {
		return
	}

	v, found = attrs.Bool(mapExpression)
	if found {
		return
	}

	v, found = attrs.Int64(mapExpression)
	if found {
		return
	}

	v, found = attrs.Float64(mapExpression)
	if found {
		return
	}

	return v, UnboundVariable(mapExpression)
}

// EvalString evaluates given expression using the attribute bag to a string
func (e *fakeExpr) EvalString(mapExpression string, attrs attribute.Bag) (v string, err error) {
	var found bool
	v, found = attrs.String(mapExpression)
	if found {
		return
	}
	return v, UnboundVariable(mapExpression)
}

// EvalPredicate evaluates given predicate using the attribute bag
func (e *fakeExpr) EvalPredicate(mapExpression string, attrs attribute.Bag) (v bool, err error) {
	var found bool
	v, found = attrs.Bool(mapExpression)
	if found {
		return
	}
	return v, UnboundVariable(mapExpression)
}

func (e *fakeExpr) Validate(expression string) error { return nil }
