// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package imaging

import (
	"context"
	"github.com/juju/errors"
	"github.com/ErfannNikooee/onvif"
	"github.com/ErfannNikooee/onvif/sdk"
	"github.com/ErfannNikooee/onvif/Imaging"
)

// Call_RelativeMove forwards the call to dev.CallMethod() then parses the payload of the reply as a MoveResponse.
func Call_RelativeMove(ctx context.Context, dev *onvif.Device, request imaging.RelativeMove) (imaging.MoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			MoveResponse imaging.MoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.MoveResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Move")
		return reply.Body.MoveResponse, errors.Annotate(err, "reply")
	}
}
