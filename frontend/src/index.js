const grpc = require('@improbable-eng/grpc-web').grpc;

const {RegisterRequest, LoginRequest, IsAdminRequest} = require('../protos/gen/auth/auth_pb.js');
const {AuthClient} = require('../protos/gen/auth/auth_grpc_web_pb.js');

var authClient = new AuthClient('http://localhost:8080', {}, {format: 'text'})

const metadata = {};

window.handleRegister = function handleRegister() {
    const email = document.getElementById('regEmail').value;
    const password = document.getElementById('regPassword').value;
    setResult('registerResult', 'Загрузка...');
    
    var request = new RegisterRequest();
    request.setEmail(email);
    request.setPassword(password);

    authClient.register(request, metadata, function(err, response) {
      if (err) {
        console.log(err);
        if (err.code == grpc.Code.Unknown) {
          setResult('registerResult', "Нет соединения с сервером");
        } else if (err.code == grpc.Code.InvalidArgument) {
          setResult('registerResult', "Неправильный ввод: " + err.message);
        } else if (err.code == grpc.Code.AlreadyExists) {
          setResult('registerResult', "Пользователь с таким email уже существует");
        } else {
          setResult('registerResult', err);
        }
      } else {
        console.log('Taked user_id =' + response.getUserId());
        // setResult('registerResult', 'ID пользователя: ' + response.getUserId())
        window.location.replace('/login')
      }
    });
}

window.handleLogin = function handleLogin() {
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
        if (err.code == grpc.Code.Unknown) {
          setResult('loginResult', "Нет соединения с сервером");
        } else if (err.code == grpc.Code.InvalidArgument) {
          setResult('loginResult', "Неправильный ввод: " + err.message);
        } else {
          setResult('loginResult', err);
        }
      } else {
        console.log('Taked token = ' + response.getToken());
        // setResult('loginResult', 'Полученный токен: ' + response.getToken())
        localStorage.setItem('token', response.getToken());
        window.location.replace('/')
      }
    });
}

window.handleCheckAdmin = function handleCheckAdmin() {
    const userId = document.getElementById('adminUserId').value;
    setResult('adminResult', 'Загрузка...');
    
    var request = new IsAdminRequest();
    request.setUserId(userId);

    authClient.isAdmin(request, metadata, function(err, response) {
      if (err) {
        console.log(err)
        if (err.code == grpc.Code.Unknown) {
          setResult('adminResult', "Нет соединения с сервером");
        } else if (err.code == grpc.Code.InvalidArgument) {
          setResult('adminResult', "Неправильный ввод" + err.message);
        } else {
          setResult('adminResult', err);
        }
      } else {
        console.log('IS ADMIN = ' + response.getIsAdmin());
        setResult('adminResult', 'IS ADMIN = ' + response.getIsAdmin())
      }
    });
}

function setResult(elementId, text) {
    document.getElementById(elementId).textContent = text;
}