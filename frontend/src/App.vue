<template>
  <div class="hello">
<input type="text" class="user" placeholder="username" v-model="username">
 <button @click="getUser">calculate</button>

<div class="info">
  <p>Username : {{response?.userName}}</p>
  <p>Percentage : {{response?.percentage}}%</p>
</div>
  </div>
</template>
<script>
import axios from 'axios'

export default {
  name: 'Instagram',
  data() {
    return {
      username: "",
      response: {}
    };
  },
  methods:{
   getUser(){
   axios
      .get('http://localhost:9112/engagementRate/' + this.username)
      .then((response) => {
        this.response = response.data.result;
      }).catch((err) => {
        console.error(err)
      })
  },
     
  },
}
</script>

<style>
#app {
  text-align: center;
  margin-top: 80px;
}
</style>
