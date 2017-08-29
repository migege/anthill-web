<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="/static/css/bootstrap.min.css" />
    <script type="text/javascript" src="/static/js/jquery.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <script type="text/javascript" src="/static/js/status-ws.js"></script>
</head>

<body>
    <div class="container">
        <div class="row">
            <h1>Realtime Status</h1>
        </div>
        <div class="row">
            <div class="input-group">
                <input id="tag" type="text" class="form-control" placeholder="type TAG here..." aria-label="type TAG here..." value="{{ .Tag }}"/>
                <span class="input-group-btn"><button id="btn-go" class="btn btn-secondary" type="button">Go!</button></span>
            </div>
        </div>
        <div class="row">
            <table class="table table-hover" id="tt">
            </table>
            <div id="extra">
            </div>
        </div>
    </div>
</body>

</html>
