import WebSocket = require('ws')
var ws = new WebSocket("ws://127.0.0.1:8899/echo");  
// onOpen被觸發時, 去嘗試連線 
ws.onopen = function(evt) {  
    console.log("Connection open ...");  
    ws.send("Hello WebSockets!");  
};  
// onMessage被觸發時, 來接收ws server傳來的訊息  
ws.onmessage = function(evt) {  
    console.log("Received Message: " + evt.data);  
};  
// 由ws server發出的onClose事件  
ws.onclose = function(evt) {  
    console.log("Connection closed.");  
};  
// 每秒發出一個現在時間的訊息
var timeInterval = setInterval(() => ws.send(""), 1000)