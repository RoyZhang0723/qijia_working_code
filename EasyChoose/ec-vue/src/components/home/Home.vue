<template>
  <body id="home_body">
  <div>
    <NavMenu></NavMenu>
    <router-view/>
  </div>
  </body>
</template>

<script>
import NavMenu from './NavMenu'
export default {
  name: 'Home',
  components: {NavMenu}
}
</script>

<style scoped>
</style>
