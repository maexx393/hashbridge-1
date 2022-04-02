// Copyright 2022 Hash Protocol
// SPDX-License-Identifier: LGPL-3.0-only

package shared

import (
	log "github.com/hashprotocol/log15"
)

func SetLogger(lvl log.Lvl) {
	logger := log.Root()
	handler := logger.GetHandler()
	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))
}
