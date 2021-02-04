<template>
  <div class="container">
    <h1 class="title is-1 mb-0 pt-5"><router-link style="color: #2c3e50" :to="'/'">Zwitscher</router-link></h1>
    <div id="nav" class="mt-0">
      <router-link :to="'/timeline/'+$route.params.id">Timeline</router-link> |
      <router-link :to="'/tweets/'+$route.params.id">Tweets</router-link> |
      <router-link :to="'/connections/'+$route.params.id">Connections</router-link> |
      <router-link :to="'/new/'+$route.params.id">New Tweet</router-link>
    </div>

    <h2 class="subtitle is-2">Timeline for @{{ $route.params.id }}</h2>
    
    <Tweet v-for="tweet in tweets" :key="tweet.id" :tweet="tweet" />

  </div>
</template>

<script>
import axios from 'axios';
import Tweet from '@/components/Tweet.vue';

export default {
  name: 'Timeline',
  components: { Tweet },
  data: function() {
    return {
      tweets: []
    };
  },
  props: {
    id: String
  },
  mounted() {
    axios
      .get("http://192.168.0.232:7080/dash?user_id="+this.$route.params.id)
      .then(response => this.tweets = response.data.content)
      .catch(error => console.log(error));
  }
}
</script>
