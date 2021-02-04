<template>
  <div class="container">
    <h1 class="title is-1 mb-0 pt-5"><router-link style="color: #2c3e50" :to="'/'">Zwitscher</router-link></h1>
    <div id="nav" class="mt-0">
      <router-link :to="'/timeline/'+$route.params.id">Timeline</router-link> |
      <router-link :to="'/tweets/'+$route.params.id">Tweets</router-link> |
      <router-link :to="'/connections/'+$route.params.id">Connections</router-link> |
      <router-link :to="'/new/'+$route.params.id">New Tweet</router-link>
    </div>

    <h2 class="subtitle is-2">Connections of @{{ $route.params.id }}</h2>

    <div class="columns">
      <div class="column">
        <h4 class="subtitle is-3">Following</h4>
        <h4 class="subtitle is-3" v-for="id in following" :key="'x'+id"><router-link :to="'/tweets/'+tweet.user_id">@{{ id }}</router-link></h4>
      </div>
      <div class="column">
        <h4 class="subtitle is-3">Followers</h4>
        <h4 class="subtitle is-3" v-for="id in followers" :key="'y'+id"><router-link :to="'/tweets/'+tweet.user_id">@{{ id }}</router-link></h4>
      </div>
    </div> 
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Connections',
  components: {
  },
  data() {
    return {
      following: [],
      followers: [],
    }
  },
  mounted() {
    axios
      .get("http:192.168.0.232:7080/getfollow?user_id="+this.$route.params.id)
      .then(response => {
        this.following = response.data.following;
        this.followers = response.data.followers;
      });
  }
}
</script>