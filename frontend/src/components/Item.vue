<script lang="ts">
import { defineComponent, reactive, computed } from "vue";
import LoadingIcon from "@/components/icons/LoadingIcon.vue";
import Plus from "@/components/icons/Plus.vue";
import { RequestStatus, formatIdr } from "@/helpers";
import { ItemStockView } from "@/view";
import { appState } from "@/App.vue";
export default defineComponent({
  components: {
    LoadingIcon,
    Plus,
  },
  setup() {
    const state = reactive({
      items: [] as Array<ItemStockView>,
      requestStatus: "Loading" as RequestStatus,
      searchInput: "",
    });

    const store = appState;

    const fetchItems = async () => {
      try {
        state.requestStatus = "Loading";

        const response = await fetch(
          `${process.env.VUE_APP_BASE_URL}/itemstocks`,
          {
            headers: {
              authorization: store.apiKey ?? "",
            },
          }
        );
        const items = await response.json();

        state.items = items;
        state.requestStatus = "Success";
      } catch (e) {
        console.log(e);
        state.requestStatus = "Error";
      }
    };

    fetchItems();

    const handleInputSearchItem = (e: any) => {
      state.searchInput = e.target.value;
    };

    const filteredItems = computed(() =>
      state.items.filter((itemStockView) =>
        `${itemStockView.item.name.toLowerCase()}${itemStockView.item.description.toLowerCase()}`.includes(
          state.searchInput.toLowerCase()
        )
      )
    );

    return {
      state,
      filteredItems,
      // Funcs
      handleInputSearchItem,
      formatIdr,
    };
  },
});
</script>
<template>
  <div>
    <div class="flex m-2 items-center">
      <span class="text-2xl font-semibold">Items</span>
      <loading-icon
        v-if="state.requestStatus === 'Loading'"
        class="text-2xl ml-2"
      />
    </div>
    <div class="flex items-center m-2">
      <router-link to="/items/new">
        <div>
          <button class="font-semibold bg-blue-500 text-white rounded-lg p-2">
            New
          </button>
        </div>
      </router-link>
      <input
        class="shadow appearance-none border rounded px-4 py-2 w-full bg-gray-100 ml-2 focus:shadow-outline"
        placeholder="Search item or description..."
        :value="state.searchInput"
        @input="handleInputSearchItem"
      />
    </div>
    <div>
      <div
        class="shadow-lg border-2 border-green-200 m-2 rounded-lg p-2 hover:bg-gray-100"
        v-for="itemStockView in filteredItems"
        :key="itemStockView.id"
      >
        <router-link :to="`/items/${itemStockView.item.id}`">
          <div class="text-lg font-bold" v-if="itemStockView.item">
            {{ itemStockView.item.name }}
          </div>
          <div class="text-sm italic" v-if="itemStockView.item">
            {{ itemStockView.item.description }}
          </div>
          <div class="font-semibold text-green-600" v-if="itemStockView.item">
            Price: {{ formatIdr(itemStockView.item.price) }}
          </div>
          <div class="font-semibold text-blue-600" v-if="itemStockView.item">
            Manuf. Price: {{ formatIdr(itemStockView.item.manufacturingPrice) }}
          </div>
          <div class="flex justify-between">
            <div
              class="font-semibold text-white bg-orange-600 rounded px-2 py-1"
            >
              In stock: {{ itemStockView.inStock }}
            </div>
            <router-link :to="`/stockins?itemId=${itemStockView.item.id}`">
              <div
                class="flex items-center bg-gray-700 text-white font-semibold px-2 py-1 rounded"
              >
                Stock in <plus class="ml-1" />
              </div>
            </router-link>
          </div>
        </router-link>
      </div>
    </div>
    <!-- <div class="m-2">
      {{ state }}
    </div> -->
  </div>
</template> 
