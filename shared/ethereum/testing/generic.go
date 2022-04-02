// Copyright 2022 Hash Protocol
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hashprotocol/hashbridge-utils/msg"
	utils "github.com/hashprotocol/hashbridge/shared/ethereum"
	"github.com/hashprotocol/log15"
)

func AssertGenericResourceAddress(t *testing.T, client *utils.Client, handler common.Address, rId msg.ResourceId, expected common.Address) {
	actual, err := utils.GetGenericResourceAddress(client, handler, rId)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual.Bytes(), expected.Bytes()) {
		t.Fatalf("Generic resoruce mismatch for ID %x. Expected address: %x Got: %x", rId, expected, actual)
	}
	log15.Info("Asserted generic resource ID", "handler", handler, "rId", rId.Hex(), "contract", actual)
}
