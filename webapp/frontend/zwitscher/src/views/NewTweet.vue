<template>
  <div class="container">
    <h1 class="title is-1 mb-0 pt-5"><router-link style="color: #2c3e50" :to="'/'">Zwitscher</router-link></h1>
    <div id="nav" class="mt-0">
      <router-link :to="'/timeline/'+$route.params.id">Timeline</router-link> |
      <router-link :to="'/tweets/'+$route.params.id">Tweets</router-link> |
      <router-link :to="'/connections/'+$route.params.id">Connections</router-link> |
      <router-link :to="'/new/'+$route.params.id">New Tweet</router-link>
    </div>

    <h2 class="subtitle is-2">Tweet as @{{ $route.params.id }}</h2>
    
    <form>
      <div class="field">
        <label class="label">Author</label>
        <div class="control">
          <input class="input" type="text" placeholder="Author..." v-model="author">
        </div>
        <p class="help">This is the author that will be displayed on your tweet.</p>
      </div>
      <div class="field">
        <label class="label">Content</label>
        <div class="control">
          <textarea class="textarea" placeholder="e.g. Hello world" v-model="content"></textarea>
        </div>
        <p class="help">This is the of your tweet.</p>
      </div>
      <div class="control">
        <a class="button is-primary" v-on:click="send_tweet()">Tweet</a>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'NewTweet',
  components: {
  },
  data() {
    return {
      author: "",
      content: "",
    }
  },
  computed: {
    tweet() {
      return {
        schema: {},
        payload: {
          user_id: this.$route.params.id,
          author: this.author,
          content: this.content,
          number_of_likes: 0,
          number_of_shares: 0,
        }
      }
    }
  },
  methods: {
    send_tweet() {
      axios
        .post("http://192.168.0.232:7080/zwitscherlos", this.tweet)
        .then(() => {
          this.author = "";
          this.content = "";
        });
    }
  }
}
</script>