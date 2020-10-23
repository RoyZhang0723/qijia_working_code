<template>
  <div>
    <el-table ref="multipleTable" :data="tableData" tooltip-effect="dark" style="width: 100%" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column prop="id" label="ID" width="180"></el-table-column>
      <el-table-column prop="subject_code" label="Subject Code" width="180"></el-table-column>
      <el-table-column prop="subject_name" label="Subject Name" width="700" show-overflow-tooltip></el-table-column>
    </el-table>
    <el-button @click="confirm" style="width: 100px;color: dodgerblue">Confirm</el-button>
  </div>
</template>

<script>
export default {
  name: 'RegisterCourse',
  data () {
    return {
      tableData: [],
      multipleSelection: [],
      registerForm: {
        username: this.$store.state.user.username,
        password: this.$store.state.user.password,
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
  mounted () {
    this.loadAllCourse()
  },
  methods: {
    loadAllCourse () {
      this.$axios.get('all_course').then(resp => {
        if (resp.status === 200) {
          this.tableData = resp.data
        }
      })
    },
    handleSelectionChange (val) {
      this.multipleSelection = val
    },
    confirm () {
      if (this.multipleSelection.length > 7) {
        alert('You can register at most 7 course at once')
      } else {
        if (this.multipleSelection[0] !== undefined) {
          this.registerForm.course1 = this.multipleSelection[0].subject_code
        }
        if (this.multipleSelection[1] !== undefined) {
          this.registerForm.course2 = this.multipleSelection[1].subject_code
        }
        if (this.multipleSelection[2] !== undefined) {
          this.registerForm.course3 = this.multipleSelection[2].subject_code
        }
        if (this.multipleSelection[3] !== undefined) {
          this.registerForm.course4 = this.multipleSelection[3].subject_code
        }
        if (this.multipleSelection[4] !== undefined) {
          this.registerForm.course5 = this.multipleSelection[4].subject_code
        }
        if (this.multipleSelection[5] !== undefined) {
          this.registerForm.course6 = this.multipleSelection[5].subject_code
        }
        if (this.multipleSelection[6] !== undefined) {
          this.registerForm.course7 = this.multipleSelection[6].subject_code
        }

        this.$axios.post('/register_course', {
          username: this.registerForm.username,
          password: this.registerForm.password,
          course1: this.registerForm.course1,
          course2: this.registerForm.course2,
          course3: this.registerForm.course3,
          course4: this.registerForm.course4,
          course5: this.registerForm.course5,
          course6: this.registerForm.course6,
          course7: this.registerForm.course7
        }).then(resp => {
          if (resp.status === 200) {
            this.$router.replace({path: '/index1'})
          }
        })
      }
    }
  }
}
</script>

<style scoped>

</style>
