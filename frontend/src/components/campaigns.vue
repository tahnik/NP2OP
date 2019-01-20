<template>
  <div class="campaigns">
    <v-layout class="mt-3" row wrap>
      <v-flex
        lg3 sm4 md4 v-for="campaign in campaigns"
        :key="campaign.Description"
        class="pa-2">
        <v-card class="mb-3">
          <v-card-title primary-title>
            <div class="mb-2">
              <label>
                Posted by {{ campaign.Name }}
              </label>
            </div>
            <label class="title font-weight-medium">
              {{ campaign.Title }}
            </label>
          </v-card-title>
          <img src="../assets/campaign.jpg" style="width: 100%" />
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
            <v-btn flat color="#1a1f71">View</v-btn>
            <v-btn flat color="#1a1f71">Donate</v-btn>
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
      campaigns: [
        {
          Name: 'Samith Silva',
          Email: 'sam@gmail.com',
          Phone: '07850318114',
          Country: 'UK',
          Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
          UserType: 'individual',
          Title: 'Funding for Samith Silva to attend Computer Science course',
          CourseName: 'Payments Everywhere',
          CourseDescription: 'Course to understand more the visa business',
          Cost: 500,
          Length: '5 months',
          Requirements: 'Being a Visa employee',
          CourseEmail: 'paymentseverywhere@visa.com',
          SchoolName: 'Visa University',
          TotalReceived: 200,
        },
        {
          Name: 'Samith Silva',
          Email: 'sam@gmail.com',
          Phone: '07850318114',
          Country: 'UK',
          Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
          UserType: 'individual',
          Title: 'Funding for Samith Silva to attend Computer Science course',
          CourseName: 'Payments Everywhere',
          CourseDescription: 'Course to understand more the visa business',
          Cost: 500,
          Length: '5 months',
          Requirements: 'Being a Visa employee',
          CourseEmail: 'paymentseverywhere@visa.com',
          SchoolName: 'Visa University',
          TotalReceived: 400,
        },
        {
          Name: 'Samith Silva',
          Email: 'sam@gmail.com',
          Phone: '07850318114',
          Country: 'UK',
          Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
          UserType: 'individual',
          Title: 'Funding for Samith Silva to attend Computer Science course',
          CourseName: 'Payments Everywhere',
          CourseDescription: 'Course to understand more the visa business',
          Cost: 500,
          Length: '5 months',
          Requirements: 'Being a Visa employee',
          CourseEmail: 'paymentseverywhere@visa.com',
          SchoolName: 'Vis2a University',
          TotalReceived: 150,
        },
        {
          Name: 'Samith Silva',
          Email: 'sam@gmail.com',
          Phone: '07850318114',
          Country: 'UK',
          Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
          UserType: 'individual',
          Title: 'Funding for Samith Silva to attend Computer Science course',
          CourseName: 'Payments Everywhere',
          CourseDescription: 'Course to understand more the visa business',
          Cost: 1000,
          Length: '5 months',
          Requirements: 'Being a Visa employee',
          CourseEmail: 'paymentseverywhere@visa.com',
          SchoolName: 'Vis2a University',
          TotalReceived: 900,
        },
        {
          Name: 'Samith Silva',
          Email: 'sam@gmail.com',
          Phone: '07850318114',
          Country: 'UK',
          Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
          UserType: 'individual',
          Title: 'Funding for Samith Silva to attend Computer Science course',
          CourseName: 'Payments Everywhere',
          CourseDescription: 'Course to understand more the visa business',
          Cost: 700,
          Length: '5 months',
          Requirements: 'Being a Visa employee',
          CourseEmail: 'paymentseverywhere@visa.com',
          SchoolName: 'Vis2a University',
          TotalReceived: 150,
        },
      ]
    }
  },
  methods: {
    getCost(campaign) {
      const { TotalReceived, Cost } = campaign;

      return TotalReceived/Cost*100;
    }
  }
}
</script>


<style lang="scss">
.campaigns {
  height: 100%;
  width: 100%;
}
</style>

