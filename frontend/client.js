const { NodeHttpTransport } = require('@improbable-eng/grpc-web-node-http-transport');
const grpc = require('@improbable-eng/grpc-web').grpc;

// Установка транспорта для Node.js
grpc.setDefaultTransport(NodeHttpTransport());

// Добавьте полифилл для XMLHttpRequest (решает вашу ошибку)
global.XMLHttpRequest = require('xhr2');

const {RegisterRequest, RegisterResponse} = require('./protos/gen/auth/auth_pb.js');
const {AuthClient} = require('./protos/gen/auth/auth_grpc_web_pb.js');

grpc.setDefaultTransport(NodeHttpTransport());

var authClient = new AuthClient('http://localhost:8080', {}, {format: 'text'})  // или 'binary'

var request = new RegisterRequest();
request.setEmail("ivan22@mail.com")
request.setPassword("password")

var metadata = {'custom-header-1': 'value1'};
authClient.register(request, metadata, function(err, response) {
  if (err) {
    console.log(err.code);
    console.log(err.message);
    console.log(err)
  } else {
    console.log(response.getUserId());
  }
});