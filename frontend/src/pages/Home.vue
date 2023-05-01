<script lang="ts" setup>
import { onMounted, reactive, ref, watchEffect } from "vue";
import Tree, { TreeNode } from "primevue/tree";
import Toolbar from "primevue/toolbar";
import Button from "primevue/button";
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';

import { Run } from "../../wailsjs/go/queries/Query.js";
import { Disconnect, GetConnectionInfo } from "../../wailsjs/go/connections/Connection.js";
import { useRouter } from "vue-router";
import HomeData from "../components/HomeData.vue";
import HomeColumns from "../components/HomeColumns.vue";

const connectionInfo = ref({} as any);
const tables = ref([] as TreeNode[]);
const treeSelect = ref({});
const activeTable = ref('');
const router = useRouter();
const activeTab = ref(0);

const getConnectionDsn = async () => {
  const ci = await GetConnectionInfo();
  connectionInfo.value = ci;
};

const getAllTables = async () => {
  const qTables = await Run("SHOW TABLES");
  tables.value = [
    {
      key: "0",
      label: "Tables",
      children: qTables.Rows.map((t: any) => {
        return {
          key: `tables-${t[Object.keys(t)[0]]}`,
          label: t[Object.keys(t)[0]],
        };
      }),
    },
  ];
};

const disconnect = async() => {
  await Disconnect();
  router.push('/');
};

onMounted(() => {
  getConnectionDsn();
  getAllTables();
});

watchEffect(() => {
  const tableKey = Object.keys(treeSelect.value)[0];
  const tableNameFromKey = tableKey?.split("tables-");
  if (!tableNameFromKey || tableNameFromKey.length < 2) {
    return;
  }

  activeTable.value = tableNameFromKey[1];
});
</script>

<template>
  <div>
    <Toolbar>
      <template #start>
      </template>

      <template #end>
        <Button label="Disconnect" severity="danger" @click="disconnect()" />
      </template>
    </Toolbar>
  </div>
  <div class="flex">
    <div class="w-3">
      <Tree
        :value="tables"
        class="w-full h-screen overflow-x-auto"
        selectionMode="single"
        v-model:selectionKeys="treeSelect"
      />
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

  <div class="statusbar p-1">
  {{ connectionInfo.User }} @ {{ connectionInfo.Host }} / {{ connectionInfo.Database }}
  <span v-if="!!activeTable">
    / {{ activeTable }}
  </span>
  </div>
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
