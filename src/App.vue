<template class="dark:text-gray-200">
  <NavBar/>
  <div class="container p-4 md:p-20 ">
    <router-view/>
  </div>
</template>

<script>
import NavBar from "@/components/NavBar";
export default {

  components: {
    NavBar
  }
}
</script>