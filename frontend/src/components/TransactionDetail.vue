<template>
  <!-- {{ uid() }} -->
  <div class="m-2">
    <div class="text-xl font-semibold ">Transaction Detail</div>
    <div v-if="state.project !== null" class="text-lg">For project: {{ state.project.name }}</div>
    <!-- <v-select :options="['Test']"></v-select> -->
    <div class="flex items-center">
      <router-link to="/">
        <button class="border-2 border-gray-500 rounded-lg font-semibold px-2 py-1">Back</button>
      </router-link>
      <button 
        class="bg-blue-500 text-white rounded-lg font-semibold px-2 py-1 ml-2"
        @click="handleSave"
      >
        Save
      </button>
      <loading-icon
        v-if="state.requestStatus === 'Loading'"
        class="text-2xl"
      />
    </div>
    <div v-if="state.requestStatus === 'Success'" class="font-semibold text-green-600">Success!</div>
    <div class="flex my-2">
      <div class="uppercase text-sm font-semibold">Search Item: </div> 
    </div>
    <div class="flex my-2">
      <input 
        class="w-full border-2 border-gray-300 rounded-lg p-2 focus:outline-none focus:shadow-outline" 
        placeholder="Item name..."
        :value="state.searchInput"
        @input="handleInputSearch"
      />
    </div>
    <div
      v-for="searchedItem in state.searchedItems"
      :key="searchedItem.item.id"
      class="p-1 border-2 border-gray-500 rounded-lg my-1"
    >
      <div 
        class="flex justify-between"
        @click="handleSelectSearchedItem(searchedItem)"
      >
        <div>{{ searchedItem.item.name }}</div>
        <div class="font-semibold">In stock: {{ searchedItem.inStock }}</div>
      </div>
    </div>
    <div v-if="state.selectedItem" class="flex items-center">
      <div class="uppercase text-sm font-semibold">
        Selected: 
      </div>
      <div class="ml-2">{{ state.selectedItem.item.name }}</div>
    </div>
    <div class="uppercase text-sm font-semibold">Qty: </div>
    <div class="flex font-semibold items-center justify-between">
      <!-- <div class="mx-1">Qty</div> -->
      <input 
        class="w-full p-2 border-2 rounded-lg focus:outline-none focus:shadow-outline" 
        placeholder="Qty..."
        :value="state.qty"
        @input="handleInputQty"
      />
      <button 
        class="border-2 border-gray-500 font-semibold p-2 rounded-lg"
        @click="handleInsertNewItemTransaction"
      >
        Insert
      </button>
    </div>
    <hr class="border-dashed border-2 border-green-300 my-2" />
    <div class="font-semibold text-lg my-2">
      Selected Items
    </div>
    <div 
      v-for="itemTransaction in state.transactionView.itemTransactions"
      :key="itemTransaction.uid"  
      class="flex items-center border-2 border-gray-500 justify-between p-2 m-1 rounded-md"
    >
      <div class="p-2">
        <div class="font-semibold">{{ itemTransaction.item.name }} x{{ itemTransaction.itemTransaction.qty }}</div>
        <div>{{ formatIdr(itemTransaction.item.price * itemTransaction.itemTransaction.qty) }}</div>
      </div>
      <div>
        <button 
          class="bg-red-600 text-white p-2 font-semibold rounded-full"
          @click="handleDeleteItemTransaction(itemTransaction.itemTransaction)"         
        >
          Delete
        </button>
      </div>
    </div>
    <hr class="border-dashed border-2 border-blue-300 my-2" />
    <div class="flex items-center text-lg font-semibold">
      <div>Custom Price?</div>
      <input 
        @click="handleTogglePriceIsCustom" 
        :checked="state.transactionView.transaction.priceIsCustom" 
        type="checkbox" 
        class="ml-2 leading-tight" 
      />
    </div>
    <div v-if="state.transactionView.transaction.priceIsCustom">
      <div>
        <input 
          class="appearance-none border-2 w-full border-gray-300 p-2 focus:shadow-outline rounded" 
          placeholder="Custom Price..." 
          :value="state.transactionView.transaction.customPrice"
          @input="handleInputTransactionCustomPrice"
        />
      </div>
      <div class="font-semibold">{{ formatIdr(state.transactionView.transaction.customPrice) }}</div>
    </div>
    <hr class="border-dashed border-2 border-red-300 my-2" />
    <div class="text-2xl font-semibold">Grand Total</div>
    <div class="font-semibold uppercase text-sm">Total:</div>
    <div class="text-2xl font-semibold text-green-500">
      {{ formatIdr(totalPrice) }}
    </div>
    <div class="font-semibold uppercase text-sm">
      Custom Price? 
      <span :class="state.transactionView.transaction.priceIsCustom ? 'text-green-500' : 'text-red-500'">
        {{ state.transactionView.transaction.priceIsCustom ? 'Yes' : 'No' }}
      </span>
    </div>
    <div v-if="state.transactionView.transaction.priceIsCustom" class="text-2xl font-semibold text-green-500">
      {{ formatIdr(state.transactionView.transaction.customPrice) }}
    </div>
    <div class="font-semibold uppercase text-sm">Final Price:</div>
    <div v-if="state.transactionView.transaction.priceIsCustom" class="text-2xl font-semibold text-green-500">
      {{ formatIdr(state.transactionView.transaction.customPrice) }}
    </div>
    <div v-else class="text-2xl font-semibold text-green-500">
      {{ formatIdr(totalPrice) }}
    </div>
    <div class="flex items-center justify-center">
      <button
        class="bg-blue-500 text-white font-semibold rounded-lg p-2"
        @click="handleSave"         
      >
        Save
      </button>
      <loading-icon
        v-if="state.requestStatus === 'Loading'"
        class="text-2xl"
      />
    </div>
    <div v-if="state.requestStatus === 'Success'" class="font-semibold text-green-600">Success!</div>
    <div 
      class="my-2 text-red-500 text-xl font-semibold flex justify-center"
    >
      Danger Zone
    </div>
    <div v-if="!isNaN(parseInt(transactionId))" class="flex justify-center">
      <button 
        class="bg-red-600 text-white font-semibold p-2 rounded-lg"
        @click="handleDeleteTransaction"
      >
        Delete
      </button>
    </div>
  </div>
  <!-- <div>
    {{ router }}
  </div> -->
</template>
<script lang="ts">
import Vue, { defineComponent, reactive, computed } from 'vue'
import { RequestStatus, formatIdr } from '@/helpers'
import { useRouter } from 'vue-router'
import { Project, ItemTransaction } from '@/model'
import ProjectDetailVue from './ProjectDetail.vue'
import { initialTransactionView, initialItemTransaction, initialItemTransactionView, initialItem } from '@/modelinitials'
import { TransactionView, ItemStockView, ItemTransactionView } from '@/view'
import { appState } from '@/App.vue'
// import vSelect from 'vue-select';
import ShortUniqueId from 'short-unique-id'
import { TransactionPostBody } from '@/postbody'
import LoadingIcon from '@/components/icons/LoadingIcon.vue'

export default defineComponent({
  components: {
    LoadingIcon
  },
  setup() {
    const router = useRouter()
    const transactionId = router.currentRoute.value.params?.id as string
    const projectId = router.currentRoute.value.query?.projectId as string
    const isNew = isNaN(parseInt(transactionId))
    const uid = new ShortUniqueId()

    const store = appState

    const state = reactive({
      requestStatus: 'NotAsked' as RequestStatus,
      transactionView: { ...initialTransactionView } as TransactionView,
      project: null as Project | null,
      searchInput: "",
      searchedItems: [] as ItemStockView[],
      selectedItem: null as ItemStockView | null,
      qty: 0,
      itemTransactionDeleteIds: [] as number[],
      // priceIsCustom: false
    })

    const totalPrice = computed(() => state.transactionView.itemTransactions.reduce((acc, itemTransaction) => {
      return acc + (itemTransaction.itemTransaction.qty * itemTransaction.item.price)
    }, 0))

    const fetchTransaction = async () => {
      try {
        const response =  await fetch(`${store.baseUrl}/transactions/view/${transactionId}`)
        state.transactionView = await response.json()
      } catch(e) {
        console.log(e)
      }
    }

    const fetchProject = async () => {
      try {
        const response = await fetch(`${store.baseUrl}/projects/${projectId}`)
        state.project = await response.json()
      } catch(e) {
        console.log(e)
      } 
    }

    if(!isNew) {
      fetchTransaction()
    }

    const searchItem = async (itemName: string) => {
      try {
        const response = await fetch(`${store.baseUrl}/itemsearch?name=${itemName}`, {
          headers: {
            "authorization": store.apiKey ?? ""
          }
        })

        return await response.json() as ItemStockView[]
      } catch(e) {
        console.log(e);
        return []
      }
    }

    fetchProject()

    const handleInputSearch = async (e: any) => {
      const searchInput = e.target.value as string
      state.searchInput = searchInput
      state.searchedItems = await searchItem(searchInput)
    } 

    const handleSelectSearchedItem = (searchedItem: ItemStockView) => {
      state.selectedItem = searchedItem
      state.searchInput = ""
      state.searchedItems = []
    }

    const handleInsertNewItemTransaction = () => {
      const newItemTransaction = {
        ...initialItemTransactionView,
        item: state.selectedItem ? state.selectedItem.item : null,
        itemTransaction: {
          ...initialItemTransaction,
          uid: uid(),
          qty: state.qty,
          itemId: state.selectedItem ? state.selectedItem.item.id : 0
        }
      } as ItemTransactionView;

      state.transactionView.itemTransactions = [newItemTransaction, ...state.transactionView.itemTransactions]
      state.selectedItem = null
      state.qty = 0
    }

    const handleInputQty = (e: any) => {
      state.qty = isNaN(parseInt(e.target.value)) ? 0 : parseInt(e.target.value)
    }

    const handleDeleteItemTransaction = (itemTransactionToDelete: ItemTransaction) => {
      const newItemTransactions = state.transactionView.itemTransactions
        .filter(itemTransaction => itemTransaction.itemTransaction.uid !== itemTransactionToDelete.uid)

      state.transactionView = {
        ...state.transactionView,
        itemTransactions: newItemTransactions
      }
      state.itemTransactionDeleteIds = [itemTransactionToDelete.id, ...state.itemTransactionDeleteIds]
    }

    const handleTogglePriceIsCustom = () => {
      state.transactionView.transaction.priceIsCustom = !state.transactionView.transaction.priceIsCustom
    }

    const handleInputTransactionCustomPrice = (e: any) => {
      state.transactionView.transaction.customPrice = isNaN(parseInt(e.target.value)) ? 0 : parseInt(e.target.value)
    }

    const handleSave = async () => {
      try {
        state.requestStatus = 'Loading'

        const response = await fetch(`${store.baseUrl}/transactionsave`, {
          method: 'POST',
          headers: {
            'authorization': store.apiKey ?? ""
          },
          body: JSON.stringify({ 
            transaction: {
              ...state.transactionView.transaction,
              projectId: isNaN(parseInt(projectId)) ? 0 : parseInt(projectId),
              created_at: undefined,
              updated_at: undefined
            },
            itemTransactions: state.transactionView.itemTransactions.map(itemTransactionView => {
              return { 
                ...itemTransactionView,
                itemTransaction: {
                  ...itemTransactionView.itemTransaction,
                  created_at: undefined,
                  updated_at: undefined
                }
              }
            }),
            itemTransactionDeleteIds: state.itemTransactionDeleteIds  
          } as TransactionPostBody) 
        })
        
        if (response.status !== 201) throw 'Error saving transaction'
        state.requestStatus = 'Success'

        router.push(`/transactions/${(await response.json()).id}?projectId=${projectId}`)
      } catch(e) {
        console.log(e);
        state.requestStatus = 'Error'
      }
    }

    const handleDeleteTransaction = async () => {
      const confirmation = confirm('Are you sure you want to delete transaction?')

      if (confirmation) {
        state.requestStatus = 'Loading'

        try {
          const response = await fetch(`${store.baseUrl}/transactions/${transactionId}`, {
            method: 'DELETE',
            headers: {
              'authorization': store.apiKey ?? ""
            }
          })

          if (response.status !== 200) throw 'Error deleting transaction'
          state.requestStatus = 'Success'
          router.push('/')
        } catch(e) {
          console.log(e)
          state.requestStatus = 'Error'
        }
      }
    }

    return  {
      state,
      router, 
      transactionId,
      uid,
      totalPrice,
      // Funcs
      handleInputSearch,
      handleSelectSearchedItem,
      handleInsertNewItemTransaction,
      handleInputQty,
      handleDeleteItemTransaction,
      handleTogglePriceIsCustom,
      handleInputTransactionCustomPrice,
      handleDeleteTransaction,
      formatIdr,
      handleSave
    }
  }
})
</script>