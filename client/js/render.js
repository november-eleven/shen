'use strict';

function onErrorHandler(e) {
    console.error(e);
}

function renderConnectCallback() {
    $('#input-send').prop('disabled', false);
    $('#input-message').prop('disabled', false);
}

function renderDisconnectCallback() {
    $('#input-send').prop('disabled', true);
    $('#input-message').prop('disabled', true);
}

function renderMessageCallback(message) {
    $(renderMessageLine(message)).hide().appendTo($('#output-channel')).show('fast');
    $('#output-wrapper').show(750);
}

function renderMessageLine(payload) {

    var style = '';
    var username = payload.username || 'John Doe';
    var message = payload.message || '';
    var me = payload.me || false;

    if(message == '') {
        return '';
    }

    if(me) {
        style = 'style="text-align: right;"';
    }

    return '<li class="list-group-item" ' + style + '><strong>@' + username + ':</strong> ' + message + '</li>';

}
