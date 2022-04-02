// Copyright 2022 Hash Protocol
// SPDX-License-Identifier: LGPL-3.0-only

package chains

import (
	"github.com/hashprotocol/hashbridge-utils/msg"
)

type Router interface {
	Send(message msg.Message) error
}

//type Writer interface {
//	ResolveMessage(message msg.Message) bool
//}
