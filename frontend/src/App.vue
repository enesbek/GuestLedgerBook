<template>
  <div>
    <h1 class="title">GUEST LEDGER BOOK</h1>
    <div class="main">
      
      <form @submit.prevent="add"  class="formElement">
        <p>Email: <input v-model="item.title"  class="inputEmail" placeholder="test@gmail.com"></p><br>
        <textarea v-model="item.message" class="inputMessage" placeholder="Test Message"></textarea><br>
        <button v-on:click="submit" class="button">Submit</button>
      </form>
      <div class="list">
        <div class="items">
          <div v-for="item of items" :key="item.id"  class="item">
            <p><b>Title: </b> {{ item.title }}</p>
            <p><b>Message: </b> {{ item.message }}</p>
          </div>
        </div>
        <button v-on:click="previousPage" class="button">PREVIOUS</button>
        <button v-on:click="nextPage" class="button">NEXT</button>
      </div>
    </div>
  </div>
</template>

<script>
//import axios from 'axios';
export default {
  name: 'App',
  data() {
    return {
      item: {
        title: null,
        message: null
      },
      items: [],
      page: 1
    }
  },
  methods: {
    async submit() {

      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({"title": this.item.title, "message": this.item.message})
      };

      await fetch("http://localhost:3000/create", requestOptions)
      await this.refreshPage()

      /*axios.post("http://localhost:3000/create", )
        .then(response => console.log(response.data))
        .then(this.refreshPage)*/

    }, 
    async refreshPage() {

      const response = await fetch("http://localhost:3000/" + this.page)
      const data = await response.json()
      this.items = await JSON.parse(JSON.stringify(data)).guests
      
      /*.then(response => {
        this.temp = response.data
        this.items = JSON.parse(JSON.stringify(this.temp)).guests
      })*/

    },
    nextPage(){
      this.page++
      this.refreshPage(this.page)
    },
    previousPage(){
      if(this.page > 1) {
        this.page--
        this.refreshPage(this.page)
      }
      
    }
  },
  mounted() {
    this.refreshPage()
  }
}
</script>

<style lang="postcss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: rgba(4, 120, 87);
}

.title{
  @apply m-5 text-3xl;
}

.main{
  @apply min-h-screen grid grid-cols-3 gap-3 items-center;
}

.formElement{
  @apply h-full bg-green-200 items-center text-black;
}

.inputEmail{
  @apply rounded mt-20 w-1/2 h-8 placeholder-black pl-2;
}

.inputMessage{
  @apply rounded w-4/5 h-64 placeholder-black pl-2 pt-2;
}

.button{
  @apply bg-green-500 text-white font-bold py-2 px-4 rounded w-1/3 m-5;
}

.list{
  @apply col-span-2 bg-green-800 h-full text-black;
}

.item{
  @apply bg-white w-4/5 h-60 mt-5 mb-0 ml-20 rounded text-left pl-2 pt-2;
}

.items{
  @apply grid grid-cols-2 gap-5;
}
</style>
