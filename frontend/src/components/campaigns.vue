<template>
  <div class="campaigns">
    <v-layout class="mt-3" row wrap>
      <v-flex
        lg3 sm4 md4 v-for="(campaign, index) in campaigns"
        :key="index"
        class="pa-2">
        <v-card class="mb-3">
          <v-card-title primary-title>
            <v-layout class="mb-2" column>
              <label>
                Posted by {{ campaign.Name }}
              </label>
              <div class="mb-2"></div>
              <label class="title font-weight-medium">
                {{ campaign.Title }}
              </label>
            </v-layout>
          </v-card-title>
          <img :src="campaign.Image" style="width: 100%" />
          <v-card-text>
            <p style="text-align: justify; margin-bottom: 0">
              {{ `${campaign.Description.substring(0,150)}...` }}
            </p>
            <v-layout align-center class="mb-2 mt-3">
              <v-progress-linear
                :value="getCost(campaign)"
                class="mr-3 mb-0 mt-0"
                :height="40"
                style="border-radius: 5px"
              ></v-progress-linear>
              <v-layout column align-center class="pt-1">
                <label class="title font-weight-light text-xs-center">
                  £{{ campaign.TotalReceived }}
                </label>
                <label class="caption text-xs-center">raised</label>
              </v-layout>
            </v-layout>
          </v-card-text>
          <v-divider>

          </v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn flat color="#1a1f71" @click="viewCampaign(campaign)">View</v-btn>
            <v-btn flat color="#1a1f71" @click="donateToCampaign(campaign)">Donate</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
    <v-dialog v-model="dialog" persistent scrollable max-width="1000">
      <v-card v-if="campaign">
        <v-layout style="height: 50px">
          <v-progress-linear
            style="position: absolute"
            :value="getCost(campaign)" height="50" class="my-0"
          >
          </v-progress-linear>
          <v-layout style="z-index: 99" align-center class="px-3">
            <label class="title white--text font-weight-light">
              Funded: £{{ campaign.TotalReceived }}
            </label>
            <v-spacer></v-spacer>
            <label
              :class="`title font-weight-light ${getCost(campaign) > 90 ? 'white--text' : ''}`"
            >
              £{{ campaign.Cost }}
            </label>
          </v-layout>
        </v-layout>
        <v-card-text style="max-height: 500px" class="pa-0">
          <v-layout>
            <v-flex sm8>
              <v-layout class="pa-3" column>
                <v-layout class="mb-3 mt-1" column>
                  <label class="headline font-weight-regular">
                    {{ campaign.Title }}
                  </label>
                  <div class="mt-2"></div>
                  <label>Posted By {{ campaign.Name }}</label>
                </v-layout>
                <img :src="campaign.Image" style="width: 100%" class="mb-3" />
                <label>{{ campaign.Description }}</label>
              </v-layout>
            </v-flex>
            <v-divider vertical></v-divider>
            <v-flex sm4>
              <v-layout class="pa-3" column>
                <v-layout class="mb-3 mt-1">
                  <label class="headline font-weight-regular">
                    {{ campaign.UserType === 'individual' ? 'Course' : 'Organisation' }} Details
                  </label>
                </v-layout>
                <v-layout column>
                  <v-subheader>Name</v-subheader>
                  <label class="subheading">{{ campaign.CourseName }}</label>
                  <v-subheader class="mt-3">Organisation/School</v-subheader>
                  <label class="subheading">{{ campaign.SchoolName }}</label>
                  <v-subheader class="mt-3">Course Email</v-subheader>
                  <label class="subheading">{{ campaign.CourseEmail }}</label>
                  <v-subheader class="mt-3">Course Length</v-subheader>
                  <label class="subheading">{{ campaign.Length }}</label>
                  <v-subheader class="mt-3">Course Requirement</v-subheader>
                  <label class="subheading">{{ campaign.Requirements }}</label>
                  <v-subheader class="mt-3">Course Description</v-subheader>
                  <label class="subheading">{{ campaign.CourseDescription }}</label>
                </v-layout>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="#1a1f71" flat @click="dialog = false">Close</v-btn>
          <v-btn color="#1a1f71" flat @click="donateToCampaign(campaign)">
            Donate
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="donateD" persistent max-width="500">
      <v-card v-if="donateCampaign">
        <transition name="payfade">
          <v-layout class="payLoad" justify-center align-center v-if="donateLoading">
            <v-layout
              v-if="!donateDone"
              column justify-center align-center
            >
              <v-progress-circular
                :size="50"
                color="#1a1f71"
                indeterminate
              ></v-progress-circular>
              <div class="mt-5"></div>
              <label class="title font-weight-light">Processing donation</label>
            </v-layout>
            <v-layout v-else column justify-center align-center>
              <svg class="checkmark" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 52 52"><circle class="checkmark__circle" cx="26" cy="26" r="25" fill="none"/><path class="checkmark__check" fill="none" d="M14.1 27.2l7.1 7.2 16.7-16.8"/></svg>
              <div class="mt-2"></div>
              <label class="title font-weight-light">Donation Complete! Thank you</label>
            </v-layout>
          </v-layout>
        </transition>
        <v-layout class="pa-3">
          <label class="headline font-weight-light">
            Donate to {{ donateCampaign.Name }}
          </label>
        </v-layout>
        <v-card-text>
          <v-subheader>How much do you want to donate</v-subheader>
          <v-slider
            class="mt-5 mx-2"
            v-model="donateAmount"
            thumb-label
            :max="donateCampaign.Cost - donateCampaign.TotalReceived"
          ></v-slider>
          <v-layout justify-center class="mb-3">
            <label class="headline">£{{ donateAmount }}</label>
          </v-layout>
          <v-subheader>Payment Details</v-subheader>
          <v-card class="pa-4 mt-3 pt-5" style="background: linear-gradient(200deg,#0c4da9,#2894d2 45%);">
            <v-text-field
              color="white"
              class="mt-5"
              mask="#### - #### - #### - ####"
              label="XXXX - XXXX - XXXX - XXXX"
              v-model="donateLongNumber"
            ></v-text-field>
            <v-layout justify-start class="credit-card">
              <v-flex sm2>
                <v-text-field
                  style="max-width: 35px"
                  color="white"
                  label="MM"
                  mask="##"
                ></v-text-field>
              </v-flex>
              <v-flex sm2 class="mr-5">
                <v-text-field
                  style="max-width: 35px"
                  color="white"
                  label="YY"
                  mask="##"
                ></v-text-field>
              </v-flex>
              <v-flex sm2 class="ml-4">
                <v-text-field
                  color="white"
                  label="CVV"
                  mask="###"
                ></v-text-field>
              </v-flex>
              <v-flex sm6 justify-end>
                <transition name="fade">
                  <v-layout align-center justify-end style="height: 100%" v-if="visa">
                    <img src="../assets/rsz_1visa.png" height="20px" />
                  </v-layout>
                  <v-layout align-center justify-end style="height: 100%" v-if="mastercard">
                    <img src="../assets/rsz_mastercard-logo-2016-640x455.png" height="20px" />
                  </v-layout>
                </transition>
              </v-flex>
            </v-layout>
          </v-card>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="#1a1f71" flat @click="donateD = false; donateCampaign = null">Close</v-btn>
          <v-btn color="#1a1f71" flat @click="donate()">Donate</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      campaign: null,
      donateD: null,
      donateCampaign: null,
      donateAmount: 0,
      donateLongNumber: null,
      donateLoading: false,
      donateDone: false,
    }
  },
  computed: {
    campaigns() {
      return this.$store.state.campaigns;
    },
    visa() {
      if (this.donateLongNumber && this.donateLongNumber.substring(0,4) === '4658') {
        return true;
      }
      return false;
    },
    mastercard() {
      if (this.donateLongNumber
        && (
          this.donateLongNumber.substring(0,2) === '55'
          || this.donateLongNumber.substring(0,2) === '51'
          || this.donateLongNumber.substring(0,2) === '53'
          || this.donateLongNumber.substring(0,2) === '52'
          || this.donateLongNumber.substring(0,2) === '54'
        )
      ) {
        return true;
      }
      return false;
    },
  },
  methods: {
    getCost(campaign) {
      const { TotalReceived, Cost } = campaign;

      return TotalReceived/Cost*100;
    },
    viewCampaign(campaign) {
      this.campaign = campaign;
      this.dialog = true;
    },
    donateToCampaign(campaign) {
      this.dialog = false;
      this.campaign = null;
      this.donateCampaign = campaign;
      this.donateD = true;
    },
    donate(index) {
      this.donateLoading = true;
      this.donateDone = false;


      setTimeout(() => {
        this.donateDone = true;
      }, 3000);

      setTimeout(() => {
        this.$store.commit('addFunding', {
          email: this.donateCampaign.Email,
          amount: this.donateAmount,
        })
        this.donateLoading = false;
        this.donateD = false;
        this.donateCampaign = null;
        this.donateLongNumber = null;
        this.donateAmount = 0;
      }, 6000);
    }
  }
}
</script>


<style lang="scss">
.v-subheader {
  padding-left: 0;
  height: 30px;
}

.v-progress-linear__bar__determinate {
  border-radius: 5px 0 0 0;
}

.fade-enter-active, .fade-leave-active {
  transition: all .3s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(10%)
}

.payfade-enter-active, .payfade-leave-active {
  transition: all .9s;
}
.payfade-enter, .payfade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.payLoad {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba($color: #FFF, $alpha: 0.8);
  z-index: 99;
}

.checkmark__circle {
  stroke-dasharray: 166;
  stroke-dashoffset: 166;
  stroke-width: 2;
  stroke-miterlimit: 10;
  stroke: #1a1f71;
  fill: none;
  animation: stroke 0.6s cubic-bezier(0.65, 0, 0.45, 1) forwards;
}

.checkmark {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  display: block;
  stroke-width: 2;
  stroke: #fff;
  stroke-miterlimit: 10;
  margin: 5% auto;
  box-shadow: inset 0px 0px 0px #1a1f71;
  animation: fill .4s ease-in-out .4s forwards, scale .3s ease-in-out .9s both;
}

.checkmark__check {
  transform-origin: 50% 50%;
  stroke-dasharray: 48;
  stroke-dashoffset: 48;
  animation: stroke 0.3s cubic-bezier(0.65, 0, 0.45, 1) 0.8s forwards;
}

@keyframes stroke {
  100% {
    stroke-dashoffset: 0;
  }
}
@keyframes scale {
  0%, 100% {
    transform: none;
  }
  50% {
    transform: scale3d(1.1, 1.1, 1);
  }
}
@keyframes fill {
  100% {
    box-shadow: inset 0px 0px 0px 30px #1a1f71;
  }
}


.campaigns {
  height: 100%;
  width: 100%;
}
</style>
