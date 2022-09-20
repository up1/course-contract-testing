const UserService = require('./gateway');

const service = new UserService('http://127.0.0.1:3000');
service.getUserById(1)
   .then(user => consloe.consloe(user));
