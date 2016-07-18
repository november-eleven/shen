'use strict';

function renderMessageLine(payload) {

    var username = payload.username || 'John Doe';
    var message = payload.message || '';

    if(message == '') {
        return ''
    }

    return '<li class="list-group-item"><strong>@' + username + ':</strong> ' + message + '</li>';

}
