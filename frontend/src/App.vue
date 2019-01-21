<template>
  <div id="app">
    <v-app>
      <nav-bar></nav-bar>
      <v-content>
        <v-container fill-height fluid>
          <router-view></router-view>
        </v-container>
      </v-content>
    </v-app>
  </div>
</template>

<script>
import NavBar from './components/navbar/index.vue';

export default {
  components: {
    NavBar,
  },
  name: 'app',
}
</script>

<style lang="scss">

.theme--light.application {
  background: linear-gradient(45deg,#0c4da9,#2894d2 45%);
}

h1, h2, h3, h4, h5, h6 {
  font-family: 'Titillium Web', sans-serif;
}

::-webkit-scrollbar {
  width: 5px;
}

::-webkit-scrollbar-thumb {
  background: #666;
}

::-webkit-scrollbar-track {
  background: #ddd;
}

.v-toolbar {
  box-shadow: none;
}

.v-card {
  border-radius: 5px;
}

</style>
