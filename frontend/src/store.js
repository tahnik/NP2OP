import Vuex from 'vuex';
import Vue from 'vue';
import createPersistedState from 'vuex-persistedstate';

Vue.use(Vuex);

const store = new Vuex.Store({
  plugins: [createPersistedState()],
  state: {
    campaigns: [
      {
        Name: 'Hermann Gmeiner',
        Email: 'sam@gmail.com',
        Phone: '07850318114',
        Country: 'UK',
        Image: 'https://i.imgur.com/H9dvBUW.jpg',
        Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
        UserType: 'organisation',
        Title: 'Funding for Hermann Gmeiner',
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
        Name: 'Mikey Wotton',
        Email: 'mikey@gmail.com',
        Phone: '07850318114',
        Country: 'UK',
        Image: 'https://i.imgur.com/fMGsksL.jpg',
        Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
        UserType: 'individual',
        Title: 'Funding for Mikey Wotton',
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
        Name: 'Zhang Jing',
        Email: 'tahnik@gmail.com',
        Phone: '07850318114',
        Country: 'UK',
        Image: 'https://i.imgur.com/TpOpw31.jpg',
        Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
        UserType: 'individual',
        Title: 'Funding for Zhang Jing',
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
        Name: 'Monipur High School',
        Email: 'peter@gmail.com',
        Phone: '07850318114',
        Country: 'UK',
        Image: 'https://i.imgur.com/uAHMYji.jpg',
        Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
        UserType: 'organisation',
        Title: 'Funding for Monipur High School',
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
        Name: 'Residential School',
        Email: 'james@gmail.com',
        Phone: '07850318114',
        Country: 'UK',
        Image: 'https://i.imgur.com/L9o4ot0.jpg',
        Description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`,
        UserType: 'organisation',
        Title: 'Funding for Residential School',
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
  },
  mutations: {
    addFunding (state, nextState) {
      const c = state.campaigns.filter((campaign) => campaign.Email === nextState.email)[0];
      c.TotalReceived += nextState.amount;
    }
  }
})

export default store;
