const wrtc = require("wrtc");

global.window = {
  RTCPeerConnection: wrtc.RTCPeerConnection
};

global.RTCPeerConnection = wrtc.RTCPeerConnection;
