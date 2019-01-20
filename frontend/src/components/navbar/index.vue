<template>
  <div class="navbar">
    <v-navigation-drawer
      fixed
      v-if="$vuetify.breakpoint.mdAndDown"
      app
      v-model="drawer"
    >
      <v-list dense>
        <template v-for="item in items">
          <v-layout
            row
            v-if="item.heading"
            align-center
            :key="item.heading"
          >
            <v-flex xs6>
              <v-subheader v-if="item.heading">
                {{ item.heading }}
              </v-subheader>
            </v-flex>
            <v-flex xs6 class="text-xs-center">
              <a href="#!" class="body-2 black--text">EDIT</a>
            </v-flex>
          </v-layout>
          <v-list-group
            v-else-if="item.children"
            v-model="item.model"
            :key="item.text"
            :prepend-icon="item.model ? item.icon : item['icon-alt']"
            append-icon=""
          >
            <v-list-tile slot="activator">
              <v-list-tile-content>
                <v-list-tile-title>
                  {{ item.text }}
                </v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile
              v-for="(child, i) in item.children"
              :key="i"
              @click="goTo(child)"
            >
              <v-list-tile-action v-if="child.icon">
                <v-icon>{{ child.icon }}</v-icon>
              </v-list-tile-action>
              <v-list-tile-content>
                <v-list-tile-title>
                  {{ child.text }}
                </v-list-tile-title>
              </v-list-tile-content>
            </v-list-tile>
          </v-list-group>
          <v-list-tile v-else @click="goTo(item)" :key="item.text">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>
                {{ item.text }}
              </v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar
      color="#1a1f71"
      dark
      app
      :clipped-left="$vuetify.breakpoint.mdAndUp"
      fixed
    >
      <v-toolbar-title
        style="width: 300px"
        :class="`ml-0 ${$vuetify.breakpoint.mdAndDown ? 'pl-0' : 'pl-3'}`">
        <v-toolbar-side-icon
          @click.stop="drawer = !drawer"
          v-if="$vuetify.breakpoint.mdAndDown"
        ></v-toolbar-side-icon>
        <span style="cursor: pointer" @click="$router.push('/')">NP2OP</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-layout justify-end style="height: 100%" v-if="$vuetify.breakpoint.mdAndUp">
        <div
          v-for="item in items"
          :key="item.text"
          style="height: 100%; position: relative"
        >
          <v-btn
            flat style="height: 100%"
            class="my-0 mx-0"
            @click="goTo(item)"
          >
            {{ item.text }}
          </v-btn>
          <transition name="slide">
            <div
              v-if="activeItem && activeItem.text === item.text"
              class="active-btn"
            ></div>
          </transition>
        </div>
      </v-layout>
    </v-toolbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      drawer: null,
      activeItem: null,
      items: [
        { text: 'Account', icon: 'mdi-account-plus', route: '/account' },
      ]
    }
  },
  mounted() {
    console.log('nop');
  },
  methods: {
    goTo(item) {
      this.activeItem = item;
      this.$router.push(item.route);
    }
  }
}
</script>

<style lang="scss">
.navbar {
  .v-toolbar__content {
    padding-left: 0;
    padding-right: 0;
  }
  .v-btn {
    border-radius: 0;
  }
  .active-btn {
    background: #fdbb0a;
    width: 100%;
    position: absolute;
    bottom: 0;
    right: 0;
    left: 0;
    height: 3px;
  }
  .slide-enter-active, .slide-leave-active {
    transition: opacity .5s;
  }
  .slide-enter, .slide-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
  }
}
</style>

