<template>
  <el-menu :default-active="1" router class="el-menu-demo" mode="horizontal" v-on:select="handleSelect" style="background-color: #545c64;text-color: #fff;active-text-color:#ffd04b">
<!--    得开启router模式，这样才能根据index来寻找路径-->
    <el-menu-item index="/index" key="首页" class="home-page">Home page</el-menu-item>
    <el-menu-item index="/stu_inf" key="学生信息">Student information</el-menu-item>
    <el-submenu index="/edit" class="edit-student">
      <template slot="title">Edit</template>
      <el-menu-item index="/a_student">Add or Update student information</el-menu-item>
      <el-menu-item index="/d_student">Delete student</el-menu-item>
    </el-submenu>
    <el-submenu index="/major" class="student-major">
      <template slot="title">Major</template>
      <el-menu-item index="/COMP">COMP</el-menu-item>
      <el-menu-item index="/AMA">AMA</el-menu-item>
      <el-menu-item index="/AF">AF</el-menu-item>
      <el-menu-item index="/ACCT">ACCT</el-menu-item>
      <el-menu-item index="/EIE">EIE</el-menu-item>
      <el-menu-item index="/ABCT">ABCT</el-menu-item>
      <el-menu-item index="/CBS">CBS</el-menu-item>
    </el-submenu>
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
