import { createApp } from 'vue';
import Navigation from './Navigation.vue';
import { createRouter, createWebHistory } from 'vue-router';


import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import Overview from './Overview.vue';
import AddCompany from "@/AddCompany.vue";
import AddAdA from './AddAdA.vue';
import AddLesson from "@/AddLesson.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Overview },
        { path: '/add-ada', component: AddAdA },
        { path: '/add-company', component: AddCompany},
        { path: '/add-lesson', component: AddLesson}
    ]
});

const app = createApp(Navigation);
app.use(router);
app.mount('#app');