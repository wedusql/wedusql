<script lang="ts" setup>
import Tree, { TreeNode } from 'primevue/tree';
import { onMounted, ref, watchEffect } from 'vue';

import { Run } from "../../wailsjs/go/queries/Query.js";

const tables = ref([] as TreeNode[]);
const treeSelect = ref({});
const props = defineProps<{
  activeTable: string,
}>();

const emits = defineEmits(['update:activeTable']);

const getAllTables = async () => {
  const qTables = await Run("SHOW TABLES");
  tables.value = [
    ...tables.value,
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

const getAllProcedures = async () => {
  const qProcedures = await Run("SHOW PROCEDURE STATUS WHERE db = DATABASE() AND type = 'PROCEDURE'");
  tables.value = [
    ...tables.value,
    {
      key: "1",
      label: "Procedures",
      children: qProcedures.Rows.map((t: any) => {
        return {
          key: `procedures-${t.Name}`,
          label: t.Name,
        };
      }),
    },
  ];
};

const getAllFunctions = async () => {
  const qFunctions = await Run("SHOW FUNCTION STATUS WHERE db = DATABASE() AND type = 'FUNCTION'");
  tables.value = [
    ...tables.value,
    {
      key: "2",
      label: "Functions",
      children: qFunctions.Rows.map((t: any) => {
        return {
          key: `functions-${t.Name}`,
          label: t.Name,
        };
      }),
    },
  ];
};


onMounted(() => {
  getAllTables();
  getAllProcedures();
  getAllFunctions();
});

watchEffect(() => {
  const tableKey = Object.keys(treeSelect.value)[0];
  const tableNameFromKey = tableKey?.split("tables-");
  if (!tableNameFromKey || tableNameFromKey.length < 2) {
    return;
  }

  emits('update:activeTable', tableNameFromKey[1]);
});
</script>

<template>
  <Tree
    :value="tables"
    class="w-full h-screen overflow-x-auto"
    selectionMode="single"
    v-model:selectionKeys="treeSelect"
  />
</template>
