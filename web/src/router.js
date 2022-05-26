import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/home.vue';
import Publish from './views/publish.vue';
import Package from './views/package.vue';

Vue.use(Router);

export default new Router({
  routes: [{
      path: '/',
      redirect: '/index'
    },
    {
      path: '/index',
      name: 'home',
      component: Home
    },
    {
      path: '/package',
      name: 'package',
      component: Package
    },
    {
      path: '/publish',
      name: 'publish',
      component: Publish
    }
  ]
});