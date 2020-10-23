<template>
  <el-table :data="tableData1" style="width: 100%">
    <el-table-column prop="id" label="ID" width="180"></el-table-column>
    <el-table-column prop="subject_code" label="Subject Code" width="180"></el-table-column>
    <el-table-column prop="subject_name" label="Subject Name" width="600"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'CompCourse',
  data () {
    return {
      tableData1: []
    }
  },
  mounted () {
    this.loadCOMPCourse()
  },
  methods: {
    loadCOMPCourse () {
      this.$axios.get('/COMP').then(resp => {
        if (resp.status === 200) {
          this.tableData1 = resp.data
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
