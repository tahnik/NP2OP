import Vuex from 'vuex';
import Vue from 'vue';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
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
  },
  mutations: {
    increment (state) {
      state.count++
    }
  }
})

export default store;
