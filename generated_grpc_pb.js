// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var generated_pb = require('./generated_pb.js');

function serialize_gomad_AddIngredientRequest(arg) {
  if (!(arg instanceof generated_pb.AddIngredientRequest)) {
    throw new Error('Expected argument of type gomad.AddIngredientRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_gomad_AddIngredientRequest(buffer_arg) {
  return generated_pb.AddIngredientRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gomad_NutritionalReportResponse(arg) {
  if (!(arg instanceof generated_pb.NutritionalReportResponse)) {
    throw new Error('Expected argument of type gomad.NutritionalReportResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_gomad_NutritionalReportResponse(buffer_arg) {
  return generated_pb.NutritionalReportResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gomad_Pizza(arg) {
  if (!(arg instanceof generated_pb.Pizza)) {
    throw new Error('Expected argument of type gomad.Pizza');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_gomad_Pizza(buffer_arg) {
  return generated_pb.Pizza.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gomad_RemoveIngredientRequest(arg) {
  if (!(arg instanceof generated_pb.RemoveIngredientRequest)) {
    throw new Error('Expected argument of type gomad.RemoveIngredientRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_gomad_RemoveIngredientRequest(buffer_arg) {
  return generated_pb.RemoveIngredientRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var GomadServiceService = exports.GomadServiceService = {
  addIngredient: {
    path: '/gomad.GomadService/AddIngredient',
    requestStream: false,
    responseStream: false,
    requestType: generated_pb.AddIngredientRequest,
    responseType: generated_pb.Pizza,
    requestSerialize: serialize_gomad_AddIngredientRequest,
    requestDeserialize: deserialize_gomad_AddIngredientRequest,
    responseSerialize: serialize_gomad_Pizza,
    responseDeserialize: deserialize_gomad_Pizza,
  },
  nutritionalReport: {
    path: '/gomad.GomadService/NutritionalReport',
    requestStream: false,
    responseStream: false,
    requestType: generated_pb.Pizza,
    responseType: generated_pb.NutritionalReportResponse,
    requestSerialize: serialize_gomad_Pizza,
    requestDeserialize: deserialize_gomad_Pizza,
    responseSerialize: serialize_gomad_NutritionalReportResponse,
    responseDeserialize: deserialize_gomad_NutritionalReportResponse,
  },
  removeIngredient: {
    path: '/gomad.GomadService/RemoveIngredient',
    requestStream: false,
    responseStream: false,
    requestType: generated_pb.RemoveIngredientRequest,
    responseType: generated_pb.Pizza,
    requestSerialize: serialize_gomad_RemoveIngredientRequest,
    requestDeserialize: deserialize_gomad_RemoveIngredientRequest,
    responseSerialize: serialize_gomad_Pizza,
    responseDeserialize: deserialize_gomad_Pizza,
  },
};

exports.GomadServiceClient = grpc.makeGenericClientConstructor(GomadServiceService);
