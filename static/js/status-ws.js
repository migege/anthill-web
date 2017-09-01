var wsUri;
var ws;

window.addEventListener("load", function(evt) {
    wsUri = "ws://" + window.location.host + "/status/stream";

    var print = function(data) {
        message = data.info
        parts = message.split("`")
        d = JSON.parse(parts[1])
        cols = d.cols
        rows = d.rows
        content = ""
        content += "<thead><tr>"
        for (var i = 0; i < cols.length; i++) {
            content += "<th>" + cols[i] + "</th>"
        }
        content += "</tr></thead>"
        content += "<tbody>"
        for (var j = 0; j < rows.length; j++) {
            content += "<tr>"
            for (var k = 0; k < rows[j].length; k++) {
                content += "<td>" + rows[j][k] + "</td>"
            }
            content += "</tr>"
        }
        content += "</tbody>"
        $('#tt').html(content)

        lines = parts[2].split("\n")
        lines.push('Last updated at: ' + new Date(parseInt(data.ts, 10)).toISOString())
        $('#extra').html("<p>" + lines.join("</p><p>") + "</p>")
    };

    var parseInfo = function(evt) {
        return JSON.parse(evt.data)
    };

    var newSocket = function() {
        ws = new WebSocket(wsUri);
        ws.onopen = function(evt) {
            if ($('#antid').val().length > 0) {
                var req = {
                    info: $('#antid').val()
                }
                ws.send(JSON.stringify(req))
            }
        }
        ws.onclose = function(evt) {
            reconnectSocket()
        }
        ws.onmessage = function(evt) {
            print(parseInfo(evt))
        }
        ws.onerror = function(evt) {}
    };

    var reconnectSocket = function() {
        if (!ws || ws.readyState == WebSocket.CLOSED) {
            newSocket()
        }
    };

    $('#btn-go').click(function() {
        reconnectSocket()
        return false
    })
})
