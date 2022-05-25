// console.log(process.env);
// import devtools from '@vue/devtools';

// console.log(process.env);
// if (process.env.NODE_ENV === 'development' && process.env.VUE_APP_USEDEVTOOLS === 'true') {
//   const devtools = require('@vue/devtools');
//   devtools.connect(
//     'http://localhost',
//     '8098'
//   );
// }
import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import ElementUI from 'element-ui';
import Element from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);
Vue.use(Element, { size: 'small', zIndex: 3000 });


Vue.config.productionTip = false;
Vue.prototype.$eventBus = new Vue();

const app = new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
