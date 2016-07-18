'use strict';

var localConnection;
var remoteConnection;
var sendChannel;
var receiveChannel;
var pcConstraint;
var dataConstraint;
var servers;

$(document).ready(function() {

    console.debug("Shen is ready!");

    $("#input-username").change(function() {
        if ($(this).val() == "") {
            $("#input-message").hide();
            $("#input-send").hide();
        } else {
            $("#input-message").show(600);
            $("#input-send").show(600);
        }
    });

    $("#input-send").click(function() {
        send();
    });

    $("#input-message").keyup(function(e) {
        if (e.keyCode == 13 && !e.shiftKey) {
            send();
        }
    });

    servers = null;
    pcConstraint = null;
    dataConstraint = null;
    console.debug('Using SCTP based data channels');

    window.localConnection = localConnection = new RTCPeerConnection(servers, pcConstraint);

    console.debug('Created local peer connection object localConnection');

    sendChannel = localConnection.createDataChannel('shen-channel', dataConstraint);
    console.debug('Created send data channel');

    localConnection.onicecandidate = iceCallback1;
    sendChannel.onopen = onSendChannelStateChange;
    sendChannel.onclose = onSendChannelStateChange;


    window.remoteConnection = remoteConnection = new RTCPeerConnection(servers, pcConstraint);
    console.debug('Created remote peer connection object remoteConnection');

    remoteConnection.onicecandidate = iceCallback2;
    remoteConnection.ondatachannel = receiveChannelCallback;

    localConnection.createOffer().then(
        gotDescription1,
        onCreateSessionDescriptionError
    );

});

$(window).on('beforeunload', function() {

    console.debug('Shen is closing...');

    if(sendChannel != null) {
        sendChannel.close();
    }

    if(receiveChannel != null) {
        receiveChannel.close();
    }

    if(localConnection != null) {
        localConnection.close();
    }

    if(remoteConnection != null) {
        remoteConnection.close();
    }

    sendChannel = null;
    receiveChannel = null;
    localConnection = null;
    remoteConnection = null;

});

function send() {

    var username = $("#input-username").val();
    var message = $("#input-message").val();
    $("#input-message").val('');

    if(sendChannel != null) {
        sendChannel.send(JSON.stringify({username: username, message: message}));
    }

    console.log('@' + username + ': ' + message);

}

function onSendChannelStateChange() {

    var readyState = sendChannel.readyState;
    console.debug('Send channel state is: ' + readyState);

    if (readyState === 'open') {
        $("#input-send").prop("disabled", false);
        $("#input-message").prop("disabled", false);
    } else {
        $("#input-send").prop("disabled", true);
        $("#input-message").prop("disabled", true);
    }

}

function iceCallback1(event) {
  console.debug('local ice callback');
  if (event.candidate) {
    remoteConnection.addIceCandidate(
      event.candidate
    ).then(
      onAddIceCandidateSuccess,
      onAddIceCandidateError
    );
    console.debug('Local ICE candidate: \n' + event.candidate.candidate);
  }
}

function iceCallback2(event) {
  console.debug('remote ice callback');
  if (event.candidate) {
    localConnection.addIceCandidate(
      event.candidate
    ).then(
      onAddIceCandidateSuccess,
      onAddIceCandidateError
    );
    console.debug('Remote ICE candidate: \n ' + event.candidate.candidate);
  }
}

function onAddIceCandidateSuccess() {
  console.debug('AddIceCandidate success.');
}

function onAddIceCandidateError(error) {
  console.debug('Failed to add Ice Candidate: ' + error.toString());
}

function receiveChannelCallback(event) {
  console.debug('Receive Channel Callback');
  receiveChannel = event.channel;
  receiveChannel.onmessage = onReceiveMessageCallback;
  receiveChannel.onopen = onReceiveChannelStateChange;
  receiveChannel.onclose = onReceiveChannelStateChange;
}

function onReceiveMessageCallback(event) {

    console.log('Message received: ' + event.data);
    var payload = JSON.parse(event.data);

    $("#output-channel").append(renderMessageLine(payload));
    $("#output-wrapper").show(750);

}

function onReceiveChannelStateChange() {

    var readyState = receiveChannel.readyState;
    console.debug('Receive channel state is: ' + readyState);

}

function gotDescription1(desc) {
  localConnection.setLocalDescription(desc);
  console.debug('Offer from localConnection \n' + desc.sdp);
  remoteConnection.setRemoteDescription(desc);
  remoteConnection.createAnswer().then(
    gotDescription2,
    onCreateSessionDescriptionError
  );
}

function gotDescription2(desc) {
  remoteConnection.setLocalDescription(desc);
  console.debug('Answer from remoteConnection \n' + desc.sdp);
  localConnection.setRemoteDescription(desc);
}

function onCreateSessionDescriptionError(error) {
  console.debug('Failed to create session description: ' + error.toString());
}


