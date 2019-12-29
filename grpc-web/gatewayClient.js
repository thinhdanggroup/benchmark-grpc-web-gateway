let Benchmarkify = require("benchmarkify");
var KeepAliveAgent = require('keep-alive-agent');
var request = require('request');

let benchmark = new Benchmarkify("Simple example").printHeader();

var agent = new KeepAliveAgent({ maxSockets: 1 });

var options = {
  agent:agent,
  headers: {"Connection":"Keep-Alive"}
}

request.defaults(options);

// Create a test suite
let bench = benchmark.createSuite("Call HTTP", { time: 30000 });


bench.add("Ping API", done => {
  request('http://localhost:8800/ping/123', function (error, response, body) {
    if (error) {
      console.log('error:', error); // Print the error if one occurred
      return;
    }
    done();
  });
});

bench.run();
