<template>
  <div class="container">
    <h1>Lektion hinzufügen</h1>
    <form @submit.prevent="addLesson" class="center-form">
      <div class="mb-3">
        <label for="name" class="form-label">Name:</label>
        <input type="text" class="form-control" v-model="newLesson.name" required>
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
      newLesson: { name: '' }
    };
  },
  methods: {
    async addLesson() {
      try {
        // Make an API call to add the lesson
        await axios.post('http://localhost:8080/api/ako/lesson/', {
          name: this.newLesson.name
        });

        // Reset the form
        this.newLesson.name = '';

        console.log('Lesson added successfully.');
      } catch (error) {
        console.error('Error adding lesson:', error);
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
