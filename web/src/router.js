import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/home.vue';
import Publish from './views/publish.vue';
import Package from './views/package.vue';
import testPackage from './views/testPackage.vue';
import exePackage from './views/exePackage.vue';
import signPackage from './views/signPackage.vue';

Vue.use(Router);

export default new Router({
  routes: [{
      path: '/',
      redirect: '/testPackage'
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
    },
    {
      path: '/testPackage',
      name: 'testPackage',
      component: testPackage
    },
    {
      path: '/signPackage',
      name: 'signPackage',
      component: signPackage
    },
    {
      path: '/exePackage',
      name: 'exePackage',
      component: exePackage
    }
  ]
});