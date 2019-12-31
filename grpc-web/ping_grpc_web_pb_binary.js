/**
 * @fileoverview gRPC-Web generated client stub for sms.core
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.sms = {};
proto.sms.core = require('./ping_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.sms.core.CoreServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.sms.core.CoreServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.sms.core.PingRequest,
 *   !proto.sms.core.PingResponse>}
 */
const methodDescriptor_CoreService_Ping = new grpc.web.MethodDescriptor(
  '/sms.core.CoreService/Ping',
  grpc.web.MethodType.UNARY,
  proto.sms.core.PingRequest,
  proto.sms.core.PingResponse,
  /**
   * @param {!proto.sms.core.PingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.sms.core.PingResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.sms.core.PingRequest,
 *   !proto.sms.core.PingResponse>}
 */
const methodInfo_CoreService_Ping = new grpc.web.AbstractClientBase.MethodInfo(
  proto.sms.core.PingResponse,
  /**
   * @param {!proto.sms.core.PingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.sms.core.PingResponse.deserializeBinary
);


/**
 * @param {!proto.sms.core.PingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.sms.core.PingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.sms.core.PingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.sms.core.CoreServiceClient.prototype.ping =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/sms.core.CoreService/Ping',
      request,
      metadata || {},
      methodDescriptor_CoreService_Ping,
      callback);
};


/**
 * @param {!proto.sms.core.PingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.sms.core.PingResponse>}
 *     A native promise that resolves to the response
 */
proto.sms.core.CoreServicePromiseClient.prototype.ping =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/sms.core.CoreService/Ping',
      request,
      metadata || {},
      methodDescriptor_CoreService_Ping);
};


module.exports = proto.sms.core;

