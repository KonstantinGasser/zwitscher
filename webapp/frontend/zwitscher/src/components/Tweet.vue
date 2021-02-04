<template>
  <div class="card m-5">
    <div>
      <h4 class="card-header-title is-3">{{ tweet.author }} <router-link class="ml-1" :to="'/tweets/'+tweet.user_id">@{{ tweet.user_id }}</router-link></h4>
    </div>
    <div class="card-content">
      <p class="is-5">{{tweet.content}}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "Tweet",
  props: {
      tweet: Object,
  }
};
</script>