<template>
  <div>
    <div class="text-xl font-semibold m-2">Project</div>
  </div>
  <div class="flex items-center m-2">
    <router-link to="/projects">
      <button class="border-2 border-gray-500 font-semibold p-2 rounded-lg">Back</button>
    </router-link>
    <button @click="handleSave" class="bg-blue-500 text-white font-semibold p-2 rounded-lg ml-2">Save</button>
    <loading-icon v-if="state.requestStatus === 'Loading'" class="text-2xl" />
    <div v-if="state.requestStatus === 'Success'" class="font-semibold text-green-600">
      Success!
    </div>
  </div>
  <form class="m-2">
    <div class="my-2">
      <div class="font-semibold text-sm">Project Name</div>
      <input 
        :value="state.project.name" 
        class="w-full border-2 border-gray-300 p-2 rounded-lg focus:shadow-outline" 
        placeholder="Project Name..." 
        @input="handleChangeProjectName"
      />
    </div>
    <div class="my-2">
      <div class="font-semibold text-sm">Project Date ({{ makeReadableDateString(new Date(state.project.startDate)) }})</div>
      <input 
        :value="makeDateString(new Date(state.project.startDate))" 
        type="date" 
        class="w-full border-2 border-gray-300 p-2 rounded-lg focus:shadow-outline" 
        placeholder="Project Date..." 
        @input="handleChangeProjectDate"  
      />
    </div>
  </form>
  <div class="m-2">
    <div class="font-semibold text-lg text-red-600">Danger Zone</div>
  </div>
</template>
<script lang="ts">
import Vue, { defineComponent, reactive } from 'vue'
import { appState } from '@/App.vue'
import { useRouter } from 'vue-router'
import { RequestStatus, makeDateString, makeReadableDateString } from '@/helpers'
import { initialProject } from '@/modelinitials'
import LoadingIcon from '@/components/icons/LoadingIcon.vue'
import { Project } from '@/model'
export default defineComponent({
  components: {
    LoadingIcon
  },
  setup() {
    const router = useRouter()
    const projectId = router.currentRoute.value.params?.id as string
    const isNew = isNaN(parseInt(projectId))
    const store = appState   

    const state = reactive({
      project: {...initialProject},
      requestStatus: 'NotAsked' as RequestStatus
    })

    const fetchProject = async (id: string) => {
      try {
        const response = await fetch(`${store.baseUrl}/projects/${projectId}`, {
          headers: {
            'authorization': store.apiKey
          }
        })

        state.project = await response.json()
      } catch(e) {
        console.log(e)
      }
    }

    if(!isNew) {
      fetchProject(projectId)
    }

    const handleSave = async () => {
      try {
        state.requestStatus = 'Loading'

        const response = await fetch(`${store.baseUrl}/projects`, {
          method: 'POST',
          headers: {
            'authorization': store.apiKey,
            'content-type': 'application/json'
          },
          body: JSON.stringify({
            ...state.project,
            created_at: undefined,
            updated_at: undefined
          } as Project)
        }) 

        if (response.status !== 201) throw 'Error saving project'
        const responseData = await response.json() as { id: number };

        state.requestStatus = 'Success'
        // fetchProject(responseData.id.toString())
        router.push(`/projects/${responseData.id}`)
      } catch(e) {
        console.log(e)
        state.requestStatus = 'Error'
      }
    }

    const handleChangeProjectName = (e: any) => {
      state.project.name = e.target.value
    }

    const handleChangeProjectDate = (e: any) => {
      state.project.startDate = e.target.value
    }
    
    return {
      state,
      // funcs
      makeDateString,
      makeReadableDateString,
      handleSave,
      handleChangeProjectName,
      handleChangeProjectDate
    }
  }
})
</script>