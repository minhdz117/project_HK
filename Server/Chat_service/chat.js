var express = require('express');
var app = express()
// const cors = require('cors')
// app.use(cors())
var server = require('http').Server(app);
const io = require('socket.io')(server,{
    cors: {
        origin: "http://localhost:8080",
        methods: ["GET", "POST"]
    }
});
const socketioJwt = require('socketio-jwt');

app.set('view engine', 'ejs');
// app.use('/public', express.static(path.join(__dirname, 'public')))
io.use(socketioJwt.authorize({
    secret: 'minhdz117',
    handshake: true
}));

io.on('connection', (socket) => {
    console.log(socket.decoded_token.name," connected");
    socket.on('user_chat',(message)=>{
        console.log(`${socket.decoded_token.name} say ${message}`)
        io.emit('server_send',`${socket.decoded_token.name} say ${message}`)
    })
});

app.get("/", (req, res) => {
    res.render("test")
})
server.listen(9000)