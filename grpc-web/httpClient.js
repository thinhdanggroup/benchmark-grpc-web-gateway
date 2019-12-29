var request = require('request');
var sleep = require('system-sleep');

var counter = 0
var total = 10000

var start = Math.floor(Date.now());

for (var i = 0; i < total ;i++) {
  const finalI = i;
  request('http://localhost:8800/ping/123', function (error, response, body) {
    if (error) {
      console.log('error:', error); // Print the error if one occurred
      return;
    }
    console.log('statusCode:', response && response.statusCode); // Print the response status code if a response was received
    // console.log('body:', body); // Print the HTML for the Google homepage.
    console.log("idx="+ finalI +" curTime= " +(Math.floor(Date.now() ) - start));

  });
  if (i % 5 == 0 && i > 0 ) {
    sleep(10);
}
}
