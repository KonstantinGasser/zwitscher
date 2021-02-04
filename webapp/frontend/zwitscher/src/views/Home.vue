<template>
  <div class="container">
    <h1 class="title is-1 pt-5">Zwitscher</h1>
    <!-- <div id="nav" class="mt-0">
      <router-link :to="'/timeline/'+$route.params.id">Timeline</router-link> |
      <router-link :to="'/tweets/'+$route.params.id">Tweets</router-link> |
      <router-link :to="'/connections/'+$route.params.id">Connections</router-link> |
      <router-link :to="'/new/'+$route.params.id">New Tweet</router-link>
    </div> -->

    <h2 class="subtitle is-2">These users might be interesting:</h2>
    
    <p style="display: inline-block" v-for="id in users" :key="id"><router-link class="is-size-3 mx-1" :to="'/tweets/'+id">@{{ id }}</router-link></p>

  </div>
</template>


<script>

export default {
  name: 'Home',
  components: {
  },
  data() {
    return {
      users: [
        2705,
        6463,
        7846,
        10350,
        11627,
        12925,
        15913,
        17743,
        21413,
        55113,
        60273,
        465973,
        657863,
        678953,
        756278,
        756379,
        767809,
        769101,
        769584,
        783214,
        792998,
        794186,
        795536,
        813286,
        814153,
        820485,
        865561,
        899701,
        972651,
        991221,
        1002351,
        1034621,
        1043241,
        1044591,
        1044771,
        1045141,
        1047451,
        1047961,
        1048661,
        1049111,
        1049981,
        1065921,
        1116721,
        1117341,
        1118161,
        1118781,
        1159201,
        1183041,
        1268821,
        1269251,
        1270191,
        1294741,
        1385211,
        1389951,
        1403031,
        1435461,
        1452261,
        1468951,
        1476651,
        1478151,
        1478471,
        1534321,
        1712051,
        1715051,
        1879831,
        1975721,
        2014851,
        2154171,
        2217251,
        2260771,
        2271071,
        2367911,
        2681811,
        2752851,
        2811851,
        2883841,
        2904691,
        2916671,
        3026061,
        3037321,
        3096071,
        3107971,
        3359851,
        3462421,
        3527151,
        3829151,
        3956091,
        4126501,
        4208931,
        4219131,
        4379101,
        4462381,
        4556941,
        4620451,
        4700301,
        4841921,
        4893411,
        4952751,
        5066391,
        5190551
      ]
    }
  }
}
</script>
