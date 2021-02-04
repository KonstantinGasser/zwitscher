<template>
  <div id="tweetlist" class="container">
    <Tweet v-for="tweet in tweets" :key="tweet.id" :tweet="tweet"></Tweet>
  <p>{{id}}</p>
  </div>
</template>

<script>
import Tweet from './Tweet.vue';
import axios from 'axios';

export default {
  components: { Tweet },
  name: "Dashboard",
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
      .get("http://192.168.0.232:7080/dash?user_id="+this.id)
      .then(response => this.tweets = response.data.content)
      .catch(error => {console.log(error)});
  }
};

</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
