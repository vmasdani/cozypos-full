<script lang="ts">
import Vue, { defineComponent, reactive } from "vue";
import { formatIdr, RequestStatus } from "@/helpers";
import { initialItem, initialProject } from "@/modelinitials";
import { Item, Project } from "@/model";
import { useRoute, useRouter } from "vue-router";
import { appState } from "../App.vue";
import LoadingIcon from "@/components/icons/LoadingIcon.vue";
import { ItemPostBody } from "../postbody";

export default defineComponent({
  components: {
    LoadingIcon,
  },
  setup() {
    const router = useRouter();
    const itemId = router.currentRoute.value.params?.id as string;
    const isNew = isNaN(parseInt(itemId ? itemId : ""));

    const state = reactive({
      requestStatus: "NotAsked" as RequestStatus,
      withInitialStock: false,
      projects: [] as Project[],
      selectedProject: { ...initialProject } as Project,
      qty: 0,
      item: { ...initialItem },
    });

    const store = appState;

    const fetchItem = async () => {
      try {
        state.requestStatus = "Loading";

        const response = await fetch(
          `${process.env.VUE_APP_BASE_URL}/items/${itemId}`,
          {
            headers: {
              authorization: store.apiKey ?? "",
            },
          }
        );

        const item = await response.json();
        console.log("item:", item);

        // state.item = item
        state.item = item;
        state.requestStatus = "Success";
      } catch (e) {
        console.log(e);
        state.requestStatus = "Error";
      }
    };

    const fetchProjects = async () => {
      try {
        const response = await fetch(
          `${process.env.VUE_APP_BASE_URL}/projects`,
          {
            headers: {
              authorization: store.apiKey ?? "",
            },
          }
        );

        state.projects = (await response.json()) as Project[];

        // if (state.projects.length > 0) {
        //   state.selectedProject = state.projects[0];
        // }
      } catch (e) {
        console.log(e);
      }
    };

    if (itemId && !isNew) {
      fetchItem();
    }

    fetchProjects();

    const handleChangeItemName = (e: any) => {
      state.item.name = e.target.value;
    };

    const handleChangeItemDescription = (e: any) => {
      state.item.description = e.target.value;
    };

    const handleChangeItemPrice = (e: any) => {
      state.item.price = isNaN(parseInt(e.target.value))
        ? 0
        : parseInt(e.target.value);
    };

    const handleChangeItemManufacturingPrice = (e: any) => {
      state.item.manufacturingPrice = isNaN(parseInt(e.target.value))
        ? 0
        : parseInt(e.target.value);
    };

    const handleToggleWithInitialStock = () => {
      state.withInitialStock = !state.withInitialStock;
    };

    const handleInputInitialStock = (e: any) => {
      state.qty = isNaN(parseInt(e.target.value))
        ? 0
        : parseInt(e.target.value);
    };

    const handleSave = async () => {
      if (state.withInitialStock && state.selectedProject.id === 0) {
        alert(
          "Please select project first, if you want to create with initial stock."
        );
      } else {
        const itemToSave: ItemPostBody = {
          item: {
            ...state.item,
            created_at: undefined,
            updated_at: undefined,
          } as Item,
          withInitialStock: state.withInitialStock,
          initialStockQty: state.qty,
          project: state.selectedProject,
        };

        try {
          state.requestStatus = "Loading";

          const response = await fetch(
            `${process.env.VUE_APP_BASE_URL}/itemsave`,
            {
              method: "POST",
              headers: {
                authorization: store.apiKey ?? "",
              },
              body: JSON.stringify(itemToSave),
            }
          );

          state.requestStatus = "Success";
          router.push("/items");
        } catch (e) {
          state.requestStatus = "Error";
        }
      }
    };

    const handleDeleteItem = () => {
      const confirmation = confirm(
        `Are you sure you want to delete item: ${state.item.name}?`
      );
    };

    const handleSelectProject = (e: any) => {
      const projectId = e.target.value as string;

      if (state.selectedProject) {
        state.selectedProject.id = isNaN(parseInt(projectId))
          ? 0
          : parseInt(projectId);
      }
    };

    return {
      state,
      router,
      isNew,
      // Funcs
      formatIdr,
      handleChangeItemName,
      handleSelectProject,
      handleChangeItemDescription,
      handleChangeItemPrice,
      handleChangeItemManufacturingPrice,
      handleToggleWithInitialStock,
      handleInputInitialStock,
      handleSave,
      handleDeleteItem,
    };
  },
});
</script>
<template>
  <div class="flex m-2 items-center">
    <div class="text-xl font-semibold">Item Detail</div>
    <div v-if="state.requestStatus === 'Loading'">
      <loading-icon class="text-2xl" />
    </div>
  </div>
  <div class="m-2">
    <div class="flex">
      <router-link to="/items">
        <div>
          <button
            class="border-2 font-semibold border-gray-500 px-2 py-1 rounded-lg"
          >
            Back
          </button>
        </div>
      </router-link>
      <div class="ml-2">
        <button
          @click="handleSave"
          class="border-2 font-semibold bg-blue-500 text-white px-2 py-1 rounded-lg"
        >
          Save
        </button>
      </div>
      <!-- <div>{{ state.item.id }}</div> -->
    </div>
    <div>
      <div class="text-sm font-semibold">Item Name</div>
      <input
        @input="handleChangeItemName"
        :value="state.item.name"
        class="border font-semibold w-full p-2 rounded-lg"
        placeholder="Item name..."
      />
    </div>
    <div>
      <div class="text-sm font-semibold">Description</div>
      <input
        @input="handleChangeItemDescription"
        :value="state.item.description"
        class="border font-semibold w-full p-2 rounded-lg"
        placeholder="Item Description..."
      />
    </div>
    <div>
      <div class="text-sm font-semibold">Price</div>
      <input
        :value="state.item.price"
        class="border font-semibold w-full p-2 rounded-lg"
        @input="handleChangeItemPrice"
        placeholder="Item Price..."
        type="number"
      />
    </div>
    <div>
      <div class="text-sm font-semibold">Funding Cost</div>
      <input
        :value="state.item.manufacturingPrice"
        class="border font-semibold w-full p-2 rounded-lg"
        @input="handleChangeItemManufacturingPrice"
        placeholder="Funding Cost..."
        type="number"
      />
    </div>
    <div v-if="isNew">
      <div class="flex my-1">
        <div class="text-sm font-semibold">With Initial Stock ?</div>
        <input
          @click="handleToggleWithInitialStock"
          :checked="state.withInitialStock"
          type="checkbox"
          class="border-4 ml-2"
        />
      </div>
      <div v-if="state.withInitialStock" class="flex items-center">
        <div>
          <input
            type="number"
            :value="state.qty"
            @input="handleInputInitialStock"
            class="border font-semibold w-full p-2 rounded-lg"
            placeholder="Initial Stock..."
          />
        </div>
        <div class="flex flex-col items-center m-1">
          <div class="text-center">For Project?</div>
          <select @input="handleSelectProject" class="appearance-none p-2">
            <option></option>
            <option
              :value="project.id"
              v-for="project in state.projects"
              :key="project.id"
            >
              {{ project.name }}
            </option>
          </select>
        </div>
      </div>
    </div>
  </div>
  <div class="m-2">
    <div class="font-semibold text-red-700">Danger Zone</div>
    <div v-if="!isNew" class="my-2">
      <button
        @click="handleDeleteItem"
        class="bg-red-700 text-white font-semibold p-2 rounded-lg"
      >
        Delete
      </button>
    </div>
  </div>
  <!-- <div>
    {{ state }}
  </div> -->
</template>