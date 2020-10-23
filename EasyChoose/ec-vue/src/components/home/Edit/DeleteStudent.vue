<template>
  <body id="poster">
  <el-form class="delete-container" label-position="left" label-width="0px">
    <h3 class="delete-title">Delete student</h3>
    <el-form-item>
      <el-input type="text" v-model="id" auto-complete="off" placeholder="Please enter the student ID you want to delete"></el-input>
    </el-form-item>
    <el-form-item style="width: 100%">
      <el-button type="primary" style="width: 100%;background: dodgerblue;border: none" v-on:click="delete_stu">删除</el-button>
    </el-form-item>
  </el-form>
  </body>
</template>

<script>
export default {
  name: 'DeleteStudent',
  data () {
    return {
      id: null
    }
  },
  methods: {
    delete_stu () {
      this.$confirm('This operation will delete this student information forever，will you choose to continue？', 'Tip', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        this.$axios.post('/d_student', {
          id: this.id
        }).then(resp => {
          if (resp.status === 200) {
            this.$router.replace({path: '/stu_inf'})
          }
        })
      })
    }
  }
}
</script>

<style scoped>
  .delete-container{
    border-radius: 15px;
    background-clip: padding-box;
    margin: 90px auto;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
  }
  .delete-title {
    margin: 0px auto 40px auto;
    text-align: center;
    color: #505458;
  }
</style>
