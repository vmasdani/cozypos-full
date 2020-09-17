<template>
  <!-- {{ state }} -->
  <div class="flex items-center m-2">
    <div class="text-2xl font-semibold">Transactions</div>
    <loading-icon v-if="state.requestStatus === 'Loading'" class="ml-2 text-2xl " /> 
  </div>
  <div class="flex items-center">
    <div 
      class="block uppercase tracking-wide text-gray-700 text-xs font-bold mx-1"
    >
      Select project:{{ " " }}
    </div>
    <div class="relative mx-1">
      <select class="block appearance-none w-full bg-gray-200 border border-blue-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
        <!-- <option>test</option> -->
        <option 
          v-for="project in state.projects" 
          :key="project.id"
          @click="selectProject(project)"
        >
          {{ project.name }}
        </option>
      </select>
    </div>
  </div>
  <div v-if="state.projectTransactionsView !== null" class="flex items-center m-2">
    <div>
      <router-link :to="`/transactions/new?projectId=${state.selectedProject.id}`">
        <button class="font-semibold bg-blue-500 text-white rounded-lg p-2">
          New
        </button>
      </router-link>
    </div>
    <input
      type="text"
      class="w-full ml-2 shadow bg-gray-100 border p-2 rounded-lg focus:shadow-outline" 
      v-model="state.searchInput"
      placeholder="Search containing items..." 
    />
  </div> 
  
  <div class="m-2" v-if="state.projectTransactionsView">
    Showing <span class="font-bold">{{ filteredProjectTransactionsView.length }}/{{ state.projectTransactionsView.transactions.length }}</span> transactions
  </div>
  <div class="mt-3">
    <div class="mx-2">
      <div 
        v-for="transactionView in filteredProjectTransactionsView" 
        :key="transactionView.transaction.id" 
        class="list-group-item my-2"
      >
        <router-link :to="`/transactions/${transactionView.transaction.id}?projectId=${state.selectedProject.id}`">
          <div class="shadow-lg border-2 border-solid border-blue-200 rounded-lg px-4 py-2">
            <div class="flex justify-between items-center">
              <h3
                :class="`text-2xl font-bold ${comparePriceIsCustomColor(transactionView)}`"
              >
                {{
                  formatIdr(transactionView.transaction.priceIsCustom
                    ? transactionView.transaction.customPrice
                    : transactionView.totalPrice
                  )  
                }}
              </h3>
              <div class="rounded-full shadow-lg text-white bg-gray-600 font-semibold px-2">{{ transactionView.transaction.cashier }}</div>
            </div>
            <div class="text-sm font-bold text-gray-800 uppercase">Orig: {{ formatIdr(transactionView.totalPrice)  }}</div>
            <div class="text-sm font-thin italic">
              {{ 
                transactionView.itemTransactions
                  .map(itemTransaction => `${itemTransaction.item.name} x${itemTransaction.itemTransaction.qty}`)
                  .join(', ')
              }}
            </div>
          </div>
        </router-link>
      </div>
    </div>
  </div>
  <!-- <div>
    <pre>
{{ filteredProjectTransactionsView }}
    </pre>
  </div> -->
</template>
<script lang="ts">
import { defineComponent, reactive, computed } from 'vue'
import { Project } from  '../model'
import { RequestStatus, formatIdr } from '../helpers'
import { ProjectTransactionsView, TransactionView } from '@/view'
import LoadingIcon from '@/components/icons/LoadingIcon.vue'
import { appState } from '@/App.vue'

export default defineComponent({
  components: {
    LoadingIcon
  },
  setup() {
    const state = reactive({
      selectedProject: null as Project | null,
      projects: [] as Project[],
      projectDetails: null,
      projectTransactionsView: null as ProjectTransactionsView | null,
      isLoading: false,
      searchInput: "",
      requestStatus: 'NotAsked' as RequestStatus
    })

    const store = appState

    const fetchInitialProject = async () => {
      try {
        state.requestStatus = 'Loading'

        const response = await fetch(`${store.baseUrl}/projects`)
        const projectsData = (await response.json()) as Project[]

        if(projectsData) {
          state.projects = projectsData
          state.requestStatus = 'Success'
        }
      } catch(e) {
        state.requestStatus = 'Error'
      }
    }

    // Fetch
    fetchInitialProject()

    const selectProject = (project: Project) => {
      console.log('Project id:', project.id)
      state.selectedProject = project
      fetchProjectTransactionsView(project.id)
    }

    const comparePriceIsCustomColor = (transactionView: TransactionView) => {
      if(!transactionView.transaction.priceIsCustom) {
        return ''
      } else {
        if(transactionView.totalPrice < transactionView.transaction.customPrice) {
          return 'text-green-500'
        } else {
          return 'text-red-500'
        }
      }
    }

    const fetchProjectTransactionsView = async (id: number) => {
      try {
        state.requestStatus = "Loading"

        const response = await fetch(`${store.baseUrl}/projects/${id}/transactions`)
        const projectTransactionsView = await response.json()

        state.projectTransactionsView = projectTransactionsView
        state.requestStatus = "Success"
      } catch(e) {
        console.log(e);
        state.requestStatus = "Error"
      }
    }

    const filteredProjectTransactionsView = computed(() => state.projectTransactionsView 
      ? state.projectTransactionsView.transactions
          .filter(transaction => {
            return transaction.itemTransactions
              .map(itemTransaction => itemTransaction.item.name.toLowerCase())
              .flat()
              .join('')
              .toLowerCase()
              .includes(state.searchInput.toLowerCase())
          })
      : []
    )

    return {
      state,
      filteredProjectTransactionsView,
      // Funcs
      selectProject,
      formatIdr,
      comparePriceIsCustomColor
    }
  }
})
</script>