<template>
  <el-table :data="tableData" style="width: 100%">
    <el-table-column prop="id" label="Student ID" width="180"></el-table-column>
    <el-table-column prop="major" label="Subject Code" width="180"></el-table-column>
    <el-table-column prop="name" label="Name" width="180"></el-table-column>
    <el-table-column prop="gender" label="Gender" width="180"></el-table-column>
    <el-table-column prop="grade" label="Grade" width="180"></el-table-column>
    <el-table-column prop="course1" label="Subject 1" width="180"></el-table-column>
    <el-table-column prop="course2" label="Subject 2" width="180"></el-table-column>
    <el-table-column prop="course3" label="Subject 3" width="180"></el-table-column>
    <el-table-column prop="course4" label="Subject 4" width="180"></el-table-column>
    <el-table-column prop="course5" label="Subject 5" width="180"></el-table-column>
    <el-table-column prop="course6" label="Subject 6" width="180"></el-table-column>
    <el-table-column prop="course7" label="Subject 7" width="180"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'StuInf',
  data () {
    return {
      tableData: []
    }
  },
  mounted () {
    this.loadStudentInf()
  },
  methods: {
    loadStudentInf () {
      this.$axios.get('/stu_inf').then(resp => {
        if (resp.status === 200) {
          this.tableData = resp.data
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
