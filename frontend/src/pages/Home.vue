<script lang="ts" setup>
import { onMounted, ref, watchEffect } from "vue";
import Button from "primevue/button";

import { GetConnectionInfo } from "../../wailsjs/go/connections/Connection.js";
import HomeData from "../components/HomeData.vue";
import HomeColumns from "../components/HomeColumns.vue";
import BToolbar from "../components/BToolbar.vue";
import BSidebar from "../components/BSidebar.vue";
import BStatusbar from "../components/BStatusbar.vue";


const activeTable = ref('');
const activeTab = ref(0);
</script>

<template>
  <div>
    <BToolbar />
  </div>
  <div class="flex">
    <div class="w-3">
      <BSidebar v-model:activeTable="activeTable" />
    </div>

    <div>
      <div class="p-inputgroup flex-1">
        <Button label="Data" :outlined="activeTab != 0" @click="activeTab = 0" />
        <Button label="Columns" :outlined="activeTab != 1" @click="activeTab = 1" />
        <Button label="Indexes" :outlined="activeTab != 2" @click="activeTab = 2" />
      </div>

      <div>
        <HomeData :tableName="activeTable" v-if="activeTab == 0" />
        <HomeColumns :tableName="activeTable" v-if="activeTab == 1" />
      </div>

    </div>
  </div>

  <BStatusbar :activeTable="activeTable" />
</template>

<style scoped>
.statusbar {
    position: fixed;
    z-index: 100; 
    bottom: 0; 
    left: 0;
    width: 100%;
    background: white;
}
</style>
