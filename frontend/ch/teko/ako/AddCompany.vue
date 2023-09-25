<template>
  <div class="container">
    <h1>Kompanie hinzufügen</h1>
    <form @submit.prevent="addCompany" class="center-form">
      <div class="mb-3">
        <label for="company" class="form-label">Kompaniename:</label>
        <input type="text" class="form-control" v-model="newCompany.name" required>
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
      newCompany: { name: '' }
    };
  },
  methods: {
    async addCompany() {
      try {
        // Make an API call to add the company
        await axios.post('http://localhost:8080//api/ako/company/', {
          name: this.newCompany.name
        });

        // Reset the form
        this.newCompany.name = '';

        console.log('Company added successfully.');
      } catch (error) {
        console.error('Error adding company:', error);
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
  height: 100vh; /* Adjust as needed to center vertically */
}

.center-form {
  width: 300px; /* Adjust the width of the form as needed */
  margin: auto;
}
</style>
