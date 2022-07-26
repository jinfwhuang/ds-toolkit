/**
 * @fileoverview gRPC-Web generated client stub for jinfwhuang.dstoolkit.identity
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('../google/api/annotations_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.jinfwhuang = {};
proto.jinfwhuang.dstoolkit = {};
proto.jinfwhuang.dstoolkit.identity = require('./login_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.jinfwhuang.dstoolkit.identity.IdentityClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.jinfwhuang.dstoolkit.identity.IdentityPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 *   !proto.jinfwhuang.dstoolkit.identity.LoginMessage,
 *   !proto.jinfwhuang.dstoolkit.identity.LoginMessage>}
 */
const methodDescriptor_Identity_RequestLogin = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.Identity/RequestLogin',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.LoginMessage,
  proto.jinfwhuang.dstoolkit.identity.LoginMessage,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.LoginMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.LoginMessage.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.LoginMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.LoginMessage)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.LoginMessage>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.IdentityClient.prototype.requestLogin =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.Identity/RequestLogin',
      request,
      metadata || {},
      methodDescriptor_Identity_RequestLogin,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.LoginMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.LoginMessage>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.IdentityPromiseClient.prototype.requestLogin =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.Identity/RequestLogin',
      request,
      metadata || {},
      methodDescriptor_Identity_RequestLogin);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.jinfwhuang.dstoolkit.identity.LoginMessage,
 *   !proto.jinfwhuang.dstoolkit.identity.LoginResp>}
 */
const methodDescriptor_Identity_Login = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.Identity/Login',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.LoginMessage,
  proto.jinfwhuang.dstoolkit.identity.LoginResp,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.LoginMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.LoginResp.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.LoginMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.LoginResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.LoginResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.IdentityClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.Identity/Login',
      request,
      metadata || {},
      methodDescriptor_Identity_Login,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.LoginMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.LoginResp>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.IdentityPromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.Identity/Login',
      request,
      metadata || {},
      methodDescriptor_Identity_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.jinfwhuang.dstoolkit.identity.LoginMessage>}
 */
const methodDescriptor_Identity_Debug = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.Identity/Debug',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.jinfwhuang.dstoolkit.identity.LoginMessage,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.LoginMessage.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.LoginMessage)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.LoginMessage>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.IdentityClient.prototype.debug =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.Identity/Debug',
      request,
      metadata || {},
      methodDescriptor_Identity_Debug,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.LoginMessage>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.IdentityPromiseClient.prototype.debug =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.Identity/Debug',
      request,
      metadata || {},
      methodDescriptor_Identity_Debug);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

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
 *   !proto.google.protobuf.Empty,
 *   !proto.jinfwhuang.dstoolkit.identity.UserList>}
 */
const methodDescriptor_UserRegistryLogin_ListAllUsers = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/ListAllUsers',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.jinfwhuang.dstoolkit.identity.UserList,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.UserList.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.UserList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.UserList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient.prototype.listAllUsers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/ListAllUsers',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_ListAllUsers,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.UserList>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient.prototype.listAllUsers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/ListAllUsers',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_ListAllUsers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.jinfwhuang.dstoolkit.identity.User,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_UserRegistryLogin_AddUser = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/AddUser',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.User,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient.prototype.addUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/AddUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_AddUser,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient.prototype.addUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/AddUser',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_AddUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.jinfwhuang.dstoolkit.identity.PubKey,
 *   !proto.jinfwhuang.dstoolkit.identity.User>}
 */
const methodDescriptor_UserRegistryLogin_GetUserByPubKey = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/GetUserByPubKey',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.PubKey,
  proto.jinfwhuang.dstoolkit.identity.User,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.PubKey} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.User.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.PubKey} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient.prototype.getUserByPubKey =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/GetUserByPubKey',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_GetUserByPubKey,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.PubKey} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.User>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient.prototype.getUserByPubKey =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/GetUserByPubKey',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_GetUserByPubKey);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.jinfwhuang.dstoolkit.identity.UserName,
 *   !proto.jinfwhuang.dstoolkit.identity.User>}
 */
const methodDescriptor_UserRegistryLogin_GetUserByUserName = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/GetUserByUserName',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.UserName,
  proto.jinfwhuang.dstoolkit.identity.User,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.UserName} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.User.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.UserName} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient.prototype.getUserByUserName =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/GetUserByUserName',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_GetUserByUserName,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.UserName} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.User>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient.prototype.getUserByUserName =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/GetUserByUserName',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_GetUserByUserName);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.jinfwhuang.dstoolkit.identity.UserName,
 *   !proto.jinfwhuang.dstoolkit.identity.User>}
 */
const methodDescriptor_UserRegistryLogin_RequestLogin = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/RequestLogin',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.UserName,
  proto.jinfwhuang.dstoolkit.identity.User,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.UserName} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.User.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.UserName} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient.prototype.requestLogin =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/RequestLogin',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_RequestLogin,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.UserName} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.User>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient.prototype.requestLogin =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/RequestLogin',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_RequestLogin);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.jinfwhuang.dstoolkit.identity.User,
 *   !proto.jinfwhuang.dstoolkit.identity.LoginResp>}
 */
const methodDescriptor_UserRegistryLogin_Login = new grpc.web.MethodDescriptor(
  '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/Login',
  grpc.web.MethodType.UNARY,
  proto.jinfwhuang.dstoolkit.identity.User,
  proto.jinfwhuang.dstoolkit.identity.LoginResp,
  /**
   * @param {!proto.jinfwhuang.dstoolkit.identity.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.jinfwhuang.dstoolkit.identity.LoginResp.deserializeBinary
);


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.jinfwhuang.dstoolkit.identity.LoginResp)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.jinfwhuang.dstoolkit.identity.LoginResp>|undefined}
 *     The XHR Node Readable Stream
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/Login',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_Login,
      callback);
};


/**
 * @param {!proto.jinfwhuang.dstoolkit.identity.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.jinfwhuang.dstoolkit.identity.LoginResp>}
 *     Promise that resolves to the response
 */
proto.jinfwhuang.dstoolkit.identity.UserRegistryLoginPromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/jinfwhuang.dstoolkit.identity.UserRegistryLogin/Login',
      request,
      metadata || {},
      methodDescriptor_UserRegistryLogin_Login);
};


module.exports = proto.jinfwhuang.dstoolkit.identity;

