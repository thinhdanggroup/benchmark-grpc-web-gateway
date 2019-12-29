var request = require('request');
var sleep = require('system-sleep');
var KeepAliveAgent = require('keep-alive-agent');

var counter = 0
var total = 10000

var start = Math.floor(Date.now());

var agent = new KeepAliveAgent({ maxSockets: 1 });

var options = {
  agent:agent,
  headers: {"Connection":"Keep-Alive"}
}

request.defaults(options);

for (var i = 0; i < total ;i++) {
  const finalI = i;
  request.post('http://localhost:6789/api/ping', {
    json: {
      timestamp: 111
    }
  }, (error, res, body) => {
    if (error) {
      console.log('error:', error); // Print the error if one occurred
      return;
    }
    console.log('statusCode:', res && res.statusCode); // Print the response status code if a response was received
    // console.log('body:', body); // Print the HTML for the Google homepage.
    console.log("idx="+ finalI +" curTime= " +(Math.floor(Date.now() ) - start));

  })
  if (i % 10 == 0 && i > 0 ) {
    sleep(10);
  }
}
