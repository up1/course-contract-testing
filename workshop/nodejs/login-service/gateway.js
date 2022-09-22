const axios = require('axios');

function UserService(api_server_url) {
  this.AXIOS = axios.create({ baseURL: api_server_url });
  this.getUserById = (id) => {
    return this.AXIOS.get(`/users/${id}`)
    .then((res) => res.data)
    .catch( e =>  { return {error: "User not found"} });
  };
}
module.exports = UserService;