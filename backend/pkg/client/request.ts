// HTTP request tests

// Check bad request handler
fetch("http://localhost:8000/ws").then((res) => {
  return console.log(res);
});

// Check not found handler
fetch("http://localhost:8000/nothing").then((res) => {
  return console.log(res);
});
