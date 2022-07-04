import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VueSocketIO from 'vue-3-socket.io'
import SocketIO from 'socket.io-client'
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";
import "bootstrap-icons/font/bootstrap-icons.css"
var socket =  new VueSocketIO({
    debug: true,
    connection: SocketIO('http://localhost:9000', {
        query: `token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoibWFuaCIsImlhdCI6MTY1NjU1NzUxNywiZXhwIjoxNjU2NjQzOTE3fQ.B3A_Ms8ToDFe0tWBXM__1NC_FEwmiCtinXM4X-tT-X4`
    }), //options object is Optional
    vuex: {
      store,
      actionPrefix: "SOCKET_",
      mutationPrefix: "SOCKET_"
    }
  })
createApp(App).use(store).use(socket).use(router).mount('#app')
