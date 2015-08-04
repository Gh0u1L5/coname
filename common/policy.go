// Copyright 2014-2015 The Dename Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package common

import (
	"fmt"

	"github.com/yahoo/coname/proto"
)

// VerifyUpdate returns nil iff replacing entry current (nil if none) with next
// is justified given the evidence in update. Globally deterministic.
func VerifyUpdate(current *proto.Entry, update *proto.SignedEntryUpdate) error {
	next := &update.NewEntry
	if current != nil {
		if current.UpdatePolicy == nil {
			return fmt.Errorf("VerifyUpdate: current.UpdatePolicy is nil")
		}
		if !VerifyPolicy(current.UpdatePolicy, update.NewEntry.PreservedEncoding, update.Signatures) {
			return fmt.Errorf("VerifyUpdate: replacing an entry requires authorization from the old key, but signature verification failed")
		}
		if next.Version < current.Version {
			return fmt.Errorf("VerifyUpdate: entry version must not decrease (got %d < %d)", next.Version, current.Version)
		}
	} else if next.Version != 0 {
		return fmt.Errorf("VerifyUpdate: registering a new entry must use version number 0 (got %d)", next.Version)
	}
	if next.UpdatePolicy == nil {
		return fmt.Errorf("VerifyUpdate: next.UpdatePolicy is nil")
	}
	if !VerifyPolicy(next.UpdatePolicy, update.NewEntry.PreservedEncoding, update.Signatures) {
		return fmt.Errorf("VerifyUpdate: update needs to be accepted by the new key, but signature verification failed")
	}
	return nil
}

// VerifyPolicy returns whether, by policy, action is justified by evidence.
// Evidence is in the form of digital signatures denoting agreement, and the
// policy contains public keys and a quorum rule.
func VerifyPolicy(policy *proto.AuthorizationPolicy, action []byte, evidence map[uint64][]byte) bool {
	have := make(map[uint64]struct{}, len(evidence))
	for id, pk := range policy.PublicKeys {
		if sig, ok := evidence[id]; ok && VerifySignature(pk, action, sig) {
			have[id] = struct{}{}
		}
	}
	switch {
	case policy.Quorum != nil:
		return CheckQuorum(policy.Quorum, have)
	default: // unknown policy
		return false
	}
}

// CheckQuorum evaluates whether the quorum requirement want can be satisfied
// by ratifications of the verifiers in have.
func CheckQuorum(want *proto.QuorumExpr, have map[uint64]struct{}) bool {
	if want == nil {
		return true // no requirements
	}
	var n uint32
	for _, verifier := range want.Candidates {
		if _, yes := have[verifier]; yes {
			n++
		}
	}
	for _, e := range want.Subexpressions {
		if CheckQuorum(e, have) {
			n++
		}
	}
	return n >= want.Threshold
}

// ListQuorum inserts all verifiers mentioned in e to out. If out is nil, a new
// map is allocated. ListQuorum is NOT intended to be used for implementing
// quorum verification, use CheckQuorum instead.
func ListQuorum(e *proto.QuorumExpr, out map[uint64]struct{}) map[uint64]struct{} {
	if e == nil {
		return make(map[uint64]struct{}, 0)
	}
	if out == nil {
		var l int
		if e.Candidates != nil {
			l = len(e.Candidates)
		}
		out = make(map[uint64]struct{}, l)
	}
	for _, verifier := range e.Candidates {
		out[verifier] = struct{}{}
	}
	for _, e := range e.Subexpressions {
		ListQuorum(e, out)
	}
	return out
}
