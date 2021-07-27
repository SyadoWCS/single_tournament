<template>
   <div class="container">
      <h1>{{ message }}</h1>
   </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import {useStore} from 'vuex'

export default {
   name: "Home",
   setup() {
      const message = ref('ログイン状態です！')
      const store = useStore()

      onMounted(async () => {
         try {
            // user情報を取得
            // ログイン情報は、Cookieに保存してあるので、
            // リクエストするだけでOK
            const { data } = await axios.get('user')
            console.log(data)
            message.value = `Hi ${data.FirstName} ${data.LastName}`
            // actionsに設定したパラメータ名を設定
            await store.dispatch('setAuth', true)
         } catch(e) {
            await store.dispatch('setAuth', false)
         }
      })

      return {
         message
      }
   }
}
</script>