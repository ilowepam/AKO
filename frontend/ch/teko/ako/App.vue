<template>
  <div>
    <h1>{{ message }}</h1>
    <button @click="increment">Increment</button>
    <button @click="decrement">Decrement</button>
    <p>Counter: {{ counter }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: 'Simple Counter App',
      counter: 0
    };
  },
  methods: {
    increment() {
      this.counter++;
    },
    decrement() {
      this.counter--;
    }
  }
};
</script>

<style scoped>
/* Add your component-specific styles here */
</style>
