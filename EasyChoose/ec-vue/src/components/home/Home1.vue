<template>
  <div>
    <NavMenu></NavMenu>
    <router-view/>
  </div>
</template>

<script>
import NavMenu from '../NavMenu1'
export default {
  name: 'Home1',
  components: {NavMenu}
}
</script>

<style scoped>

</style>
