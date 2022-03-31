// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package tss

import (
	"crypto/elliptic"
	"errors"
	"time"

	"github.com/binance-chain/tss-lib/common"
)

type (
	Parameters struct {
		ec                      elliptic.Curve
		partyID                 *PartyID
		parties                 *PeerContext
		partyCount              uint
		threshold               uint
		safePrimeGenTimeout     time.Duration
		unsafeKGIgnoreH1H2Dupes bool
	}

	ReSharingParameters struct {
		*Parameters
		newParties    *PeerContext
		newPartyCount uint
		newThreshold  uint
	}
)

const (
	defaultSafePrimeGenTimeout = 5 * time.Minute
)

// Exported, used in `tss` client
func NewParameters(ec elliptic.Curve, ctx *PeerContext, partyID *PartyID, partyCount, threshold uint, optionalSafePrimeGenTimeout ...time.Duration) *Parameters {
	var safePrimeGenTimeout time.Duration
	if threshold >= partyCount {
		panic(errors.New("NewParameters: t<n necessarily with the dishonest majority assumption"))
	}
	if 0 < len(optionalSafePrimeGenTimeout) {
		if 1 < len(optionalSafePrimeGenTimeout) {
			panic(errors.New("GeneratePreParams: expected 0 or 1 item in `optionalSafePrimeGenTimeout`"))
		}
		safePrimeGenTimeout = optionalSafePrimeGenTimeout[0]
	} else {
		safePrimeGenTimeout = defaultSafePrimeGenTimeout
	}
	return &Parameters{
		ec:                  ec,
		parties:             ctx,
		partyID:             partyID,
		partyCount:          partyCount,
		threshold:           threshold,
		safePrimeGenTimeout: safePrimeGenTimeout,
	}
}

func (params *Parameters) EC() elliptic.Curve {
	return params.ec
}

func (params *Parameters) Parties() *PeerContext {
	return params.parties
}

func (params *Parameters) PartyID() *PartyID {
	return params.partyID
}

func (params *Parameters) PartyCount() uint {
	return params.partyCount
}

func (params *Parameters) Threshold() uint {
	return params.threshold
}

func (params *Parameters) SafePrimeGenTimeout() time.Duration {
	return params.safePrimeGenTimeout
}

// Getter. The H1, H2 dupe check is disabled during some benchmarking scenarios to allow reuse of pre-params.
func (params *Parameters) UNSAFE_KGIgnoreH1H2Dupes() bool {
	return params.unsafeKGIgnoreH1H2Dupes
}

// Setter. The H1, H2 dupe check is disabled during some benchmarking scenarios to allow reuse of pre-params.
func (params *Parameters) UNSAFE_setKGIgnoreH1H2Dupes(unsafeKGIgnoreH1H2Dupes bool) {
	if unsafeKGIgnoreH1H2Dupes {
		common.Logger.Warn("UNSAFE_setKGIgnoreH1H2Dupes() has been called; do not use these shares in production.")
	}
	params.unsafeKGIgnoreH1H2Dupes = unsafeKGIgnoreH1H2Dupes
}

// ----- //

// Exported, used in `tss` client
func NewReSharingParameters(ec elliptic.Curve, ctx, newCtx *PeerContext, partyID *PartyID, partyCount, threshold, newPartyCount, newThreshold uint) *ReSharingParameters {
	params := NewParameters(ec, ctx, partyID, partyCount, threshold)
	return &ReSharingParameters{
		Parameters:    params,
		newParties:    newCtx,
		newPartyCount: newPartyCount,
		newThreshold:  newThreshold,
	}
}

func (rgParams *ReSharingParameters) OldParties() *PeerContext {
	return rgParams.Parties() // wr use the original method for old parties
}

func (rgParams *ReSharingParameters) OldPartyCount() uint {
	return rgParams.partyCount
}

func (rgParams *ReSharingParameters) NewParties() *PeerContext {
	return rgParams.newParties
}

func (rgParams *ReSharingParameters) NewPartyCount() uint {
	return rgParams.newPartyCount
}

func (rgParams *ReSharingParameters) NewThreshold() uint {
	return rgParams.newThreshold
}

func (rgParams *ReSharingParameters) OldAndNewParties() []*PartyID {
	return append(rgParams.OldParties().IDs(), rgParams.NewParties().IDs()...)
}

func (rgParams *ReSharingParameters) OldAndNewPartyCount() uint {
	return rgParams.OldPartyCount() + rgParams.NewPartyCount()
}

func (rgParams *ReSharingParameters) IsOldCommittee() bool {
	partyID := rgParams.partyID
	for _, Pj := range rgParams.parties.IDs() {
		if partyID.KeyInt().Cmp(Pj.KeyInt()) == 0 {
			return true
		}
	}
	return false
}

func (rgParams *ReSharingParameters) IsNewCommittee() bool {
	partyID := rgParams.partyID
	for _, Pj := range rgParams.newParties.IDs() {
		if partyID.KeyInt().Cmp(Pj.KeyInt()) == 0 {
			return true
		}
	}
	return false
}
