<template>
  <div>
    <div v-if="store.loggedIn">
      <nav class="sticky top-0 flex items-center justify-between flex-wrap bg-gray-700 p-3 z-20">
        <div class="flex items-center flex-shrink-0 text-white mr-6">
          <a href="#">
            <span class="font-semibold text-xl tracking-tight">Cozy PoS</span>
          </a>
        </div>
        <div class="block lg:hidden">
          <button 
            class="flex items-center px-3 py-2 border rounded text-white border-white-400 hover:text-white hover:border-white"
            @click="handleToggleExpandNavbar"
          >
            <svg class="fill-current h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Menu</title><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>
          </button>
        </div>
        <div :class="`w-full ${ state.expandNavbar ? 'block' : 'hidden' } flex-grow lg:flex lg:items-center lg:w-auto`">
          <div class="text-sm lg:flex-grow">
            <a href="#" class="block mt-4 lg:inline-block lg:mt-0 text-white font-semibold hover:text-white mr-4">
              Transactions
            </a>
            <a href="#/items" class="block mt-4 lg:inline-block lg:mt-0 text-white font-semibold hover:text-white mr-4">
              Items
            </a>
            <a href="#/projects" class="block mt-4 lg:inline-block lg:mt-0 text-white font-semibold hover:text-white mr-4">
              Projects
            </a>
          </div>
          <div>
            <button 
              class="bg-red-600 text-white font-semibold py-1 px-2 text-sm rounded-lg lg:mt-0 mt-4"
              @click="handleLogout"
            >
              Logout
            </button>
            <!-- <a href="#" class="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0">Download</a> -->
          </div>
        </div>
      </nav>
      <router-view v-if="store.loggedIn" />
    </div>
    <login @pressLogin="handlePressLogin" :onLogin="handleLogin" myComponent="Test Prop" v-else />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, watch } from 'vue';
import Login from './components/Login.vue'
import vSelect from 'vue-select';

export const appState = reactive({
  apiKey: localStorage.getItem("apiKey"),
  loggedIn: false,
  baseUrl: process.env.VUE_APP_BASE_URL
})

export default defineComponent({
  name: 'App',
  components: {
    Login,
  },
  setup() {
    const state = reactive({
      loggedIn: false,
      testProp: `a`,
      expandNavbar: false
    })

    const store = appState

    if(store.apiKey !== null && store.apiKey !== "") {
      store.loggedIn = true
    }

    // const handleIncrementTestProp = () => {
    //   state.testProp = `${state.testProp}a`
    // }

    const handleLogin = (username: string, password: string) => {
      state.testProp = `${state.testProp}a`
      // state.loggedIn = true
      appState.loggedIn = true
      console.log('onLogin', username, password)
    }

    const handleLogout = () => {
      // state.loggedIn = false
      localStorage.removeItem("apiKey")
      appState.loggedIn = false
      appState.apiKey = ""
    }

    const handlePressLogin = () => {
      console.log('Emit press login!')
    }

    const handleToggleExpandNavbar = () => {
      state.expandNavbar = !state.expandNavbar
    }

    return {
      state,
      store,
      // Funcs
      handleLogin,
      handleLogout,
      handlePressLogin,
      handleToggleExpandNavbar
    }
  }
});
</script>
