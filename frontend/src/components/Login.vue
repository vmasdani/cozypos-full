<template>
  <div class="bg-gray-700 h-screen flex items-center justify-center">
    <div class="shadow-2xl flex flex-col items-center bg-white rounded-lg p-5">
      <div class="text-xl font-bold" @click="handleChangeApiKey">
        Cozy PoS 
        <!-- {{ store.apiKey }} -->
      </div>
      <loading-icon v-if="state.requestStatus === 'Loading'" class="text-2xl" />
      <div v-if="state.requestStatus === 'Error'" class="font-semibold text-red-600">{{ state.errorMessage }}</div>
      <div>
        <input 
          class="shadow bg-gray-200 appearance-none border text-center text-gray-700 rounded w-full p-2 leading-tight my-3 focus:outline-none focus:shadow-outline font-bold" 
          type="text"
          placeholder="Username..."
          @input="changeUsername"
          :value="state.username"
        />
      </div>
      <div>
        <input 
          class="shadow bg-gray-200 appearance-none border text-center text-gray-700 rounded w-full p-2 leading-tight my-3 focus:outline-none focus:shadow-outline font-bold" 
          type="password" 
          placeholder="Password..." 
          @input="changePassword"
          :value="state.password"
        />
      </div>
      <div>
        <button 
          class="shadow bg-gray-500 text-white hover:bg-gray-600 font-bold p-2 rounded-lg"
          @click="handleLogin"
        >
          Login
        </button>
      </div>
    </div> 
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { appState } from '@/App.vue'
import { RequestStatus } from '@/helpers'
import LoadingIcon from '@/components/icons/LoadingIcon.vue'

export default defineComponent({
  components: {
    LoadingIcon
  },
  setup(props, ctx) {
    const state = reactive({
      requestStatus: 'NotAsked' as RequestStatus,
      username: "",
      password: "",
      apiKey: "",
      errorMessage: ""
    })
    
    const store = appState

    const changeUsername = (e: any) => {
      state.username = e.target.value.trim()
    }

    const changePassword = (e: any) => {
      state.password = e.target.value.trim()
    }

    const handleLogin = async () => {
      // ctx.emit('pressLogin', 'test')
      // props.onLogin('testusername', 'testpassword')

      try {
        state.requestStatus = 'Loading'

        const response = await fetch(`${store.baseUrl}/login`, {
          method: 'POST',
          body: JSON.stringify({ username: state.username, password: state.password })
        })

        if (response.status !== 200) throw await response.text()
        const apiKey = await response.text()

        localStorage.setItem('apiKey', apiKey)
        store.apiKey = apiKey
        store.loggedIn = true
        state.requestStatus = 'Success'
      } catch(e) {
        console.log(e)
        state.requestStatus = 'Error'
        state.errorMessage = e
      }
    }

    const handleChangeApiKey = () => {
      store.apiKey = ' changed!'
    }

    return {
      props,
      state,
      store,
      // Funcs
      changeUsername,
      changePassword,
      handleLogin,
      handleChangeApiKey
    }
  }
})
</script>