<template>
  <div class="container">
    <h1>AdA hinzufügen</h1>
    <form @submit.prevent="addAdA" class="center-form">
      <!-- Form content -->
      <div class="mb-3">
        <label for="name" class="form-label">Name:</label>
        <input type="text" class="form-control" v-model="newAdA.name" required>
      </div>
      <div class="mb-3">
        <label for="rank" class="form-label">Rang:</label>
        <select v-model="newAdA.rank" class="form-select" required>
          <option v-for="rank in rankOptions" :key="rank.rank" :value="rank.rank">{{ rank.rank }}</option>
        </select>
      </div>
      <div class="mb-3">
        <label for="company" class="form-label">Kompaniename:</label>
        <select v-model="newAdA.company" class="form-select" required>
          <option v-for="company in companyOpts" :key="company.name" :value="company.name">{{company.name}}</option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary">Hinzufügen</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      newAdA: { name: '', rank: '', company: '' },
      rankOptions: [],
      companyOpts: []
    };
  },
  created() {
    this.fetchRankOptions();
    this.fetchCompanyOpts();
  },
  methods: {
    async fetchRankOptions() {
      try {
        const response = await axios.get('http://localhost:8080/api/ako/rank/all');
        this.rankOptions = response.data;
        this.newAdA.rank = this.rankOptions[0] ? this.rankOptions[0].rank : '';
      } catch (error) {
        console.error('Error fetching rank options:', error);
      }
    },
    async fetchCompanyOpts() {
      try {
        const response = await axios.get('http://localhost:8080/api/ako/company/all');
        this.companyOpts = response.data;
        this.newAdA.company = this.companyOpts[0] ? this.companyOpts[0].name : '';
      } catch (error) {
        console.error('Error fetching rank options:', error);
      }
    },
    async addAdA() {
      try {
        const payload = {
          Name: this.newAdA.name,
          Rank: this.newAdA.rank,
          CompanyName: this.newAdA.company
        };

        const response = await axios.post('http://localhost:8080/api/ako/ada/', payload);

        console.log('AdA added successfully:', response.data);

        this.newAdA = {
          name: '',
          rank: this.rankOptions[0] ? this.rankOptions[0].rank : '',
          company: this.companyOpts[0] ? this.companyOpts[0].name : ''
        };
      } catch (error) {
        console.error('Error adding AdA:', error);
      }
    }
  }
};
</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.center-form {
  width: 300px;
  margin: auto;
}
</style>
