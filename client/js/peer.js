'use strict';

var id;
var discovery;
var handshake;

var peers;
var linked;

var handler = {
    onConnect: function() {},
    onClose: function() {},
    onMessage: function(message) {},
};

function connect(options) {

    handler = options || handler;
    peers = {};
    linked = {};

    $.ajax({
            url: '/api/login',
            method: 'POST',
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            crossDomain: true,
    })
    .done(onLogin)
    .fail(onErrorHandler);

}

function onLogin(payload) {

    id = payload.id;

    discovery = setTimeout(onPeerDiscovery, 0);
    handshake = setTimeout(onHandshake, 0);

}

function onPeerDiscovery() {

    console.debug('Execute Peer Discovery...');

    $.ajax({
        url: '/api/peers',
        method: 'GET',
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        crossDomain: true,
    })
    .done(onPeerDiscoverySuccess)
    .fail(onErrorHandler)
    .always(onRegister);

    var max = 60000;
    var min = 100;

    discovery = schedule(onPeerDiscovery, 100, 60000);

}

function onRegister() {

    $.ajax({
        url: '/api/peers',
        method: 'POST',
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: id,
        crossDomain: true,
    })
    .fail(function() {});

}

function onPeerDiscoverySuccess(list) {

    list.map(function(payload) {

        if(payload == id) {
            return;
        }

        if(linked[payload] || peers[payload]) {
            return;
        }

        newPeer(id, payload, true);

    });

}

function onHandshake() {

    $.ajax({
        url: '/api/peers/' + id,
        method: 'GET',
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        crossDomain: true,
    })
    .done(onHandshakeSuccess)
    .fail(onErrorHandler);

    handshake = schedule(onHandshake, 100, 5000);

}

function onHandshakeSuccess(list) {

    list.map(function(payload) {

        if(payload.id == id) {
            return;
        }

        if(linked[payload.id]) {
            return;
        }

        try {

            if(peers[payload.id] === undefined) {
                newPeer(id, payload.id);
            }

            var p = peers[payload.id];
            p.signal(payload.signal);

        } catch(e) {
            console.error('Handshake had an error:', e);
            removePeer(id, payload.id, true);
        }

    });

}

function removePeer(local, remote, clean) {

    var request = JSON.stringify({id: local});

    if(clean) {
        delete linked[remote];
        delete peers[remote];
    }

    $.ajax({
        url: '/api/peers/' + remote,
        method: 'DELETE',
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: request,
        crossDomain: true,
    })
    .fail(function() {});

}

function newPeer(local, remote, initiator) {

    var init = initiator || false;

    var peer = new SimplePeer({
        initiator: init,
        trickle: false,
        channelName: 'shen-channel',
        config: {
            'iceServers': [
                { 'urls': 'stun:stun.services.mozilla.com' },
                { 'urls': 'stun:stun.l.google.com:19302' },
            ],
        },
    });

    peer.on('signal', function(signal) {

        var request = JSON.stringify({id: local, signal: signal});

        $.ajax({
            url: '/api/peers/' + remote,
            method: 'POST',
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            data: request,
            crossDomain: true,
        });

        console.debug('Try to connect on remote peer: ' + remote);

    });

    peer.on('connect', function() {

        linked[remote] = true;

        console.debug('Connect: ' + remote);
        removePeer(id, remote);
        onConnect();

    });

    peer.on('data', function(data) {
        onReceive(JSON.parse(data));
    });

    peer.on('close', function() {

        console.debug('Close: ' + remote);
        removePeer(id, remote, true);
        onClose();

    });

    peers[remote] = peer;
    return peer;

}

function disconnect() {

    var payload = JSON.stringify({id: id});

    $.ajax({
        url: '/api/logout',
        method: 'POST',
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        data: payload,
        crossDomain: true,
    });

    Object.keys(peers).map(function(key) {
        try {

            var peer = peers[key];
            peer.destroy();

        } catch(e) {
            // Silent fail
        }
    });

    peers = {};
    linked = {};
    id = undefined;

}

function onConnect() {
    if (Object.keys(peers).length > 0 && handler != null) {
        handler.onConnect();
    }
}

function onClose() {
    if(Object.keys(peers).length === 0 && handler != null) {
        handler.onClose();
    }
}

function onWrite(username, message) {

    var payload = JSON.stringify({username: username, message: message});

    Object.keys(peers).map(function(key) {
        try {

            var peer = peers[key];
            peer.send(payload);

        } catch(e) {
            delete peers[ key ];
            delete linked[ key ];
        }
    });

    if (handler != null) {
        handler.onMessage({username: username, message: message, me: true});
    }

}

function onReceive(message) {
    if (handler != null) {
        handler.onMessage(message);
    }
}

function schedule(callback, min, max) {
    return setTimeout(callback, (Math.random() * (max - min) + min));
}
