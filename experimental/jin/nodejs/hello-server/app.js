const http = require('http');

const hostname = '0.0.0.0';
const port = 8080;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  const resp_content = "hello at: " + (new Date()).getTime() + "\n"
  console.log(resp_content)
  res.end(resp_content);
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
