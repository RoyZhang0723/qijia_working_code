import Vue from 'vue'
import Router from 'vue-router'
import Login from '../components/Login'
import AppIndex from '../components/home/AppIndex'
import StuInf from '../components/home/StuInf'
import CompCourse from '../components/home/Course/CompCourse'
import AMACourse from '../components/home/Course/AMACourse'
import AddStudent from '../components/home/Edit/AddStudent'
import DeleteStudent from '../components/home/Edit/DeleteStudent'
import ABCTCourse from '../components/home/Course/ABCTCourse'
import AppIndex1 from '../components/home/AppIndex1'
import Home from '../components/home/Home'
import Home1 from '../components/home/Home1'
import RegisterCourse from '../components/home/RegisterCourse'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/home',
      name: 'Home',
      component: Home,
      redirect: '/index',
      children: [
        {
          path: '/index',
          name: 'AppIndex',
          component: AppIndex,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/stu_inf',
          name: 'StuInf',
          component: StuInf,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/COMP',
          name: 'COMCourse',
          component: CompCourse,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/AMA',
          name: 'AMACourse',
          component: AMACourse,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/ABCT',
          name: 'ABCTCourse',
          component: ABCTCourse,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/a_student',
          name: 'AddStudent',
          component: AddStudent,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/d_student',
          name: 'DeleteStudent',
          component: DeleteStudent,
          meta: {
            requireAuth: true
          }
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/home1',
      name: 'Home1',
      component: Home1,
      redirect: '/index1',
      children: [
        {
          path: '/index1',
          name: 'AppIndex1',
          component: AppIndex1,
          meta: {
            requireAuth: true
          }
        },
        {
          path: '/register_course',
          name: 'RegisterCourse',
          component: RegisterCourse,
          meta: {
            requireAuth: true
          }
        }
      ]
    }
  ]
})
