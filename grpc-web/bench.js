var Benchmarkify = require("benchmarkify");
var KeepAliveAgent = require('keep-alive-agent');
var request = require('request');
const {PingRequest} = require('./ping_pb.js');
const {CoreServiceClient: CoreServiceClient} = require('./ping_grpc_web_pb.js');
const {CoreServiceClient: CoreServiceClientBinary} = require('./ping_grpc_web_pb.js');

var benchmark = new Benchmarkify("Benchmark: gRPC-Web vs grpc-gateway vs spring-http-base-grpc").printHeader();
var bench = benchmark.createSuite("Call HTTP", {time: 30000});
global.XMLHttpRequest = require('xhr2');
var agent = new KeepAliveAgent({maxSockets: 1});

var options = {
    agent: agent,
    headers: {"Connection": "Keep-Alive"}
};

request.defaults(options);


// var target = {
//     backend: 'http://localhost:6789/api/ping',
//     gateway: 'http://localhost:8800/ping/123',
//     envoy: 'http://localhost:8080',
// };

var target = {
  backend: 'http://bench-server:6789/api/ping',
  gateway: 'http://bench-gateway:80/ping/123',
  envoy: 'http://bench-envoy:8080',
};

var clientText = new CoreServiceClient(target.envoy, null, null);
var clientBinary = new CoreServiceClientBinary(target.envoy, null, null);

bench.add("Ping to server", done => {
    request.post(target.backend, {
        json: {
            timestamp: 111
        }
    }, (error, res, body) => {
        if (error) {
            console.log('error:', error); // Print the error if one occurred
        }
        done();
    });
});

bench.add("Ping to gateway", done => {
  request(target.gateway, function (error, response, body) {
    if (error) {
      console.log('error:', error); // Print the error if one occurred
    }
    done();
  });
});

//===

var pingRequest = new PingRequest();
pingRequest.setTimestamp(123);

bench.add("Ping with grpc-web text", done => {
  clientText.ping(pingRequest, {}, (err, response) => {
      if (err) {
          console.log(err);
      } 
      done();
  });
});

// ===
// var pingRequest = new PingRequest();
// pingRequest.setTimestamp(123);

// var metadata = {'content-type': 'application/grpc-web+proto'};
var metadata = {};
bench.add("Ping with grpc-web binary", done => {
  clientBinary.ping(pingRequest, metadata, (err, response) => {
        if (err) {
            console.log(err);
        }
        done();
    });
});

bench.run();