import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';
import Overview from './Overview.vue';
import AddAdA from './AddAdA.vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Overview },
        { path: '/add-ada', component: AddAdA }
    ]
});

const app = createApp(App);
app.use(router);
app.mount('#app');