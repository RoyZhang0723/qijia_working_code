<template>
  <body id="poster">
    <el-form class="login-container" label-position="left" label-width="0px">
      <h3 class="login-title">选课系统</h3>
      <el-form-item>
        <el-input type="text" v-model="LoginForm.username" auto-complete="off" placeholder="学生证号"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input type="password" v-model="LoginForm.password" auto-complete="off" placeholder="密码"></el-input>
      </el-form-item>
      <el-form-item style="width: 100%">
        <el-button type="primary" style="width: 100%;background: dodgerblue;border: none" v-on:click="login">管理员登录</el-button>
      </el-form-item>
      <el-form-item style="width: 100%">
        <el-button type="primary" style="width: 100%;background: dodgerblue;border: none" v-on:click="login1">学生登录</el-button>
      </el-form-item>
    </el-form>
  </body>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      LoginForm: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    login () {
      this.$axios.post('/login1', {
        username: this.LoginForm.username,
        password: this.LoginForm.password
      }).then(successResponse => {
        if (successResponse.data.code === 200) {
          this.$store.commit('login', this.LoginForm)
          this.$router.replace({path: '/index'})
          // var path = this.$route.query.redirect
          // this.$router.replace({path: path === '/' || path === undefined ? '/index' : path})
        }
      }).catch(fail => {})
    },
    login1 () {
      this.$axios.post('/login', {
        username: this.LoginForm.username,
        password: this.LoginForm.password
      }).then(successResponse => {
        if (successResponse.data.code === 200) {
          this.$store.commit('login', this.LoginForm)
          this.$router.replace({path: '/index1'})
          // var path = this.$route.query.redirect
          // this.$router.replace({path: path === '/' || path === undefined ? '/index' : path})
        }
      }).catch(fail => {})
    }
  }
}
</script>

<style>
  #poster {
    background: url("../assets/eva.jpg") no-repeat;
    background-position: center;
    height: 100%;
    width: 100%;
    background-size: cover;
    position: fixed;
  }
  body{
    margin: 0px;
  }
  .login-container{
    border-radius: 15px;
    background-clip: padding-box;
    margin: 90px auto;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
  }
  .login-title {
    margin: 0px auto 40px auto;
    text-align: center;
    color: #505458;
  }
</style>
