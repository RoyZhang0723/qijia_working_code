<template>
  <el-menu :default-active="1" router class="el-menu-demo" mode="horizontal" v-on:select="handleSelect" style="background-color: #545c64;text-color: #fff;active-text-color:#ffd04b">
    <!--    得开启router模式，这样才能根据index来寻找路径-->
    <el-menu-item index="/index1" key="我的课程" class="my-course">My course</el-menu-item>
    <el-menu-item index="/register_course" key="注册课程" class="register-course">Register Course</el-menu-item>
    <span style="position: absolute;padding-top: 20px;right: 43%;font-size: 20px;font-weight: bold;color: aliceblue">EasyChoose -- your student information manager</span>
  </el-menu>
</template>

<script>
export default {
  name: 'NavMenu',
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>

<style scoped>
</style>
