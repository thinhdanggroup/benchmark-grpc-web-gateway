const {PingRequest} = require('./ping_pb.js');
const {CoreServiceClient} = require('./ping_grpc_web_pb.js');
global.XMLHttpRequest = require('xhr2');

var client = new CoreServiceClient('http://' + "localhost" + ':8080', null, null);

// simple unary call
var request = new PingRequest();
request.setTimestamp(123);

client.ping(request, {}, (err, response) => {
    console.log(response.getMessage());
});

