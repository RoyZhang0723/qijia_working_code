<template>
  <el-table :data="tableData2" style="width: 100%">
    <el-table-column prop="id" label="ID" width="180"></el-table-column>
    <el-table-column prop="subject_code" label="Subject Code" width="180"></el-table-column>
    <el-table-column prop="subject_name" label="Subject Name" width="600"></el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'AMACourse',
  data () {
    return {
      tableData2: []
    }
  },
  mounted () {
    this.loadAMACourse()
  },
  methods: {
    loadAMACourse () {
      this.$axios.get('/AMA').then(resp => {
        if (resp.status === 200) {
          this.tableData2 = resp.data
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
