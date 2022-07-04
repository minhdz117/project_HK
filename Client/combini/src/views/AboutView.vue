<template>
  <div class="about">
    <navbar-component></navbar-component>
    <main class="offcanvas-outside bg-light">
      <div class="container-fluid">
        <h1>This is an about page</h1>
        <input type="text" id="chat" v-model="text"><br>
        <button id="send" @click="send(text)">send</button>
      </div>
    </main>
  </div>
</template>

<script>
import NavbarComponent from '../components/NavbarComponent.vue'
export default {
  name: 'AboutView',
  data() {
    return {
      text: "",
    }
  },
  methods: {
    send(data) {
      // $socket is socket.io-client instance
      this.$socket.emit('user_chat', data)
    }
  },
  components: {
    NavbarComponent
  }
}
</script>