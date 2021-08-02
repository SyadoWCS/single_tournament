<template>
   <img v-bind:src="image_src">
</template>

<script>
//import { defineComponent } from '@vue/composition-api'

export default {
   name: "CreateTournament",
   data() {
      return {
         image_src: require("@/assets/tournament.png")
      }
   }
}
</script>
