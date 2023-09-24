<template>
  <div>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <router-link to="/" class="navbar-brand">Übersicht</router-link>
      <router-link to="/add-ada" class="navbar-brand">AdA hinzufügen</router-link>
    </nav>

    <router-view></router-view>
  </div>
</template>

<script>
export default {
};
</script>

<style>
/* Add any custom styles if needed */
</style>