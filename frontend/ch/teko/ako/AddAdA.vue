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
        <input type="text" class="form-control" v-model="newAdA.company" required>
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
      rankOptions: [] // Array to store rank options fetched from the API
    };
  },
  created() {
    this.fetchRankOptions(); // Fetch rank options when the component is created
  },
  methods: {
    async fetchRankOptions() {
      try {
        const response = await axios.get('http://localhost:8080/api/ako/rank/all');
        this.rankOptions = response.data;
        this.newAdA.rank = this.rankOptions[0] ? this.rankOptions[0].rank : ''; // Set default rank
      } catch (error) {
        console.error('Error fetching rank options:', error);
      }
    },
    addAdA() {
      // Add logic to handle adding AdA
      console.log('AdA added:', this.newAdA);
      // Reset the form
      this.newAdA = { name: '', rank: this.rankOptions[0] ? this.rankOptions[0].rank : '', company: '' };
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
  height: 100vh; /* Adjust as needed to center vertically */
}

.center-form {
  width: 300px; /* Adjust the width of the form as needed */
  margin: auto;
}
</style>
