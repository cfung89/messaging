var ws = new WebSocket("ws://localhost:8000/ws");

ws.addEventListener("open", (event) => {
  console.log("WebSocket is open");
  ws.send("Hello Server");
  console.log("Sending Hello");
  // ws.close();
});

ws.addEventListener("close", (event) => {
  console.log("WebSocket is closed");
});

ws.addEventListener("message", (event) => {
  console.log("WebSocket data:", event.data);
});

ws.addEventListener("error", (event) => {
  console.error("WebSocket error:", event);
});
