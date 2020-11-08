<script lang="ts">
import { defineComponent, reactive } from "vue";
import { ProjectsView } from "@/view";
import { formatIdr, RequestStatus, makeReadableDateString } from "@/helpers";
import LoadingIcon from "@/components/icons/LoadingIcon.vue";

export default defineComponent({
  components: {
    LoadingIcon,
  },
  setup() {
    const state = reactive({
      requestStatus: "NotAsked" as RequestStatus,
      projectsView: null as ProjectsView | null,
    });

    const fetchProjectsView = async () => {
      try {
        state.requestStatus = "Loading";

        const response = await fetch(
          `${process.env.VUE_APP_BASE_URL}/projectsview`
        );
        const projectsViewData = await response.json();

        state.projectsView = projectsViewData;
        state.requestStatus = "Success";
      } catch (e) {
        console.log(e);
        state.requestStatus = "Error";
      }
    };

    fetchProjectsView();

    return {
      state,
      // Funcs
      formatIdr,
      makeReadableDateString,
    };
  },
});
</script>
<template>
  <div class="flex items-center">
    <div class="text-2xl font-semibold m-2">Project</div>
    <div v-if="state.requestStatus === 'Loading'" class="ml-1">
      <loading-icon class="text-2xl" />
    </div>
  </div>
  <div class="flex m-2 items-center">
    <router-link to="/projects/new">
      <button class="bg-blue-500 text-white font-semibold p-2 rounded-lg">
        New
      </button>
    </router-link>
    <input
      class="ml-2 w-full border-2 border-gray-400 p-2 rounded-lg"
      placeholder="Search project.."
    />
  </div>
  <div class="m-2">
    <ul v-if="state.projectsView">
      <li
        v-for="project in state.projectsView.projects"
        :key="project.project.id"
        class="shadow shadow-md border-2 border-gray-500 p-2 rounded-lg hover:bg-gray-100"
      >
        <router-link :to="`/projects/${project.project.id}`">
          <div>
            <div class="text-xl font-semibold">{{ project.project.name }}</div>
            <div>
              Date:
              {{ makeReadableDateString(new Date(project.project.startDate)) }}
            </div>
            <div>
              Revenue:
              <span class="font-semibold">{{ formatIdr(project.income) }}</span>
            </div>
            <div>
              Funding Cost:
              <span class="font-semibold">{{
                formatIdr(project.totalManufacturingPrice)
              }}</span>
            </div>
          </div>
        </router-link>
      </li>
    </ul>
  </div>
</template>