<template>
  <div>
    <div class="flex m-2 items-center">
      <div>
        <router-link to="/items">
          <button class="border-2 border-gray-500 px-2 py-1 rounded-lg">Back</button>
        </router-link>
      </div>
      <div class="text-xl font-semibold ml-2">Stock In</div>
      <div v-if="state.requestStatus === 'Loading'">
        <loading-icon class="text-2xl" />
      </div>
    </div>
    <!-- <div>
      {{ itemId }}
    </div> -->
    <div v-if="state.itemStockInsView !== null" class="m-2">
      <div class="text-xl" >
        {{ state.itemStockInsView.item.name }}
      </div>
    </div>
    <div class="flex m-2 items-center">
      <div class="uppercase text-sm font-semibold">Select Project:</div>
      <div class="ml-2">
        <select class="block bg-gray-200 appearance-none p-2 rounded-lg">
          <option
            v-for="project in state.projects"
            :key="project.id"
            @click="handleSelectProject(project)"
          >
            {{ project.name }}
          </option>
        </select>
      </div>
    </div>
    <form>
      <div class="flex m-2">
        <input 
          class="border-2 border-gray-400 px-2 rounded-lg font-semibold" 
          type="number"
          :value="state.qty"
          placeholder="Qty..." 
          @input="handleInputStockInQty"
        />
        <button 
          type="submit" 
          class="ml-1 font-semibold rounded-lg p-2 bg-blue-500 text-white"
          @click="handleInsertStockIn"
        >
          Insert
        </button>
      </div>
    </form>
    <div v-if="state.itemStockInsView" class="m-2">
      <div 
        v-for="stockInView in state.itemStockInsView.stockIns" 
        :key="stockInView.stockIn.id"
        class="border-2 border-purple-500 p-2 my-1 rounded-lg"
      >
        <div class="flex justify-between">
          <div class="text-xl font-semibold">
            x{{ stockInView.stockIn.qty }}
          </div>
          <div class="bg-gray-600 px-2 py-1 text-white rounded-full font-semibold">{{ stockInView.stockIn.pic }}</div>
        </div>
        <div>On: {{ new Date(stockInView.stockIn.createdAt).toString() }}</div>
        <div v-if="stockInView.project">For project: <span class="font-semibold p-1 rounded-lg bg-orange-700 text-white">{{ stockInView.project.name }}</span></div>
      </div>
    </div>
    <!-- <div v-if="state.itemStockInsView != null">
      {{ state.itemStockInsView.stockIns }}
    </div> -->
  </div>
</template>
<script lang="ts">
import Vue, { defineComponent, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { RequestStatus } from '@/helpers'
import { appState } from '@/App.vue'
import LoadingIcon from '@/components/icons/LoadingIcon.vue'
import { ItemStockInsView } from '@/view'
import { StockInPostBody } from '@/postbody'
import { Project, StockIn } from '@/model'
import { initialStockIn } from '@/modelinitials'
export default defineComponent({
  components: { 
    LoadingIcon
  },
  setup() {
    const router = useRouter()
    const itemId = router.currentRoute.value.query?.itemId as string

    const state = reactive({
      requestStatus: 'NotAsked' as RequestStatus,
      itemStockInsView: null as ItemStockInsView | null,
      projects: [] as Project[],
      selectedProject: null as Project | null,
      qty: 0
    })
     
    const store = appState

    const fetchItemStockIns = async () => {
      try {
        state.requestStatus = 'Loading'
        
        const response = await fetch(`${process.env.VUE_APP_BASE_URL}/items/${itemId}/stockins`, {
          headers: {
            "authorization": store.apiKey
          }
        })

        state.itemStockInsView = await response.json() as ItemStockInsView 

        state.requestStatus = 'Success'
      } catch(e) {
        state.requestStatus = 'Error'
      }
    }

    const fetchProjects = async () => {
      try {
        const response = await fetch(`${store.baseUrl}/projects`, {
          headers: {
            'authorization': store.apiKey
          }
        })

        const projects = await response.json() as Project[]
        state.projects = projects 
        state.selectedProject = projects.length > 0 ? projects[0] : null
     
      } catch(e) {
        console.log(e)
      }
    }

    if(itemId) {
      fetchItemStockIns()
    }
    fetchProjects()

    const handleInsertStockIn = async (e: any) => {
      e.preventDefault()
      
      try {
        state.requestStatus = 'Loading'

        const response = await fetch(`${store.baseUrl}/itemstockinsadd`, {
          method: 'POST',
          headers: {
            'authorization': store.apiKey
          },
          body: JSON.stringify({ 
            ...initialStockIn,
            qty: state.qty,
            itemId: isNaN(parseInt(itemId)) ? 0 : parseInt(itemId),
            projectId: state.selectedProject ? state.selectedProject.id : 0
          } as StockIn)
        })

        if (response.status !== 201) throw 'Error adding stockin'
        
        state.requestStatus = 'Success'
        fetchItemStockIns()
      } catch(e) {
        console.log(e)
        state.requestStatus = 'Error'
      }
    }

    const handleSelectProject = (project: Project) => {
      state.selectedProject = project
    }

    const handleInputStockInQty = (e: any) => {
      state.qty = isNaN(parseInt(e.target.value)) ? 0 : parseInt(e.target.value)
    }

    return {
      itemId,
      state,
      store,
      // Funcs
      handleInsertStockIn,
      handleSelectProject,
      handleInputStockInQty
    }
  }
})
</script>