'use strict';

$(document).ready(function() {

    console.debug('Shen is ready!');

    $('#input-username').change(function() {
        if ($(this).val() == '') {
            $('#input-message').hide(400);
            $('#input-send').hide(400);
        } else {
            $('#input-message').show(600);
            $('#input-send').show(600);
        }
    });

    $('#input-send').click(function() {
        send();
    });

    $('#input-message').keyup(function(e) {
        if (e.keyCode == 13 && !e.shiftKey) {
            send();
        }
    });

    connect({
      onConnect: renderConnectCallback,
      onClose: renderDisconnectCallback,
      onMessage: renderMessageCallback,
    });

});

$(window).on('beforeunload', function() {

    console.debug('Shen is closing...');
    disconnect();

});

function send() {

    var username = $('#input-username').val();
    var message = $('#input-message').val();
    $('#input-message').val('');

    onWrite(username, message);

    console.log('@' + username + ': ' + message);

}
