<template>
  <div class="card m-5 is-centered" style="border-radius: 1rem">
    <div class="card-header p-2 pl-4">
      <p class="is-size-5 has-text-left" style="overflow-wrap: normal"><strong>{{ tweet.author }} <router-link class="ml-1 mr-2" :to="'/tweets/'+tweet.user_id">@{{ tweet.user_id }}</router-link></strong> {{tweet.content}}</p>
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