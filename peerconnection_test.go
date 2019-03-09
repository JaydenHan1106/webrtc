package webrtc

import (
	"reflect"
	"testing"

	"github.com/pions/webrtc/pkg/rtcerr"
	"github.com/stretchr/testify/assert"
)

func TestPeerConnection_SetConfiguration(t *testing.T) {
	// Note: These tests don't include ICEServer.Credential,
	// ICEServer.CredentialType, or Certificates because those are not supported
	// in the WASM bindings.

	for _, test := range []struct {
		name    string
		init    func() (*PeerConnection, error)
		config  Configuration
		wantErr error
	}{
		{
			name: "valid",
			init: func() (*PeerConnection, error) {
				pc, err := NewPeerConnection(Configuration{
					PeerIdentity:         "unittest",
					ICECandidatePoolSize: 5,
				})
				if err != nil {
					return pc, err
				}

				err = pc.SetConfiguration(Configuration{
					ICEServers: []ICEServer{
						{
							URLs: []string{
								"stun:stun.l.google.com:19302",
							},
							Username: "unittest",
						},
					},
					ICETransportPolicy:   ICETransportPolicyAll,
					BundlePolicy:         BundlePolicyBalanced,
					RTCPMuxPolicy:        RTCPMuxPolicyRequire,
					PeerIdentity:         "unittest",
					ICECandidatePoolSize: 5,
				})
				if err != nil {
					return pc, err
				}

				return pc, nil
			},
			config:  Configuration{},
			wantErr: nil,
		},
		{
			name: "closed connection",
			init: func() (*PeerConnection, error) {
				pc, err := NewPeerConnection(Configuration{})
				assert.Nil(t, err)

				err = pc.Close()
				assert.Nil(t, err)
				return pc, err
			},
			config:  Configuration{},
			wantErr: &rtcerr.InvalidStateError{Err: ErrConnectionClosed},
		},
		{
			name: "update PeerIdentity",
			init: func() (*PeerConnection, error) {
				return NewPeerConnection(Configuration{})
			},
			config: Configuration{
				PeerIdentity: "unittest",
			},
			wantErr: &rtcerr.InvalidModificationError{Err: ErrModifyingPeerIdentity},
		},
		{
			name: "update BundlePolicy",
			init: func() (*PeerConnection, error) {
				return NewPeerConnection(Configuration{})
			},
			config: Configuration{
				BundlePolicy: BundlePolicyMaxCompat,
			},
			wantErr: &rtcerr.InvalidModificationError{Err: ErrModifyingBundlePolicy},
		},
		{
			name: "update RTCPMuxPolicy",
			init: func() (*PeerConnection, error) {
				return NewPeerConnection(Configuration{})
			},
			config: Configuration{
				RTCPMuxPolicy: RTCPMuxPolicyNegotiate,
			},
			wantErr: &rtcerr.InvalidModificationError{Err: ErrModifyingRTCPMuxPolicy},
		},
		{
			name: "update ICECandidatePoolSize",
			init: func() (*PeerConnection, error) {
				pc, err := NewPeerConnection(Configuration{
					ICECandidatePoolSize: 0,
				})
				if err != nil {
					return pc, err
				}
				offer, err := pc.CreateOffer(nil)
				if err != nil {
					return pc, err
				}
				err = pc.SetLocalDescription(offer)
				if err != nil {
					return pc, err
				}
				return pc, nil
			},
			config: Configuration{
				ICECandidatePoolSize: 1,
			},
			wantErr: &rtcerr.InvalidModificationError{Err: ErrModifyingICECandidatePoolSize},
		},
		{
			name: "update ICEServers, no TURN credentials",
			init: func() (*PeerConnection, error) {
				return NewPeerConnection(Configuration{})
			},
			config: Configuration{
				ICEServers: []ICEServer{
					{
						URLs: []string{
							"stun:stun.l.google.com:19302",
							"turns:google.de?transport=tcp",
						},
						Username: "unittest",
					},
				},
			},
			wantErr: &rtcerr.InvalidAccessError{Err: ErrNoTurnCredencials},
		},
	} {
		pc, err := test.init()
		if err != nil {
			t.Errorf("SetConfiguration %q: init failed: %v", test.name, err)
		}

		err = pc.SetConfiguration(test.config)
		if got, want := err, test.wantErr; !reflect.DeepEqual(got, want) {
			t.Errorf("SetConfiguration %q: err = %v, want %v", test.name, got, want)
		}
	}
}
