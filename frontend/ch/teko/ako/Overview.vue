<template>
  <div>
    <h1>Overview</h1>
    <table class="table">
      <thead>
      <tr>
        <th v-for="(header, index) in headers" :key="index">{{ header }}</th>
        <th>Status</th>
      </tr>
      </thead>
      <tbody>
      <tr v-if="rows.length > 0" v-for="(row, rowIndex) in rows" :key="rowIndex">
        <td>{{ row.ada.id }}</td>
        <td>{{ row.ada.name }}</td>
        <td>{{ row.ada.rank }}</td>
        <td>{{ row.ada.companyName }}</td>
        <td v-for="(lesson, lessonIndex) in row.lesson" :key="lessonIndex">
          <span
              v-if="lesson.status"
              class="green-checkmark"
          >&#10004;</span>
          <span
              v-else
              class="red-x"
              @click="setLessonStatus(row.ada.id, lesson.name)"
          >&#10008;</span>
        </td>
        <td>{{ getGreenPercentage(row.lesson) }}%</td>
      </tr>
      <tr v-else>
        <td :colspan="headers.length">No data available</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      headers: [],
      rows: []
    };
  },
  methods: {
    setLessonStatus(pAdaId, pLessonName) {
      axios.post('http://localhost:8080/api/ako/lesson/status/', {
        adaId: pAdaId,
        lessonName: pLessonName
      });
      window.location.reload();
    },
    getGreenPercentage(lessons) {
      const totalLessons = lessons.length;
      if (totalLessons === 0) return 0;

      const greenLessons = lessons.filter(lesson => lesson.status).length;
      return (greenLessons / totalLessons) * 100;
    },
    fetchData() {
      axios
          .get('http://localhost:8080/api/ako/overview/')
          .then(response => {
            this.headers = response.data.headers;
            this.rows = response.data.rows;
          })
          .catch(error => {
            console.error('Error fetching data:', error);
          });
    }
  },
  mounted() {
    this.fetchData();
  }
};
</script>

<style scoped>
/* Add your custom styles for the table if needed */
.table {
  width: 100%;
  border-collapse: collapse;
}

.table th,
.table td {
  border: 1px solid #ccc;
  padding: 8px;
  text-align: left;
}

.table th {
  background-color: #f0f0f0;
}

.green-checkmark {
  color: green;
  margin-right: 5px;
  cursor: pointer;
}

.red-x {
  color: red;
  margin-right: 5px;
  cursor: pointer;
}
</style>
