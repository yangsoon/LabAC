import Vue from 'vue'
import VueRouter from 'vue-router'


import Edit from '../components/Edit'
import PaperList from '../components/PaperList'
import Home from '../components/Home'

import store from '../store/main'

Vue.use(VueRouter);


let router = new VueRouter({
  routes:[
    {
      path: '/',
      redirect: 'paper'
    },
    {
      path:'/home',
      name: 'home',
      component: Home
    },
    {
      path:'/edit',
      name: 'edit',
      component: Edit
    },
    {
      path:'/paper',
      name: 'paper',
      component: PaperList
    }
  ]
});

router.beforeEach((to, from, next) => {
  if(to.name === 'edit'){
    if (!localStorage.getItem("labac")){
      next('/');
      return
    }
  }
  store.commit("setActiveMenu", to.name);
  next()
});

// router.afterEach((to, from) => {
//   store.commit("setActiveMenu", to.name);
// });

export default router