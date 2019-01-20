<template>
  <div class="register">
    <v-layout justify-center align-center style="height: 100%">
      <v-flex md5 lg4 class="px-3">
        <v-card>
          <v-progress-linear
            height="2"
            :indeterminate="loading"
            :background-opacity="loading ? '0.3' : '0'"
            color="#1a1f71"
            class="my-0"
          ></v-progress-linear>
          <v-card-title primary-title>
            <v-layout justify-center>
              <h3 class="headline mb-0 font-weight-light">
                {{ register ? 'Register': 'Sign In' }}
              </h3>
            </v-layout>
          </v-card-title>

          <v-card-text>
            <v-text-field
              label="Email"
              v-model="formData.email"
              :rules="rules.email"
            ></v-text-field>
            <v-expand-transition>
              <div v-if="register">
                <v-text-field
                  label="Name"
                  v-model="formData.name"
                  :rules="rules.name"
                ></v-text-field>
                <v-text-field
                  label="Phone"
                  :rules="rules.phone"
                  v-model="formData.phone"
                ></v-text-field>
                <v-select
                  :items="[
                    'Individual',
                    'Organisation'
                  ]"
                  label="Account Type"
                  :rules="rules.accountType"
                  v-model="formData.accountType"
                ></v-select>
              </div>
            </v-expand-transition>
          </v-card-text>

          <v-card-actions class="pb-3">
            <v-btn flat color="#faaa13" @click="register = !register" small>
              {{ register ? 'I have an account' : 'I don\'t have an account' }}
            </v-btn>
            <v-spacer></v-spacer>
            <v-btn
              flat color="#faaa13"
              @click="submit()">{{ register ? 'Register' : 'Sign in' }}</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </div>
</template>

<script>
export default {
  data() {
    return {
      register: false,
      loading: false,
      formData: {
        name: null,
        email: null,
        phone: null,
        accountType: null,
      },
      rules: {
        email: [
          v => !!v || 'Please enter email',
        ],
        name: [
          v => !!v || 'Please enter name',
        ],
        phone: [
          v => !!v || 'Please enter phone',
        ],
        accountType: [
          v => !!v || 'Please select an account type',
        ],
      }
    }
  },
  methods: {
    submit() {
      if (this.register) {
        return this.registerUser();
      }
      return this.signIn();
    },
    registerUser() {
      this.loading = true;

      setTimeout(() => {
        this.loading = false;
      }, 2000);
    },
    signIn() {
      this.loading = true;

      setTimeout(() => {
        this.loading = false;
      }, 2000);
    },
  }
}
</script>


<style lang="scss">
.register {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: linear-gradient(45deg,#0c4da9,#2894d2 45%);
}
</style>
