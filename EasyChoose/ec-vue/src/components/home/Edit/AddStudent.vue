<template>
  <body id="poster">
  <el-form class="add-container" label-position="left" label-width="0px">
    <h3 class="add-title">Add or Update student information</h3>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.id" auto-complete="off" placeholder="Student ID"></el-input>
    </el-form-item>
    <el-form-item>
    <el-input type="text" v-model="NewStudentForm.major" auto-complete="off" placeholder="Major"></el-input>
  </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.name" auto-complete="off" placeholder="Name"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.gender" auto-complete="off" placeholder="Gender"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.grade" auto-complete="off" placeholder="Grade"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course1" auto-complete="off" placeholder="Subject 1"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course2" auto-complete="off" placeholder="Subject 2"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course3" auto-complete="off" placeholder="Subject 3"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course4" auto-complete="off" placeholder="Subject 4"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course5" auto-complete="off" placeholder="Subject 5"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course6" auto-complete="off" placeholder="Subject 6"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input type="text" v-model="NewStudentForm.course7" auto-complete="off" placeholder="Subject 7"></el-input>
    </el-form-item>
    <el-form-item style="width: 100%">
      <el-button type="primary" style="width: 100%;background: dodgerblue;border: none" v-on:click="confirm">确认</el-button>
    </el-form-item>
  </el-form>
  </body>
</template>

<script>
export default {
  name: 'AddStudent',
  data () {
    return {
      NewStudentForm: {
        id: null,
        major: '',
        name: '',
        gender: '',
        grade: '',
        course1: '',
        course2: '',
        course3: '',
        course4: '',
        course5: '',
        course6: '',
        course7: ''
      }
    }
  },
  methods: {
    confirm () {
      this.$axios.post('/a_student', {
        id: this.NewStudentForm.id,
        major: this.NewStudentForm.major,
        name: this.NewStudentForm.name,
        gender: this.NewStudentForm.gender,
        grade: this.NewStudentForm.grade,
        course1: this.NewStudentForm.course1,
        course2: this.NewStudentForm.course2,
        course3: this.NewStudentForm.course3,
        course4: this.NewStudentForm.course4,
        course5: this.NewStudentForm.course5,
        course6: this.NewStudentForm.course6,
        course7: this.NewStudentForm.course7
      }).then(successResponse => {
        if (successResponse.status === 200) {
          this.$router.replace({path: '/stu_inf'})
        }
      }).catch(fail => {})
    }
  }
}
</script>

<style scoped>
  .add-container{
    border-radius: 15px;
    background-clip: padding-box;
    margin: 90px auto;
    width: 350px;
    padding: 35px 35px 15px 35px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
  }
  .add-title {
    margin: 0px auto 40px auto;
    text-align: center;
    color: #505458;
  }
</style>
