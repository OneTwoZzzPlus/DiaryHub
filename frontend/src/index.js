// const { NodeHttpTransport } = require('@improbable-eng/grpc-web-node-http-transport');
const grpc = require('@improbable-eng/grpc-web').grpc;
// grpc.setDefaultTransport(NodeHttpTransport());
// global.XMLHttpRequest = require('xhr2');

const {RegisterRequest, LoginRequest, IsAdminRequest} = require('../protos/gen/auth/auth_pb.js');
const {AuthClient} = require('../protos/gen/auth/auth_grpc_web_pb.js');

var authClient = new AuthClient('http://localhost:8080', {}, {format: 'text'})

const metadata = {};

function handleRegister() {
    const email = document.getElementById('regEmail').value;
    const password = document.getElementById('regPassword').value;
    setResult('registerResult', 'Загрузка...');
    
    var request = new RegisterRequest();
    request.setEmail(email);
    request.setPassword(password);

    authClient.register(request, metadata, function(err, response) {
      if (err) {
        console.log(err)
        setResult('registerResult', err)
      } else {
        console.log('USER ID =' + response.getUserId());
        setResult('registerResult', 'USER ID =' + response.getUserId())
      }
    });
}

function handleLogin() {
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;
    setResult('loginResult', 'Загрузка...');
    
    var request = new LoginRequest();
    request.setEmail(email);
    request.setPassword(password);
    request.setAppId(1);

    authClient.login(request, metadata, function(err, response) {
      if (err) {
        console.log(err)
        setResult('loginResult', err)
      } else {
        console.log('TOKEN = ' + response.getToken());
        setResult('loginResult', 'TOKEN = ' + response.getToken())
      }
    });
}

function handleCheckAdmin() {
    const userId = document.getElementById('adminUserId').value;
    setResult('adminResult', 'Загрузка...');
    
    var request = new IsAdminRequest();
    request.setUserId(userId);

    authClient.isAdmin(request, metadata, function(err, response) {
      if (err) {
        console.log(err)
        setResult('adminResult', err)
      } else {
        console.log('IS ADMIN = ' + response.getIsAdmin());
        setResult('adminResult', 'IS ADMIN = ' + response.getIsAdmin())
      }
    });

}

function setResult(elementId, text) {
    document.getElementById(elementId).textContent = text;
}

document.getElementById('registerBtn').addEventListener('click', handleRegister)
document.getElementById('loginBtn').addEventListener('click', handleLogin)
document.getElementById('adminCheckBtn').addEventListener('click', handleCheckAdmin)