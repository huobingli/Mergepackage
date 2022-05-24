import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';

// 测试工具
// import Tool from './views/Tool.vue';

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
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // component: () => import( /* webpackChunkName: "Home" */ './views/Home.vue')
    }
  ]
});