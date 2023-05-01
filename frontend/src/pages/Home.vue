<script lang="ts" setup>
import { onMounted, reactive, ref, watchEffect } from "vue";
import Tree, { TreeNode } from "primevue/tree";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Toolbar from "primevue/toolbar";
import Button from "primevue/button";

import { Run } from "../../wailsjs/go/queries/Query.js";
import { Disconnect, GetConnectionInfo } from "../../wailsjs/go/connections/Connection.js";
import { useRouter } from "vue-router";

const connectionInfo = ref({} as any);
const tables = ref([] as TreeNode[]);
const treeSelect = ref({});
const activeTable = ref('');
const queryResult = reactive({} as any);
const router = useRouter();

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
          key: `tables-${t.Tables_in_mysql}`,
          label: t.Tables_in_mysql,
        };
      }),
    },
  ];
};

const getAllTableRows = async (tableName: string) => {
  const result = await Run(`SELECT * FROM ${tableName}`);
  queryResult.Columns = (result.Columns || []).map((c: string, i: number) => {
    return {
      title: c,
      dataIndex: c,
      key: c,
    };
  });
  queryResult.Rows = result.Rows;
};

const disconnect = async() => {
  await Disconnect();
  router.push('/');
};

const setActiveTable = () => {
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
  if (!!activeTable.value) {
    getAllTableRows(activeTable.value);
  }
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
      <DataTable :value="queryResult.Rows || []">
        <Column
          v-for="c in queryResult.Columns || []"
          :key="c.key"
          :field="c.title"
          :header="c.title"
        />
      </DataTable>
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
