const socket = new WebSocket('ws://localhost:1009/auctee/ws');

const connect = (cb) => {
  console.log('connecting');

  socket.onopen = () => {
    console.log('Successfully Connected');
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = (event) => {
    console.log('Socket Closed Connection: ', event);
  };

  socket.onerror = (error) => {
    console.log('Socket Error: ', error);
  };
};

// call api to get data
const sendMsg = (body) => {
  JSON.stringify(body);
  console.log("line 27",body);
  socket.send("helo");
};

export { connect, sendMsg };
