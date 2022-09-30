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

package startuphandler

import (
	"context"
	"sync"

	api "github.com/gravitational/teleport/lib/teleterm/api/protogen/golang/v1"

	"github.com/gravitational/trace"
)

// Handler holds values needed to orchestrate the procedure of setting up the tshd events service
// client on app startup.
type Handler struct {
	// mu is used to ensure we don't somehow end up with two concurrent calls to
	// ResolveTshdEventsServerAddress.
	mu sync.Mutex
	// WaitForTshdEventsServerAddressC gets closed after the address becomes available, that is the
	// Electron app calls ResolveTshdEventsServerAddress.
	WaitForTshdEventsServerAddressC chan struct{}
	// TshdEventsServerAddress becomes available after the Electron app makes a call to
	// ResolveTshdEventsServerAddress.
	TshdEventsServerAddress string
	// waitForTshdEventsClientC is closed after the APIServer creates a tshd events client and injects
	// it into daemon.Service.
	waitForTshdEventsClientC chan struct{}
}

func New() (*Handler, error) {
	return &Handler{
		WaitForTshdEventsServerAddressC: make(chan struct{}),
		waitForTshdEventsClientC:        make(chan struct{}),
	}, nil
}

// RPC handlers

// ResolveTshdEventsServerAddress is called by the Electron app after the tshd events server starts.
// It'll return an error if called more than once within the application lifetime – there's no need
// to do so, if it's called more than once then it's a sign the Electron app is buggy.
func (h *Handler) ResolveTshdEventsServerAddress(ctx context.Context, req *api.ResolveTshdEventsServerAddressRequest) (*api.ResolveTshdEventsServerAddressResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	select {
	case <-h.WaitForTshdEventsServerAddressC:
		// The channel is closed so the address must have been resolved already.
		return nil, trace.AlreadyExists("repeated call to resolve tshd events server address")
	default:
	}

	h.TshdEventsServerAddress = req.Address
	close(h.WaitForTshdEventsServerAddressC)
	return &api.ResolveTshdEventsServerAddressResponse{}, nil
}

// WaitForTshdEventsClient is called by the Electron app. The function returns after the APIServer
// creates a tshd events client and injects it into daemon.Service.
//
// The Electron app waits with performing any other calls until this call finishes. This makes sure
// that we don't use daemon.Service while the tshd events client is not ready yet.
func (h *Handler) WaitForTshdEventsClient(ctx context.Context, req *api.WaitForTshdEventsClientRequest) (*api.WaitForTshdEventsClientResponse, error) {
	select {
	case <-ctx.Done():
		return nil, trace.Wrap(ctx.Err())
	case <-h.waitForTshdEventsClientC:
		return &api.WaitForTshdEventsClientResponse{}, nil
	}
}

// Other methods

// MarkTshdEventsClientAsReady closes waitForTshdEventsClientC, allowing the WaitForTshdEventsClient
// handler to return a response.
//
// In a normal operation, this function won't get called more than once but we added some
// precautions just to be sure.
func (h *Handler) MarkTshdEventsClientAsReady() {
	h.mu.Lock()
	defer h.mu.Unlock()

	select {
	case <-h.waitForTshdEventsClientC:
		// The channel was already closed.
		return
	default:
	}

	close(h.waitForTshdEventsClientC)
}
