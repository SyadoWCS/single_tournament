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
        const message = ref('You are not logged in!')

        onMounted(async () => {
            try {
                // user情報を取得
                // ログイン情報は、Cookieに保存してあるので、
                // リクエストするだけでOK
                const { data } = await axios.get('user')
                message.value = `Hi ${data.first_name} ${data.last_name}`

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