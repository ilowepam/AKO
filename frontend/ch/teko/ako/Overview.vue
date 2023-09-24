<template>
  <div>
    <h1>Übersicht</h1>
    <table class="table">
      <thead>
      <tr>
        <th>Name</th>
        <th>Rang</th>
        <th>Kompaniename</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in data" :key="index">
        <td>{{ item.name }}</td>
        <td>{{ item.rank }}</td>
        <td>{{ item.company }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      data: [
        { name: 'John Doe', rank: 'General', company: 'Alpha' },
        { name: 'Jane Smith', rank: 'Major', company: 'Bravo' },
        // ... add more data as needed
      ]
    };
  }
};
</script>
